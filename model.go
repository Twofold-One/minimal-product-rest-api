package main

import (
	"database/sql"
)

type product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
}

func (p *product) getProduct(db *sql.DB) error {
	return db.QueryRow(`select name, price from products where id=$1`, p.ID).Scan(&p.Name, &p.Price)
}

func (p *product) updateProduct(db *sql.DB) error {
	_, err := db.Exec(`update products set name=$1, price=$2 where id=$3`, p.Name, p.Price, p.ID)
	return err
}

func (p *product) deleteProduct(db *sql.DB) error {
	_, err := db.Exec(`delete from products where id=$1`, p.ID)
	return err
}

func (p *product) createProduct(db *sql.DB) error {
	err := db.QueryRow(
		`insert into products(name, price) values($1, $2) returning id`,
		p.Name, p.Price).Scan(&p.ID)

	if err != nil {
		return err
	}
	return nil
}

func (p *product) getProducts(db *sql.DB, start, count int) ([]product, error) {
	rows, err := db.Query(
		`select id, name, price from products limit $1 offset $2`, count, start)
	
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []product{}

	for rows.Next() {
		var p product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}