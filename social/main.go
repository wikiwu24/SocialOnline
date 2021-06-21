package main

import ("fmt"
        "github.com/gorilla/mux"
        "log"
        "net/http"  
        jwtmiddleware "github.com/auth0/go-jwt-middleware"
        jwt "github.com/form3tech-oss/jwt-go"


)

func main() {
 
    fmt.Println("started-service!!!!!!!!!!!")
    // only sign up successfully，we can run search and upload
    // middleware sits between the client and the server, it can check whether the request has the valid token
    // set up the jwtMiddleware
    jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
        ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
            return []byte(mySigningKey), nil
        },
        SigningMethod: jwt.SigningMethodHS256,
    })



    r := mux.NewRouter()
    r.Handle("/upload",  jwtMiddleware.Handler(http.HandlerFunc(uploadHandler))).Methods("POST", "OPTIONS")
    r.Handle("/search",  jwtMiddleware.Handler(http.HandlerFunc(searchHandler))).Methods("GET", "OPTIONS")
    r.Handle("/signin", http.HandlerFunc(signinHandler)).Methods("POST", "OPTIONS")
    r.Handle("/signup", http.HandlerFunc(signupHandler)).Methods("POST", "OPTIONS")
    // OPTIONS: 前后端在不同的域名上（domain）（cross domain）通过options特殊请求
    //http.HandleFunc("/upload", uploadHandler)
    log.Fatal(http.ListenAndServe(":8080", r))// equivalent to start Tomcat server in Java
    // 8080 is the port that is be listened
    // nil represents the default http router
    // or you can costomerize http router


}