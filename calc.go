package main

import (
	"net/http"
	"fmt"
	"strconv"
)




func calculateHandler(w http.ResponseWriter, r *http.Request) {
	number1, _ := strconv.Atoi(r.FormValue("number1"))
	number2, _ := strconv.Atoi(r.FormValue("number2"))
    fmt.Fprintf(w, "<h1>Result: %v</h1>",
        number1 + number2)
}

func main() {
	http.HandleFunc("/calculate/", calculateHandler)

	http.ListenAndServe(":8080", nil)
}
