package main

import (
	"fmt"
	"html"
	"net/http"
	"os"
)

func main() {

	// 1つ目のハンドラは /hello で呼び出される
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<html><body>Hello</body></html>")
	})

	// 1つ目のハンドラは /hello で呼び出される
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		val := r.FormValue("say")

		fmt.Fprint(w, "<html><body>echo</body></html>",
			html.EscapeString(val))
	})

	// Web-server 8080
	err := http.ListenAndServe(":8080", nil)

	// Error Check
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
