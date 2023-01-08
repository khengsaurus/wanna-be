package util

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func RunQuery(db *sql.DB, query string) ([]interface{}, error) {
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	data, err := ToJsonEncodable(rows)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// https://stackoverflow.com/questions/42774467
func ToJsonEncodable(rows *sql.Rows) ([]interface{}, error) {
	columnTypes, err := rows.ColumnTypes()

	if err != nil {
		return nil, err
	}

	count := len(columnTypes)
	finalRows := []interface{}{}

	for rows.Next() {

		scanArgs := make([]interface{}, count)

		for i, v := range columnTypes {

			switch v.DatabaseTypeName() {
			case "VARCHAR", "TEXT", "UUID", "TIMESTAMP":
				scanArgs[i] = new(sql.NullString)
			case "BOOL":
				scanArgs[i] = new(sql.NullBool)
			case "INT4":
				scanArgs[i] = new(sql.NullInt64)
			default:
				scanArgs[i] = new(sql.NullString)
			}
		}

		err := rows.Scan(scanArgs...)

		if err != nil {
			return nil, err
		}

		data := map[string]interface{}{}

		for i, v := range columnTypes {

			if z, ok := (scanArgs[i]).(*sql.NullBool); ok {
				data[v.Name()] = z.Bool
				continue
			}

			if z, ok := (scanArgs[i]).(*sql.NullString); ok {
				data[v.Name()] = z.String
				continue
			}

			if z, ok := (scanArgs[i]).(*sql.NullInt64); ok {
				data[v.Name()] = z.Int64
				continue
			}

			if z, ok := (scanArgs[i]).(*sql.NullFloat64); ok {
				data[v.Name()] = z.Float64
				continue
			}

			if z, ok := (scanArgs[i]).(*sql.NullInt32); ok {
				data[v.Name()] = z.Int32
				continue
			}

			data[v.Name()] = scanArgs[i]
		}

		finalRows = append(finalRows, data)
	}

	return finalRows, nil
}

func SendRes(w http.ResponseWriter, data []interface{}, err error) {
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
	}
}
