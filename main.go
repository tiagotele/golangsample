package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/prometheus/client_golang/prometheus/promhttp"

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

		mux.Handle("/metrics", promhttp.Handler())
		mux.HandleFunc("/", homeFunc)
		mux.HandleFunc("/list", listFunc)
		mux.HandleFunc("/other", requestOtherServiceFunc)
		

		
		errorChannel <- http.ListenAndServe(":8081", mux)
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

func requestOtherServiceFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/other" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	a := nextInt()
	b := nextInt()
	resp, err := http.Get(fmt.Sprintf("http://aluno:8080/demo/sum?a=%d&b=%d", a, b))

	if err != nil {
		log.Fatalln("Error on request: ", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(fmt.Sprintf("Sum of %d + %d = %s", a, b, string(body)))
	w.Write([]byte(fmt.Sprintf("Sum of %d + %d = %s", a, b, string(body))))
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}

func nextInt() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}
