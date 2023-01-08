package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/khengsaurus/wanna-be/controller"
)

const port = 8080

func main() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		panic(envErr)
	}

	host, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}

	router := chi.NewRouter()
	router.Get("/", home)
	router.Route("/admin", controller.AdminRouter)
	router.Route("/users", controller.UsersRouter)
	router.Route("/expenses", controller.ExpensesRouter)

	fmt.Printf("Listening on %s:%v\n", host, port)

	err = http.ListenAndServe(fmt.Sprintf(":%v", port), router)
	if err != nil {
		panic(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	appId := os.Getenv("APP_ID")
	host, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Write([]byte(fmt.Sprintf("Hello from %s:%d, appId:%s", host, port, appId)))
	}
}
