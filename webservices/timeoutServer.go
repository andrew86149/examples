package webservices

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func myHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func timeHandler2(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is:"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", Body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>\n", t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}

func MtimoutS() {
	PORT := ":8001"
	arguments := os.Args
	if len(arguments) != 2 {
		PORT = ":" + arguments[2]
	}
	fmt.Println("Using port number: ", PORT)

	m := http.NewServeMux()
	srv := &http.Server{
		Addr:         PORT,
		Handler:      m,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	m.HandleFunc("/time", timeHandler2)
	m.HandleFunc("/", myHandler2)

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
