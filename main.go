package main

import (
	"fmt"
	"net/http"

	"mood/funcs"
)

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	

	http.HandleFunc("/", funcs.Home)
	http.HandleFunc("/artist", funcs.PageArtist)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}