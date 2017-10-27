package main

import (
   
	"net/http"
	"strconv"// imports 
	"time"  // used
	"text/template"
	"math/rand"
)

type Templatedata struct{
	Message string
	//Count int
	Guess int

}



func templateHandler(w http.ResponseWriter, r *http.Request){
	
	message :="Guess a number between 1 and 20"
	    rand.Seed(time.Now().UTC().UnixNano())
		target:=rand.Intn(20-1)// picks a number between 1 and 20
		var cookie, err = r.Cookie("target")
		if err == nil {
			// If we could read it, try to convert its value to an int.
			target, _ = strconv.Atoi(cookie.Value)
		}
	
		Guess,_ := strconv.Atoi(r.FormValue("guess"))

		if Guess == target{
			message ="Congrats "+strconv.Itoa(Guess)+" was the answer"
		}else if Guess < target{
		   message="Try Again your guess  was  too low"// checks guess sent 
		}else {										  //  from the form
			message="Try Again your guess was too high"
		 }

		cookie = &http.Cookie{
			Name:    "target",
			Value:   strconv.Itoa(target),
			Expires: time.Now().Add(72 * time.Hour),
		}
		http.SetCookie(w, cookie)// set cookie
	
		
			
	 t, _ := template.ParseFiles("template/guess.html")// parses file
	 t.Execute(w, &Templatedata {Message:message, Guess:Guess})// sends data to html file
	
}

func main() {
	
	
	http.Handle("/", http.FileServer(http.Dir("./static")))// used when index is called 
	
	http.HandleFunc("/guess", templateHandler)// used when temp is called

    http.ListenAndServe(":8080", nil)
}

 