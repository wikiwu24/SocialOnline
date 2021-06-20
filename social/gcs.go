package main

import (
    "context"
    "fmt"
    "io"

    "cloud.google.com/go/storage"
)

const (
    BUCKET_NAME = "wuke-around-123"  
)

// upload files to GCS
// io.Reader： fileReader父类，从request获取文件
func saveToGCS(r io.Reader, objName string)(string, error) {
	// 返回值中还有返回一个url string交给elasticsearch存储
	// context 用来存储一些参数
	ctx:= context.Background();
    //constuct the connection
	// newClient 在第二个参数应该提供一个身份验证，userCredentials
	// 在这个，例子中没写，就是默认的，默认的是虚拟机的service account
	client, err := storage.NewClient(ctx);
    if(err != nil){
		return "", err
	}
	// get object;
	object := client.Bucket(BUCKET_NAME).Object(objName)

	wc := object.NewWriter(ctx)
	// io.Reader 是 file 的source， copy到wc
	if _, err = io.Copy(wc, r); err != nil{
		return "", err
	}
	if err := wc.Close(); err != nil {
        return "", err
    }

    if err := object.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
        return "", err
    }


    // return URL
	attrs, err := object.Attrs(ctx)
	if err != nil{
		return "", err
	}
	fmt.Printf("Image is saved to GCS: %s\n", attrs.MediaLink)

    return attrs.MediaLink, nil

}
