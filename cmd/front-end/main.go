package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"net/http"
)

//go:embed config/config.yaml
var configFile []byte

func getHomepage(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "homepage.html", nil)
}

func getAccounts(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "accounts.html", nil)
}

func getTransaction(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "transaction.html", nil)
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.FormValue("printTest"))
	tpl.ExecuteTemplate(w, "homepage.html", nil)
}

var tpl *template.Template

func main() {
	tpl, _ = tpl.ParseGlob("templates/*.html")
	http.HandleFunc("/", getHomepage)
	http.HandleFunc("/accounts", getAccounts)
	http.HandleFunc("/transaction", getTransaction)
	http.HandleFunc("/test", test)
	fmt.Println("Running on port 8080")
	http.ListenAndServe(":8080", nil)
}
