package main

import (
	"Background-Worker/src/go/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	// 用户相关路由
	router.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
	router.HandleFunc("/login", handlers.LoginUser).Methods("POST")
	router.HandleFunc("/captcha", handlers.GenerateCaptcha).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
