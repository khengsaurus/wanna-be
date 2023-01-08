package controller

import (
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
	router.Get("/{userId}", GetByUser(consts.UsersTable))
	router.Post("/", CreateUser)
}

var ExpensesRouter = func(router chi.Router) {
	router.Get("/", QueryHandler(fmt.Sprintf("SELECT * FROM %s", consts.ExpensesTable)))
	router.Get("/{userId}", GetByUser(consts.ExpensesTable))
	router.Get("/{userId}/total", GetTotalExpense)
	router.Post("/", CreateExpense)
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

		data, err := util.RunQuery(db, query)
		util.SendRes(w, data, err)
	}
}

func GetByUser(table string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "userId")
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		db := database.GetConnection()
		if db == nil {
			fmt.Print("Failed to create db connection")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer db.Close()

		data, err := util.RunQuery(db, fmt.Sprintf("SELECT * FROM %s WHERE userId = %s", table, id))
		util.SendRes(w, data, err)
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

	w.Write([]byte(fmt.Sprintf("Postgres db ping %s", status)))
}
