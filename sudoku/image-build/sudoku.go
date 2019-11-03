package main

import (
	"fmt"
	"net/http"
)

func getHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
func solve(w http.ResponseWriter, r *http.Request) {
	if !Ready() {
		w.WriteHeader(500)
		return
	}
	fmt.Fprintf(w, "%s", Sudoku(r.FormValue("sudoku")))
}

func main() {
	server := http.Server{Addr: "0.0.0.0:1234"}
	http.HandleFunc("/health", getHealth)
	http.HandleFunc("/solve", solve)
	server.ListenAndServe()
}
