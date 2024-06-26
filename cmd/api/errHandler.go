package main

import (
	"fmt"
	"log"
	"net/http"
)

func (app *Application) errHandlerLog(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (app *Application) errHandlerNoti(err error, s string) {
	if err != nil {
		fmt.Println(s + err.Error())
	}
}

func (*Application) errHandlerhttp(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
