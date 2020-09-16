package main

import (
	// "fmt"
	"html/template"
	"net/http"
)

func main() {
	// hello := "hello yu bn"
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		/*	fmt.Fprintf(w, `<!doctype html>
			<html>
			<head><title>Browser Starter</title></head>
			<link rel="stylesheet" type="text/css" href="/public/main.css">
			<body>
			<h1>shadow-cljs - Browser</h1>
			<div id="app"></div>

			<script type="text/javascript"  src="/public/main.js"></script>
			<script>starter.browser.init();</script>
			</body>
			</html>`)*/

		t, _ := template.ParseFiles("public/index.html")
		t.Execute(w, nil)

	})

	http.ListenAndServe(":9000", nil)
}
