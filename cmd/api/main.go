package main

import "net/http"

type Application struct {
	port string
}

func main() {
	var app Application
	app.port = ":8080"

	err := http.ListenAndServe(app.port, app.routes())

	app.errHandlerLog(err)
}
