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

		// validate credit card
		Validate := validateCC(NewCreditCard.Number)

		fmt.Fprintf(w, "{ \"valid\": %t }", Validate)
	})

	http.ListenAndServe(":8080", nil)
}
