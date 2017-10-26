package main

import (
   
	"net/http"
	"text/template"
)

type Templatedata struct{
	Message string
	
}

func templateHandler(w http.ResponseWriter, r *http.Request){
	 t, _ := template.ParseFiles("template/guess.html")
	 t.Execute(w, Templatedata {Message: "Guess a number between 1 and 20"} )
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	
	http.HandleFunc("/guess", templateHandler)

    http.ListenAndServe(":8080", nil)
}

 