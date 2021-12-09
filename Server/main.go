package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("../Front/index.html")
	if err != nil {
		log.Println("parse template err->", err)
		return
	}
	var name string = "hello"
	var users []string = []string{"zhangsan", "lisi", "wangwu"}
	tpl.Execute(w, map[string]interface{}{
		"name":  name,
		"users": users,
	})
}

func adminpage(w http.ResponseWriter, r *http.Request) {

	var user = make(map[string]interface{})
	user["name"] = "admin"
	user["age"] = 23
	data, err := json.Marshal(user)
	if err != nil {
		log.Println("json marshal map data err->", err)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(data)
}

func main() {

	http.HandleFunc("/", homepage)
	http.HandleFunc("/get_json", adminpage)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Println("running error->", err)
		return
	}
}
