package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type Users struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (app *Application) routes() http.Handler {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	r.Use(cors.Handler)

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

		var data Users

		err = json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			fmt.Print("error json ")
			return
		}

		fmt.Println("asdasd")
		// if username == "" || email == "" || password == "" {

		// 	fmt.Println("ooooooooooooooooooooo")
		// 	http.Error(w, "Missing required form values", http.StatusBadRequest)
		// 	return
		// }

		_, err = stmt.Exec(data.Username, data.Email, data.Password, data.Role)
		if err != nil {
			fmt.Println("asdasdasdsd")
			fmt.Println(err.Error())
			app.errHandlerhttp(w, err)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User Created!"))
	})

	return r
}
