package main

import (
	"log"
	"net/http"

	user "github.com/VDmdm/goservice_test_task/controller"
)

func main() {
	http.HandleFunc("/", user.SayHelloName)
	http.HandleFunc("/insert-user", user.InsertUser)
	http.HandleFunc("/login", user.Login)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
