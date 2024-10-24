package main

import (
	"fmt"
	"net/http"

	"mood/funcs"
)

func main() {
	http.HandleFunc("/", funcs.Home)
	http.HandleFunc("/Artist/{id}", funcs.PageArtist)
	http.HandleFunc("/search", funcs.Search)
	// http.HandleFunc("/404", funcs.ErrorHandlerr)

	http.HandleFunc("/assets/", funcs.CssHandler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
