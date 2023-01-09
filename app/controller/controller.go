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
		conn, errorStatus := database.GetConnFromReqCtx(r)
		if errorStatus != http.StatusOK {
			w.WriteHeader(errorStatus)
			return
		}

		data, err := conn.RunQuery(query)
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

		conn, errorStatus := database.GetConnFromReqCtx(r)
		if errorStatus != 200 {
			w.WriteHeader(errorStatus)
			return
		}

		data, err := conn.RunQuery(fmt.Sprintf("SELECT * FROM %s WHERE userId = %s", table, id))
		util.SendRes(w, data, err)
	}
}

func Ping(w http.ResponseWriter, r *http.Request) {
	status := "failed"
	conn, errorStatus := database.GetConnFromReqCtx(r)
	if errorStatus != 200 {
		w.WriteHeader(errorStatus)
		return
	}

	err := conn.Ping()
	if err == nil {
		status = "success"
	} else {
		fmt.Println("Failed to ping db")
	}

	w.Write([]byte(fmt.Sprintf("Postgres db ping %s", status)))
}
