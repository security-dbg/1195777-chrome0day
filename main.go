package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)

	log.Println("Listening on :3000... test url http://localhost:3000/1195777.html")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
