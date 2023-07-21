package main

import (
    "fmt"
    "net/http"
    "strconv"
    "log"
)


func root_handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!\nThis request is a %s", r.URL.Path[1:], r.Method)
}

func example_1_handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%d", 2+2)
}

func example_2_handler(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query()
    fmt.Fprintf(w, "%s", q["number"])
}

func example_3_handler(w http.ResponseWriter, r *http.Request) {
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
    fmt.Fprintf(w, "%d",  num1 + num2)
}

func main() {
    http.HandleFunc("/", root_handler)
    http.HandleFunc("/example-1/", example_1_handler)
    http.HandleFunc("/example-2/", example_2_handler)
    http.HandleFunc("/example-3/", example_3_handler)
    http.HandleFunc("/rock_paper_scissors/", rock_paper_scissors)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
