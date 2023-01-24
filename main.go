package main

import(
	"net/http"
	"youthcenter/handler"
)

func main(){
	http.HandleFunc("/", handler.MainPage)

	http.ListenAndServe(":8000", nil)
}