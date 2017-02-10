package main

import (
	"flag"
	"net/http"
)

func main() {
	dir := flag.String("dir", "staticwebsite", "Directory of Static Website")
	port := flag.String("port", "8000", "Port number of dev server")
	flag.Parse()
	http.ListenAndServe(":"+*port, http.FileServer(http.Dir(*dir)))
}
