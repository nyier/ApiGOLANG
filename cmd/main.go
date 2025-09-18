package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Kubernetes with Go!")
	})
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "{\"status\": \"ok\"}")
	})

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	html := `<html>
	<head>
		<title>Isaac Rodriguez</title>
	</head>
	<body>


		<h1>Hello from Kubernetes with Go!</h1>
	</body>
	</html>`
	fmt.Fprintf(w, html)
}
