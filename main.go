package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", errHandler)
	log.Println("Beginning to serve on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func errHandler(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	status, err := strconv.Atoi(code)
	if err != nil {
	 log.Printf("Failed to convert status:%v", err)
	 w.WriteHeader(500)
	 fmt.Fprintf(w, "Unknown code")
	 return
	}
	w.WriteHeader(status)
	fmt.Fprintf(w, "Code=%s", code)
   }
