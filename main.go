package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func handleRequests() {

	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func main() {

	Articles = []Article{
		Article{Title: "Earphone", Desc: "In Ear earphone", Content: "1 Ear phone"},
		Article{Title: "Airpod", Desc: "Apple Airpods", Content: "Wireless Apple airpod"},
	}
	handleRequests()

}
