package main_test

import (
	"log"
	"os"
	"testing"
	// "net/http"
	// "net/http/httptest"
	// "strconv"
	// "encoding/json"
	// "bytes"
)

var a main.App

func TestMain(m *testing.M) {
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
	)

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}	
}

func clearTable() {
	a.DB.Exec("delete from products")
	a.DB.Exec("alter sequence products_id_seq restart with 1")
}

const tableCreationQuery = `create table if not exists products
(
	id serial,
	name text not null,
	price numeric(10,2) not null default 0.00,
	constraint products_pkey primary key (id)
)`