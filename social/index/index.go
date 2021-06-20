package main

import (
        "context"
        "fmt"

        "github.com/olivere/elastic/v7"
)

const (
        POST_INDEX = "post"
        USER_INDEX = "user"
        ES_URL = "http://10.128.0.2:9200" // use internal IP
)
func main(){


	client, err := elastic.NewClient(
        elastic.SetURL(ES_URL),
        elastic.SetBasicAuth("elastic", "12345678"))
    if err != nil {
        panic(err)
    }
    // elastic 插入的数据可以比较随意，每条数据可以不一样，也可以插入mapping中没有的key-value，如课上的例子：插入了date：2021-0503
	// 通过定义index中的property，可以把最关心的数据类型规定了，其他的交给elasticsearch来判断最合适的存储类型
	exists, err := client.IndexExists(POST_INDEX).Do(context.Background())
	if err != nil{
		panic(err)
	}
	if !exists{
		mapping := `{
			"mappings":{
				"properties" :{
					"id":       {"type":"keyword"},
					"user" :    {"type":"keyword"},
					"message":  { "type": "text" },
                    "url":      { "type": "keyword", "index": false },
                    "type":     { "type": "keyword", "index": false } // property中的index是索引

				}
			}
		}`
		_,err := client.CreateIndex(POST_INDEX).Body(mapping).Do(context.Background())
		    if err != nil{
			   panic(err)
		    }
	}
	exists, err = client.IndexExists(USER_INDEX).Do(context.Background())
	if err != nil {
		panic(err)
	}
	
	if !exists {
		mapping := `{
					"mappings": {
						"properties": {
							"username": {"type": "keyword"},
							"password": {"type": "keyword"},
							"age":      {"type": "long", "index": false},
							"gender":   {"type": "keyword", "index": false}
						}
					}
					}`
	_, err = client.CreateIndex(USER_INDEX).Body(mapping).Do(context.Background())
	    if err != nil {
		    panic(err)
	    }
	}
	fmt.Println("Indexes are created.")

	

}

