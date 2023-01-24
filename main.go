package main

import(
	"net/http"
	"youthcenter/handler"
	"youthcenter/handler/user"
)

func main(){
	http.HandleFunc("/", handler.MainPage)
	http.HandleFunc("/user", user.RegistrationPage)

	http.ListenAndServe(":8000", nil)
}