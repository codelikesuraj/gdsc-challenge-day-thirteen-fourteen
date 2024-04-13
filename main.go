package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = "3000"

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))

	fmt.Printf("Server listening on %q\n", PORT)
	fmt.Printf("Visit the URL http://127.0.0.1:%s to check if the connection was successful\n", PORT)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil))
}
