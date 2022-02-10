package main

import (
	"html/template"
	"log"
	"net"
	"net/http"
	"path/filepath"
)

var major string
var minor string

func main() {
	sm := http.NewServeMux()
	sm.HandleFunc("/", serveIndexTemplate)
	log.Println("Starting Redweb - v" + major + "-" + minor)
	l, err := net.Listen("tcp4", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.Serve(l, sm))
	log.Println("Listening...")
}

func serveIndexTemplate(w http.ResponseWriter, r *http.Request) {
	title := "Test Redweb RPM - v" + major + "-" + minor
	log.Println(r.Method, "/")
	// Testing
	// lp := filepath.Join("templates", "index.html")
	// Live
	lp := filepath.Join("/usr/share/redweb/", "index.html")
	tmpl, _ := template.ParseFiles(lp)
	tmpl.Execute(w, title)
}
