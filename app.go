package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type App struct {
	DB *sql.DB
	Router *mux.Router
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("postgresql://%s:%s@localhost:5432/%s", user, password, dbname)

	var err error
	a.DB, err = sql.Open("pgx", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
}

func (a *App) Run(addr string) {}