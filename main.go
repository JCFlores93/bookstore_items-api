package main

import (
	"github.com/JCFlores93/bookstore_items-api/app"
	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func main() {
	app.StartApplication()
}
