package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Movies struct {
	ID    int
	Title string
	Imdb  int
}

const (
	host     = "localhost"
	database = "todos"
	user     = "mert"
	password = "123mert"
)

var connectionString string = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)

func Imdb9SQL() map[string]int {

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	res, err2 := db.Query("SELECT * FROM inventory WHERE ımdb = 9;")

	if err2 != nil {
		log.Fatal(err)
	}

	var movie Movies

	result := make(map[string]int)

	for res.Next() {

		err := res.Scan(&movie.ID, &movie.Title, &movie.Imdb)
		result[movie.Title] = movie.Imdb
		if err != nil {
			log.Fatal(err)
		}

	}
	return result
}

func Imdb8SQL() map[string]int {

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	res, err2 := db.Query("SELECT * FROM inventory WHERE ımdb = 8;")

	if err2 != nil {
		log.Fatal(err)
	}

	var movie Movies

	result := make(map[string]int)

	for res.Next() {

		err := res.Scan(&movie.ID, &movie.Title, &movie.Imdb)
		result[movie.Title] = movie.Imdb
		if err != nil {
			log.Fatal(err)
		}

	}
	return result
}

func Imdb7SQL() map[string]int {

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	res, err2 := db.Query("SELECT * FROM inventory WHERE ımdb = 7;")

	if err2 != nil {
		log.Fatal(err)
	}

	var movie Movies

	result := make(map[string]int)

	for res.Next() {

		err := res.Scan(&movie.ID, &movie.Title, &movie.Imdb)
		result[movie.Title] = movie.Imdb
		if err != nil {
			log.Fatal(err)
		}

	}
	return result
}

func Imdb6SQL() map[string]int {

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	res, err2 := db.Query("SELECT * FROM inventory WHERE ımdb = 6;")

	if err2 != nil {
		log.Fatal(err)
	}

	var movie Movies

	result := make(map[string]int)

	for res.Next() {

		err := res.Scan(&movie.ID, &movie.Title, &movie.Imdb)
		result[movie.Title] = movie.Imdb
		if err != nil {
			log.Fatal(err)
		}

	}
	return result
}

func Imdb5SQL() map[string]int {

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	res, err2 := db.Query("SELECT * FROM inventory WHERE ımdb = 5;")

	if err2 != nil {
		log.Fatal(err)
	}

	var movie Movies

	result := make(map[string]int)

	for res.Next() {

		err := res.Scan(&movie.ID, &movie.Title, &movie.Imdb)
		result[movie.Title] = movie.Imdb
		if err != nil {
			log.Fatal(err)
		}

	}

	return result
}

func FilterK(key map[string]int) {
	for k, v := range key {
		fmt.Println(k, v)
		continue
	}

}
