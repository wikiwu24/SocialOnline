// elastic 插入数据
package main

import(
	"context"
	"github.com/olivere/elastic/v7"//官方library
)

const (
	ES_URL = "http://10.128.0.2:9200" 

)
//以下function：传入一个query，返回搜索结果
// 可以重复用在很多地方
func readFromES(query elastic.Query, index string)(*elastic.SearchResult, error){
	//1.connection
	client, err := elastic.NewClient(
		elastic.SetURL(ES_URL),
		elastic.SetBasicAuth("elastic","12345678"))
	if err != nil{
		return nil, err
	}
    // client.Search() 方法返回的是pointer
	searchResult,err := client.Search().Index(index).Query(query).Pretty(true).Do(context.Background())
	if err != nil{
		return nil, err
	}
	return searchResult,nil
}

// i interface{} 设置存储的object的类型为interface， 这样能存储任何类型的object
// elasticserach 会帮我们把object转成json
func saveToEs(i interface{}, index string, id string) error{
	client, err := elastic.NewClient(
		elastic.SetURL(ES_URL),
		elastic.SetBasicAuth("elastic", "12345678"),
	)
	if err != nil{
		return err
	}

	_, err = client.Index().// Index 可以理解为update
	     Index(index). // insert into post
		 Id(id).//Item. id
		 BodyJson(i). // 具体插入内容
		 Do(context.Background())
	return err
}