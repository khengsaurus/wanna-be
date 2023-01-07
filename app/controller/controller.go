package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/khengsaurus/wanna-be/consts"
	"github.com/khengsaurus/wanna-be/database"
	"github.com/khengsaurus/wanna-be/middlewares"
	"github.com/khengsaurus/wanna-be/util"
)

var AdminRouter = func(router chi.Router) {
	router.Use(middlewares.VerifyHeader("Authorization", "bearer token"))
	router.Get("/", Ping)
}

var UsersRouter = func(router chi.Router) {
	router.Get("/", QueryHandler(fmt.Sprintf("SELECT * FROM %s", consts.UsersTable)))
}

var ExpensesRouter = func(router chi.Router) {
	router.Get("/", QueryHandler(fmt.Sprintf("SELECT * FROM %s", consts.ExpensesTable)))
}

func QueryHandler(query string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		db := database.GetConnection()
		if db == nil {
			fmt.Print("Failed to create db connection")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		defer db.Close()

		rows, err := db.Query(query)
		if err != nil {
			fmt.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data, err := util.ToJsonEncodable(rows)
		if err != nil {
			fmt.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(data)

		if err != nil {
			fmt.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func Ping(w http.ResponseWriter, r *http.Request) {
	status := "failed"
	db := database.GetConnection()
	if db != nil {
		status = "success"
	} else {
		fmt.Print("Failed to create db connection")
	}

	defer db.Close()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Postgres db ping %s", status)))
}
