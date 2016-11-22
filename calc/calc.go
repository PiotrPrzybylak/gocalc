package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	number1, _ := strconv.Atoi(r.FormValue("number1"))
	number2, _ := strconv.Atoi(r.FormValue("number2"))
	operation := r.FormValue("operation")

	var result int
	if operation == "-" {
		result = number1 - number2
	} else {
		result = number1 + number2
	}

	fmt.Fprintf(w, "<h1> %v %v %v = %v</h1>",
		number1,
		operation,
		number2,
		result)
}

func main() {
	http.HandleFunc("/calculate/", calculateHandler)

	http.ListenAndServe(":8080", nil)
}
