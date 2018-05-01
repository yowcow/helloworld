package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var addr string
	var help bool
	flag.StringVar(&addr, "addr", ":9990", "address to bind")
	flag.BoolVar(&help, "help", false, "help")
	flag.Parse()

	if help || addr == "" {
		flag.Usage()
		os.Exit(1)
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Println(req.Method, req.URL)
		fmt.Fprintf(w, "Hello world")
	})
	svr := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	log.Println("booting up to say 'Hello world' listening to address " + addr)
	log.Println(svr.ListenAndServe())
}
