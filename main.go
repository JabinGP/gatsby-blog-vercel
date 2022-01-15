package main

import (
	"net/http"

	handler "github.com/JabinGP/gatsby-blog-vercel/api"
)

func main() {
	http.HandleFunc("/api/search", handler.Handler)
	err := http.ListenAndServe(":3022", nil)
	if err != nil {
		panic(err)
	}
}
