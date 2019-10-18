package storage

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type SQLightStorage struct {
	path string
}

func NewSQLightStorage(path string) *SQLightStorage {
	return &SQLightStorage{
		path: path,
	}
}

func (ms SQLightStorage) Get() []QuoteInfo {
	db := connect()
	defer db.Close()

	rows, err := db.Query("select * from Quotes")
	checkErr(err)
	defer rows.Close()

	quotes := []QuoteInfo{}
	for rows.Next() {
		var quote string
		var id int
		err := rows.Scan(&id, &quote)
		checkErr(err)

		quotes = append(quotes, QuoteInfo{
			Who:   "Ingen",
			When:  "Now",
			Quote: quote,
		})
	}

	err = rows.Err()
	checkErr(err)

	return quotes
}

func (ms *SQLightStorage) Store(quote QuoteInfo) {
	db := connect()
	defer db.Close()

	tx, err := db.Begin()
	checkErr(err)

	stmt, err := tx.Prepare("INSERT INTO Quotes(Quote) values(?)")
	checkErr(err)
	defer stmt.Close()

	_, err = stmt.Exec(quote.Quote)
	checkErr(err)

	tx.Commit()
}

func connect() *sql.DB {
	// db, err := sql.Open("sqlite3", ":memory:")
	db, err := sql.Open("sqlite3", "./sqlite3.sql")
	checkErr(err)

	checkErr(db.Ping())

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Quotes (Id INTEGER PRIMARY KEY, Quote text NOT NULL)")
	checkErr(err)

	return db
}

func checkErr(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
