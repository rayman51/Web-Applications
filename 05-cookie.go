package main

import (
   
	"net/http"
	"strconv"
	"time"
	"text/template"
)

type Templatedata struct{
	Message string
	Count int
}


func templateHandler(w http.ResponseWriter, r *http.Request){
	count := 0
	
		// Try to read the cookie.
		var cookie, err = r.Cookie("count")
		if err == nil {
			// If we could read it, try to convert its value to an int.
			count, _ = strconv.Atoi(cookie.Value)
		}
	
		// Increase count by 1 either way.
		count += 1
	
		// Create a cookie instance and set the cookie.
		// You can delete the Expires line (and the time import) to make a session cookie.
		cookie = &http.Cookie{
			Name:    "count",
			Value:   strconv.Itoa(count),
			Expires: time.Now().Add(72 * time.Hour),
		}
		http.SetCookie(w, cookie)
	
		// Use a template to display out how many requests have been made.
		
		
	 t, _ := template.ParseFiles("template/guess.html")
	 t.Execute(w, Templatedata {Message: "Guess a number between 1 and 20", Count: count} )
	
}

func main() {
	
	
	http.Handle("/", http.FileServer(http.Dir("./static")))
	
	http.HandleFunc("/guess", templateHandler)

    http.ListenAndServe(":8080", nil)
}

 