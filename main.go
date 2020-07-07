package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/yalalovsm/solver/math"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/factorial/{number:[0-9]+}", factorialHandler).Methods("GET")

	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}

	log.Fatal(srv.ListenAndServe())
}

func factorialHandler(w http.ResponseWriter, r *http.Request) {
	number, _ := strconv.Atoi(mux.Vars(r)["number"])

	f := math.FactorialTree(number)

	responseWith(w, http.StatusOK, f.String())
}

func responseWith(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	w.Write([]byte(message))
}
