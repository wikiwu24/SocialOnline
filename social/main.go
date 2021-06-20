package main

import ("fmt"
        "github.com/gorilla/mux"
        "log"
        "net/http"  
)

func main() {
 
    fmt.Println("started-service")

    r := mux.NewRouter()
    r.Handle("/upload", http.HandlerFunc(uploadHandler)).Methods("POST", "OPTIONS")
    r.Handle("/search", http.HandlerFunc(searchHandler)).Methods("GET", "OPTIONS")
    // OPTIONS: 前后端在不同的域名上（domain）（cross domain）通过options特殊请求
    //http.HandleFunc("/upload", uploadHandler)
    log.Fatal(http.ListenAndServe(":8080", r))// equivalent to start Tomcat server in Java
    // 8080 is the port that is be listened
    // nil represents the default http router
    // or you can costomerize http router


}