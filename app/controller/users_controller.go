package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/khengsaurus/wanna-be/consts"
	"github.com/khengsaurus/wanna-be/database"
)

type CreateUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user CreateUserReq
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Printf("%v", err)
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

	sql := fmt.Sprintf(
		`INSERT INTO %s(username, password) VALUES ($1, $2)`,
		consts.UsersTable,
	)

	_, err = db.Exec(sql, user.Username, user.Password)

	if err != nil {
		fmt.Printf("Failed to write to table %s: %v", consts.UsersTable, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
