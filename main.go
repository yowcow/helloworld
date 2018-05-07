package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var version = "version-string"

func main() {
	var addr string
	var help bool
	var ver bool
	flag.StringVar(&addr, "a", ":9990", "address to bind")
	flag.BoolVar(&help, "h", false, "print help")
	flag.BoolVar(&ver, "v", false, "print version")
	flag.Parse()

	if help || addr == "" {
		flag.Usage()
		os.Exit(1)
	}
	if ver {
		fmt.Println(version)
		os.Exit(0)
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Println(req.Method, req.URL)
		w.Header().Set("X-Hello-World-Version", version)
		fmt.Fprintln(w, "Hello world")
	})
	svr := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	log.Println("booting up to say 'Hello world' listening to address " + addr)
	log.Println(svr.ListenAndServe())
}
