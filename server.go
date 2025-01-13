/*
Rendering html
Ref: https://philipptanlak.com/web-frontends-in-go/#rendering-html-with-go
*/
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("index.html")
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/validatecc", func(w http.ResponseWriter, r *http.Request) {

		type CreditCard struct {
			Number string
		}

		var NewCreditCard CreditCard

		cc := json.NewDecoder(r.Body).Decode(&NewCreditCard)
		if cc != nil {
			log.Fatal(cc)
		}

		if NewCreditCard.Number == "" {
			fmt.Fprintf(w, "{ \"message\": \"Please provide a credit card number\" }")
		}

		// validate credit card
		Validate := validateCC(NewCreditCard.Number)

		fmt.Fprintf(w, "{ \"valid\": %t }", Validate)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
