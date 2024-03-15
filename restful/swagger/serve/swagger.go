package serve

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

var PORT = ":1234"

func Mswagger() {
	arguments := os.Args
	if len(arguments) >= 3 {
		PORT = ":" + arguments[2]
	}
	mux := mux.NewRouter()

	getMux := mux.Methods(http.MethodGet).Subrouter()
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	getMux.Handle("/docs", sh)
	getMux.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	s := http.Server{
		Addr:         PORT,
		Handler:      mux,
		ErrorLog:     nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	log.Println("Listening to", PORT)
	err := s.ListenAndServe()
	if err != nil {
		log.Printf("Error starting server: %s\n", err)
		return
	}
}
