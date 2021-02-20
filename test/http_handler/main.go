package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/user/{id}", paramHandler)
	r.HandleFunc("/regex/{id:[0-9]+}", paramHandler)
	r.Use(logging)

	err := http.ListenAndServe(":5000", r)
	if err != nil {
		log.Fatal(err)
	}

}
func logging(next http.Handler) http.Handler {
	fmt.Println("text", next)
	fmt.Scanln()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello. world!\n"))
}

func paramHandler(w http.ResponseWriter, r *http.Request) {
	variables := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, variables["id"])
}
