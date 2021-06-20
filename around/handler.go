package main

import (
    "encoding/json"
    "fmt"
    "net/http"
	"path/filepath"
	"github.com/pborman/uuid"

)

var (
    mediaTypes = map[string]string{
        ".jpeg": "image",
        ".jpg":  "image",
        ".gif":  "image",
        ".png":  "image",
        ".mov":  "video",
        ".mp4":  "video",
        ".avi":  "video",
        ".flv":  "video",
        ".wmv":  "video",
    }
)

// 为什么 ResponseWriter 是传入值， Request 传入的是pointer
// Responsewriter 在Go中是interface，Request是struct， 相当于class
// interface 不能被实例化，所以不存在引用（pointer)
// upload 基本逻辑：
// 从httprequest 拿到 post 内容
// 创建一个Post Object
// 调用post.go 中的savePost方法
// 先存储GCS， 有了url后存入elasticsearch
func uploadHandler (w http.ResponseWriter, r *http.Request){
	// Parse from body of request to get a json object.
	fmt.Println("received one post request")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
    // 支持跨域访问
    if r.Method == "OPTIONS" {
        return
    }
	// create the post we want to upload
	p:= Post{
        Id: uuid.New(),// use uuid library to create a global unique id
		User: r.FormValue("user"), // extrat the user name from form-data 
		Message: r.FormValue("message"),

	}
    // get file from form-data, header can contain the information about the file including the suffix
	file,header,err := r.FormFile("media_file")
	if err != nil{
		http.Error(w, "Media file is not available", http.StatusBadRequest)
		fmt.Printf("Media file is not available %v\n", err)
		return
	}
    // get the type of file according to the prefix.
	suffix := filepath.Ext(header.Filename)
	if t, ok := mediaTypes[suffix]; ok{
		p.Type = t

	}else{
		p.Type = "unknown"
	}

	err = savePost(&p,file)
	if err != nil{
		http.Error(w, "Failed to save post to GCS or Elasticsearch", http.StatusInternalServerError)
        fmt.Printf("Failed to save post to GCS or Elasticsearch %v\n", err)

		return 
	}
    fmt.Println("Post is saved successfully.")




    /*
	   abandon the following methods since the post we want to upload may not be in the json format
	
	 decoder := json.NewDecoder(r.Body)// return an json format string
	 // convert the json string into Post
	   var p Post
	 // the following code is equavalent to try catch in Java, used for handle expections
	 // if there is any err --- panic
     // decoder.Decode(&p): convert into Post object, the function Decode should take the reference of p instead of the copy
	 // so that you can really change the value of p
	 // There is another way to do so: (var p *Post; decoder.Decode(p))
	 if err := decoder.Decode(&p); err != nil{
	  // panic is equivalent to "throw exception" in java, stop the program
	   panic(err) 
	 }

	// Print in response body
	fmt.Fprintf(w, "Post received: %s\n", p.Message)*/
}

// 根据用户的参数调用searchByUser || searchByKeywords
func searchHandler (w http.ResponseWriter, r *http.Request){

	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
    w.Header().Set("Content-Type", "application/json")
    if r.Method == "OPTIONS" {
        return
    }

	user := r.URL.Query().Get("user")
	keywords := r.URL.Query().Get("keywords")
    
	var posts []Post
	var err error
    
	// 注意： 不能写成user != nil, user 的类型是string，和nil不匹配
	if user != ""{
		posts, err = searchPostByUser(user)
	}else{
		posts, err = searchPostByKeywords(keywords)
	}

	
	if err != nil{
		http.Error(w, "Failed to read data from ElasticSearch", http.StatusInternalServerError)
        return 
	}

	// 将结果以json格式返回给前端
	// 在Go中，将Go的object convert成json的api是Marshal
	js, err := json.Marshal(posts)
	if err != nil{
		http.Error(w, "Falied to parse post into json", http.StatusInternalServerError)
		return 
	}
	// write the result into response body
	w.Write(js)

}