package storage

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3" // db-driver
)

// SQLightStorage uses a sqlite3 db to store quoteInfo in
type SQLightStorage struct {
	path string
}

// NewSQLightStorage creates new SQLightStorage
func NewSQLightStorage(path string) *SQLightStorage {
	return &SQLightStorage{
		path: path,
	}
}

// Get fetches all QuoteInfo from the db
func (ms SQLightStorage) Get() []QuoteInfo {
	db := connect()
	defer db.Close()

	rows, err := db.Query("select * from Quotes ORDER BY Time DESC")
	checkErr(err)
	defer rows.Close()

	quotes := []QuoteInfo{}
	for rows.Next() {
		var quote, who, where string
		var when time.Time

		var id int
		err := rows.Scan(&id, &quote, &who, &where, &when)
		checkErr(err)

		quotes = append(quotes, QuoteInfo{
			Quote: quote,
			Who:   who,
			Where: where,
			When:  when,
		})
	}

	err = rows.Err()
	checkErr(err)

	return quotes
}

// Store inserts one new quote into the db
func (ms *SQLightStorage) Store(quote QuoteInfo) {
	db := connect()
	defer db.Close()

	tx, err := db.Begin()
	checkErr(err)

	stmt, err := tx.Prepare("INSERT INTO Quotes(Quote, Who, Location, Time) VALUES(?, ?, ?, ?)")
	checkErr(err)
	defer stmt.Close()

	_, err = stmt.Exec(quote.Quote, quote.Who, quote.Where, quote.When)
	checkErr(err)

	tx.Commit()
}

func connect() *sql.DB {
	// db, err := sql.Open("sqlite3", ":memory:")
	db, err := sql.Open("sqlite3", "./sqlite3.sql")
	checkErr(err)

	checkErr(db.Ping())

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS Quotes (
			Id INTEGER PRIMARY KEY,
			Quote text NOT NULL,
			Who text NOT NULL,
			Location text NOT NULL,
			Time datetime NOT NULL)`)
	checkErr(err)

	return db
}

func checkErr(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
