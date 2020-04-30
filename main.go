package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type PageVariables struct {
	Sentiment string
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	http.HandleFunc("/", AnalyzePage)
	http.HandleFunc("/results", ResultsPage)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func AnalyzePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
  <!DOCTYPE html>
  <html>
  <head>
    <title>Hello World</title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
  </head>
  <body>
    <div class="container">
      <p>Docker and Cypress Tutorial</p>
	  <form action="/results">
	  	<button type="submit" class="btn btn-primary">Click Me!</button>
	  </form>
   </div>
  </body>
  </html>
  `)
}

func ResultsPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
	<!DOCTYPE html>
	<html>
	<head>
	  <title>Hello World</title>
	  <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
	</head>
	<body>
	  <div class="results container">
		<p>Thanks for clicking the button!</p>
		<form action="/">
		  <button type="submit" class="btn btn-primary">Go Back</button>
		  </form>
	  </div>
	</body>
	</html>
  `)
}