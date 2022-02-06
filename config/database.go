package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "localhost"
	database = "todos"
	user     = "mert"
	password = "123mert"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func DataBase() {
	var connectionString string = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)

	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	err = db.Ping()
	checkError(err)
	fmt.Println("Veri Tabanına Bağlanıldı...")

	_, err = db.Exec("DROP TABLE IF EXISTS inventory;")
	checkError(err)
	fmt.Println("Tablo Silindi Yeni Tablo Oluşturuldu..")

	_, err = db.Exec("CREATE TABLE inventory (id serial PRIMARY KEY, title VARCHAR(50), ımdb INTEGER);")
	checkError(err)

	// sqlStatment, err := db.Prepare("INSERT INTO inventory (title, ımdb) VALUES (?, ?);")
	// checkError(err)

	// res, err := sqlStatment.Exec("Sevgi", 5.0)
	// checkError(err)
	// rowCount, err := res.RowsAffected()
	// fmt.Printf("%d", rowCount)

	checkError(err)
	fmt.Println("Tablo Oluşturuldu...")

}

func InsertDataBase(title string, imdb float64) {
	var connectionString string = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)

	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	sqlStatment, err := db.Prepare("INSERT INTO inventory (title, ımdb) VALUES (?, ?);")
	checkError(err)

	res, err := sqlStatment.Exec(title, imdb)
	checkError(err)
	rowCount, err := res.RowsAffected()
	fmt.Printf("%d", rowCount)
	checkError(err)
}
