package main

import (
	"io"
	"net/http"
	"os"
	"fmt"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello from web server " + os.Getenv("HELLO_IDENTIFIER"))
}

func main() {
	fmt.Println("Web server " + os.Getenv("HELLO_IDENTIFIER") + " is running...")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}

