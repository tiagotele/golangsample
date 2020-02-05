package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	errorChannel := make(chan error)
	fmt.Println("Starting app")

	go func() {
		fmt.Println("Starting signal")
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errorChannel <- fmt.Errorf("%s", <-c)
	}()

	go func() {

		fmt.Println("Starting server")

		mux := http.NewServeMux()

		mux.HandleFunc("/", homeFunc)
		mux.HandleFunc("/list", listFunc)

		errorChannel <- http.ListenAndServe(":8080", mux)
	}()

	fmt.Fprintf(os.Stderr, "server killed: %s", <-errorChannel)
}

func homeFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	fmt.Println("Home endpoint")
	w.Write([]byte(`Home endpoint`))
}

func listFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/list" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	fmt.Println("List endpoint")
	w.Write([]byte(`List endpoint`))
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	fmt.Println("Not found")
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}
