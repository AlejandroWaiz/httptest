package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main () {
	http.HandleFunc("/double", doubleHandler)

	err := http.ListenAndServe(":8080", handler())

	if err != nil {
		panic(err)
	}


}

func handler() http.Handler {

	r := http.NewServeMux()
	r.HandleFunc("/double", doubleHandler)

	return r

}

func doubleHandler (w http.ResponseWriter, r *http.Request) {

	time.Sleep(1*time.Second)

	text := r.FormValue("v")

	if text == "" {

		http.Error(w, "Missing value", http.StatusBadRequest)

		return

	}

	value, err := strconv.Atoi(text)

	if err != nil {

		http.Error(w, "Not a number", 400)
		return

	}

	_, err = fmt.Fprintln(w, value*2)

	if err != nil {

		fmt.Printf("Error printing the value: %v", err)

	}


}

//curl "http://localhost:8080/double?v=<INSERTYOURVALUEHERE>"
