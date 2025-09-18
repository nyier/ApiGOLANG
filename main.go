package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := `<html>
		<head>
			<title>Isaac Rodriguez</title>
		</head>
		<body>
			<h1>Hello from Kubernetes with Go!</h1>
			<p>This is my CV running on Go + Containers</p>
		</body>
		</html>`
		fmt.Fprintf(w, html)
	})

	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed:", err)
	}
}
