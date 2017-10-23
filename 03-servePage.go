package main

import (
	"net/http"
	"strconv"
	"time"
	"text/template"
)

type Counter struct {
	Count int
}

func cookieHandler(w http.ResponseWriter, r *http.Request) {
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
    t, _ := template.ParseFiles("index.html")
    t.Execute(w, &Counter{Count: count})
}

func main() {
	// Call cookie handler for every request, unless specified below.
	http.HandleFunc("/", cookieHandler)
	// Send a 404 for favicon.ico
	http.Handle("/favicon.ico", http.NotFoundHandler())
	// Serve on port 8080.
	http.ListenAndServe(":8080", nil)
}