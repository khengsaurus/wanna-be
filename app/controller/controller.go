package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/khengsaurus/wanna-be/consts"
	"github.com/khengsaurus/wanna-be/database"
	"github.com/khengsaurus/wanna-be/util"
)

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
			return
		}

		defer db.Close()

		rows, err := db.Query(query)
		if err != nil {
			fmt.Print(err)
		}

		data, err := util.ToJsonEncodable(rows)
		if err != nil {
			fmt.Print(err)
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(data)

		if err != nil {
			fmt.Print(err)
		}
	}
}
