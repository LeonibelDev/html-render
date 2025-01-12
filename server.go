/*
Rendering html
Ref: https://philipptanlak.com/web-frontends-in-go/#rendering-html-with-go
*/
package main

import (
	"encoding/json"
	"html/template"
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

		var cc CreditCard

		err := json.NewDecoder(r.Body).Decode(&cc)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// validate cc
		isValid := validateCC(cc.Number)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]bool{
			"valid": isValid,
		})
	})

	http.ListenAndServe(":8080", nil)
}
