package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (app *Application) routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Lock"))
	})

	r.Post("/signup", func(w http.ResponseWriter, r *http.Request) {
		db := app.connectToDB()
		if db == nil {
			http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		stmt, err := db.Prepare("INSERT INTO Users (Username, Email, Password, role) VALUES (?, ?, ?, ?)")
		if err != nil {
			app.errHandlerhttp(w, err)
			return
		}
		defer stmt.Close()

		err = r.ParseForm()
		if err != nil {
			app.errHandlerhttp(w, err)
			return
		}

		username := r.FormValue("Username")
		email := r.FormValue("Email")
		password := r.FormValue("Password")
		role := r.FormValue("role")

		if username == "" || email == "" || password == "" || role == "" {
			http.Error(w, "Missing required form values", http.StatusBadRequest)
			return
		}

		_, err = stmt.Exec(username, email, password, role)
		if err != nil {
			app.errHandlerhttp(w, err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User Created!"))
	})

	return r
}
