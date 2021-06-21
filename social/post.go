package main

import(
    "mime/multipart"
    "reflect"

    "github.com/olivere/elastic/v7"

)
const(
    POST_INDEX = "post"
)
 
type Post struct {
	Id      string `json:"id"` 
    User    string `json:"user"`
    Message string `json:"message"`
    Url     string `json:"url"` // 这里程序只支持每个post传一个图片Url， 要扩展：变成string array
    Type    string `json:"type"`// 分为图片或视频
	// json:"":similiar to Jackson in Java, conver into json
	 
}
// Add functions to support user-based search and keyword-based search.
// functions will call readFromES function in elasticsearch.go
func searchPostByUser(user string)([]Post, error){
    query := elastic.NewTermQuery("user", user)// select * from post where user = ?
    searchResult, err := readFromES(query,POST_INDEX)
    if err != nil {
        return nil, err
    }
    // check every item in serachResult, if the item 符合 Post的格式，就把item cast成Post
    var ptyp Post
    var posts []Post
    for _, item := range searchResult.Each(reflect.TypeOf(ptyp)){
        p := item.(Post) // cast
        posts = append(posts, p)
    }
    return posts, nil
}

func searchPostByKeywords(keywords string)([]Post, error){
    query := elastic.NewMatchQuery("message", keywords)// select * from post where user = ?
    // 如果有很多关键词，取并集
    query.Operator("AND")
    // 如果keywords为空，没有提供关键词，则返回所有的post
    if keywords == ""{
        query.ZeroTermsQuery("all")
    }
    searchResult, err := readFromES(query, POST_INDEX)
    if err != nil {
        return nil, err
    }
    // check every item in serachResult, if the item 符合 Post的格式，就把item cast成Post
    var ptyp Post
    var posts []Post
    for _, item := range searchResult.Each(reflect.TypeOf(ptyp)){
        p := item.(Post) // cast
        posts = append(posts, p)
    }
    return posts, nil
    
}
//multipart.File http request 中的文件
func savePost(post *Post,  file multipart.File) error{
    mediaLink, err := saveToGCS(file, post.Id)
    if err != nil{
        return err
    }

    //先存GCS，得到url
    post.Url = mediaLink
    return saveToEs(post, POST_INDEX, post.Id)

}

func deletePost(userId string, postId string)error{
    query := elastic.NewBoolQuery()
    query.Must(elastic.NewTermQuery("user",userId))
    query.Must(elastic.NewTermQuery("id",postId))

    return deleteFromES(query, POST_INDEX)
   
}

