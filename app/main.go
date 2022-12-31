package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

const port = 8080

func main() {
	host, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}

	router := chi.NewRouter()
	router.Get("/ping", test)
	router.Get("/", home)

	fmt.Printf("Running on %s:%v\n", host, port)
	err = http.ListenAndServe(fmt.Sprintf(":%v", port), router)
	if err != nil {
		panic(err)
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Pong"))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	host, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write([]byte(fmt.Sprintf("Hello from %s", host)))
	}
}
