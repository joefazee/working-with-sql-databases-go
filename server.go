package main

import (
	"encoding/json"
	"log"
	"net/http"
	"working-database/models"
)

func (app *application) serve() error {
	srv := http.Server{
		Handler: app.handlers(),
		Addr:    ":4005",
	}

	err := srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (app *application) handlers() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.homePage)

	return mux

}

func (app *application) homePage(w http.ResponseWriter, req *http.Request) {

	f := models.Filter{Page: 1, PageSize: 20}
	users, meta, err := app.Models.Users.GetAll(f)

	if err != nil {
		log.Fatalln(err)
	}

	res := struct {
		Users []models.User
		Meta  models.Metadata
	}{
		Users: users,
		Meta:  meta,
	}

	js, err := json.Marshal(res)
	if err != nil {
		log.Fatalln(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(js)

}
