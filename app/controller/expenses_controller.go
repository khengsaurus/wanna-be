package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/khengsaurus/wanna-be/consts"
	"github.com/khengsaurus/wanna-be/database"
	"github.com/khengsaurus/wanna-be/util"
)

type CreateExpenseReq struct {
	UserId   int     `json:"userId"`
	Currency string  `json:"currency"`
	Amount   float32 `json:"amount"`
	Note     string  `json:"note"`
}

func GetTotalExpense(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "userId")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	conn, errorStatus := database.GetPgConnFromReq(r)
	if errorStatus != 200 {
		w.WriteHeader(errorStatus)
		return
	}

	sql := fmt.Sprintf(`
	SELECT 
		currency, ROUND(SUM(amount) * 100)/100 as total 
	FROM %s WHERE userId=%s
	GROUP BY currency`,
		consts.ExpensesTable,
		id,
	)

	data, err := conn.RunQuery(sql)
	util.SendRes(w, data, err)
}

func CreateExpense(w http.ResponseWriter, r *http.Request) {
	var expense CreateExpenseReq
	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		fmt.Printf("%v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	conn, errorStatus := database.GetPgConnFromReq(r)
	if errorStatus != 200 {
		w.WriteHeader(errorStatus)
		return
	}

	sql := fmt.Sprintf(
		`INSERT INTO %s(userId, currency, amount, note, date) VALUES ($1, $2, $3, $4, $5)`,
		consts.ExpensesTable,
	)

	_, err = conn.Exec(sql, expense.UserId, expense.Currency, expense.Amount, expense.Note, time.Now())

	if err != nil {
		fmt.Printf("Failed to write to table %s: %v", consts.ExpensesTable, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
