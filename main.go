package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

/**
 * @api {get} / Request User information
 * @apiName Get
 *
 */
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!\n") // send data to client side
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":8063", r) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
