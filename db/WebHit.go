package db

import (
	"fmt"

	"gitlab.com/gregcrobertson/millicent/model"
)

const webhitTable = "app.webhit"

// InsertWebHit ...
func InsertWebHit(wh model.WebHit) (int, error) {

	statement := fmt.Sprintf("INSERT INTO %s (useragent, ip, port, forwardedfor) VALUES (?, ?, ?, ?)", webhitTable)

	result, err := db.Exec(statement, wh.UserAgent, wh.IP, wh.Port, wh.ForwardedFor)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return int(id), err
	}

	return int(id), nil
}

// SelectWebHits ...
func SelectWebHits() ([]model.WebHit, error) {

	statement := fmt.Sprintf("SELECT useragent, ip, port, forwardedfor FROM %s", webhitTable)

	rows, err := db.Query(statement)

	if err != nil {
		return []model.WebHit{}, err
	}

	defer rows.Close()

	var webhits []model.WebHit

	for rows.Next() {

		var useragent string
		var ip string
		var port string
		var forwardedfor string

		if err := rows.Scan(&useragent, &ip, &port, &forwardedfor); err != nil {
			// Check for a scan error
			// Query rows will be closed with defer
			return []model.WebHit{}, err
		}

		webhits = append(webhits, model.WebHit{
			UserAgent:    useragent,
			IP:           ip,
			Port:         port,
			ForwardedFor: forwardedfor,
		})
	}

	// Any error encountered during iteration ...
	err = rows.Err()
	if err != nil {
		return webhits, err
	}

	return webhits, nil
}

// DeleteWebHits ...
func DeleteWebHits() (int, error) {

	statement := fmt.Sprintf("DELETE FROM %s", webhitTable)

	rows, err := db.Exec(statement)

	if err != nil {
		return 0, err
	}

	r, err := rows.RowsAffected()

	if err != nil {
		return int(r), err
	}

	return int(r), nil
}

// @TODO: Convert into a proper CRUD, re-use for ICIJ project ...
