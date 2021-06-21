package main

import(
	"fmt"
	// reading information from elasticsearch and convert it into user object
	"reflect"
	// store user information in elasticsearch
	"github.com/olivere/elastic/v7"
)
const(
	USER_INDEX = "user"
)
type User struct{
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int64  `json:"age"`
	Gender   string `json:"gender"`
}
// check the user in the database
func checkUser(username, password string)(bool, error){
	// 读取elasticsearch data 的function
	// 调用自己写的readFromES,搜索username和password的组合存不存在
    //
	query := elastic.NewBoolQuery() // select user from USER_INDEX where id == id and pw = pw
	query.Must(elastic.NewTermQuery("username", username))
	query.Must(elastic.NewTermQuery("password", password))
	searchResult,err := readFromES(query, USER_INDEX)
	
	if(err != nil){
		return false, err
	}
	var utype User
	// TypeOf should not take the name of the class
	// 加入class 的对应的object
    // linear scan the result and check if the username&password matches 
	for _, item := range searchResult.Each(reflect.TypeOf(utype)){
       u := item.(User)
	   if u.Password == password && u.Username == username{
		   return true, nil
	   }
	}
	return false, nil
}


    // run addUser when register a user
	// user is the pointer of the User object
	func addUser(user *User)(bool, error){
		query := elastic.NewTermQuery("username", user.Username)
		searchResult, err := readFromES(query, USER_INDEX)
		if err != nil {
			return false, err
		}
		// if the username has already exist, return false
		// if we do not check the duplicate and the duplicate actually exist:
		// elasticsearch will update the original data
		if searchResult.TotalHits() > 0 {
			return false, nil
		}

		err = saveToEs(user, USER_INDEX, user.Username)
		if err != nil {
			return false, err
		}
		fmt.Printf("User is added : %s\n", user.Username)   
		return true, nil
	
	}


