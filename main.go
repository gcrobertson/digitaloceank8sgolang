package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
	"gitlab.com/gregcrobertson/millicent/db"
	"gitlab.com/gregcrobertson/millicent/service"
)

/* Mux creates http router, listens for requests, and delegates responses
 *
 *
 */
func main() {

	// MySQL ...
	err := db.OpenMySQL()
	if err != nil {
		// @TODO: handle error
	}
	defer db.CloseMySQL()

	// MUX ...
	router := mux.NewRouter()
	router.HandleFunc("/getWebHits", getWebHits)
	router.HandleFunc("/", index)

	// Middleware ...
	router.Use(clientTrackingMiddleware)

	staticDir := "/public/"
	router.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))
	http.ListenAndServe(":8080", router)
}

/*
 *
 *
 *
 *
 */
func index(respWriter http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles(
		"view/index.gohtml",
		"view/_header.gohtml",
	)
	tpl = template.Must(tpl, err)
	tpl.ExecuteTemplate(respWriter, "index.gohtml", nil)
}

/* @TODO
 *
 *
 *
 *
 */
func clientTrackingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// https://github.com/gorilla/mux#middleware
		// https://golang.org/pkg/net/http/#Request
		// https://stackoverflow.com/questions/27234861/correct-way-of-getting-clients-ip-addresses-from-http-request
		wh := service.WebHit(r)

		_, err := db.InsertWebHit(wh) // id

		if err != nil {
			// @TODO: Logging
		}

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

/* @TODO
 *
 *
 *
 *
 */
func getWebHits(respWriter http.ResponseWriter, req *http.Request) {

	/*	@TODO:  security! I know this is not good!
	 *			and passed via the URL for the simple reason of practicing K8s cronjobs.
	 *
	 *
	 */
	param := req.URL.Query().Get("fake-api-key-x") // ensure it comes from the cronjob
	if param != "qQVO8StCP6v4vGMQ3LYz" {
		// Email ...
		s := fmt.Sprintf("Unsuccessful attempt to email! The API key passed was %s\n", param)
		service.SendEmail(s)

		// HTTP Response ...
		respWriter.Write([]byte(strconv.Itoa(http.StatusForbidden)))
		return
	}

	webhits, err := service.GetWebHits()
	if err != nil {

		// Email
		s := fmt.Sprintf("Unsuccessful attempt to email! The error retrieving DB records was %s\n", err.Error())
		service.SendEmail(s)

		// Response
		respWriter.Write([]byte(strconv.Itoa(http.StatusInternalServerError)))
		return
	}

	err = service.SendWebHitsEmail(webhits)
	if err != nil {

		// Email
		s := fmt.Sprintf("Unsuccessful attempt to email! The error sending email was %s\n", err.Error())
		service.SendEmail(s)

		// Response
		respWriter.Write([]byte(strconv.Itoa(http.StatusInternalServerError)))
		return

	}

	// Delete all records from table ...
	r, err := service.DeleteWebHits()
	if err != nil {

		// Email
		s := fmt.Sprintf("Unable to delete web hits! The count [%v] and error sending email was %s\n", r, err.Error())
		service.SendEmail(s)

		// Response
		respWriter.Write([]byte(strconv.Itoa(http.StatusInternalServerError)))
		return
	}

	// Response
	respWriter.Write([]byte(strconv.Itoa(http.StatusOK)))
	return
}
