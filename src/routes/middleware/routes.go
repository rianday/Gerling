package middleware

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", index)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hellow")
}

func StartWebServer() {
	http.ListenAndServe(":8080", nil)
}
