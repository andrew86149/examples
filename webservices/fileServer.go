package webservices

import (
	"fmt"
	"log"
	"net/http"
)

var PORTf = ":8765"

func defaultHandler2(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusOK)
	Body := "Thanks for visiting!\n"
	fmt.Fprintf(w, "%s", Body)
}

func MfileServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultHandler2)

	fileServer := http.FileServer(http.Dir("/tmp/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	fmt.Println("Starting server on:", PORTf)
	err := http.ListenAndServe(PORTf, mux)
	fmt.Println(err)
}
