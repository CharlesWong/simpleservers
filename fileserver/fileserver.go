package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// This works and strip "/static/" fragment from path
	addr := flag.String("addr", ":8080", "proxy listening address")
	folder := flag.String("folder", "./static/", "static file folder")
	fs := http.FileServer(http.Dir(*folder))
	http.Handle("/", http.StripPrefix("/", fs))

	log.Fatal(http.ListenAndServe(*addr, nil))
}
