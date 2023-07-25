package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "Hi there, I love %s!\nThis request is a %s", r.URL.Path[1:], r.Method)
}

func example1Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%d", 2+2)
}

func example2Handler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	fmt.Fprintf(w, "%s", q["number"])
}

func example3Handler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	num1, err := strconv.Atoi(q["number1"][0])
	if err != nil {
		fmt.Fprintf(w, "'%s' is not parseable as a number", q["number1"][0])
		return
	}
	num2, err := strconv.Atoi(q["number2"][0])
	if err != nil {
		fmt.Fprintf(w, "'%s' is not parseable as a number", q["number2"][0])
		return
	}
	fmt.Fprintf(w, "%d", num1+num2)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/example-1/", example1Handler)
	http.HandleFunc("/example-2/", example2Handler)
	http.HandleFunc("/example-3/", example3Handler)
	http.HandleFunc("/rock_paper_scissors/", rockPaperScissors)
	fmt.Println("Running at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
