package controller

import (
	"fmt"
	"net/http"

	"github.com/khengsaurus/wanna-be/database"
	"github.com/khengsaurus/wanna-be/util"
)

func PingPg(w http.ResponseWriter, r *http.Request) {
	status := "failed"
	conn, errorStatus := database.GetPgConnFromReq(r)
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

func GetPgConnPoolDetails(w http.ResponseWriter, r *http.Request) {
	pgConnPool, err := database.GetPgConnPoolFromReq(r)
	if err != nil {
		fmt.Printf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res := pgConnPool.GetRepr()
	util.SendRes(w, res, nil)
}
