package api

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" //hidden
)

const (
	host     = "localhost"
	port     = 5433
	user     = "arkdev"
	password = "password"
	dbname   = "ark_devnet"
)

//Datastore interface
type Datastore interface {
	GetTransactions() ([]*Transaction, error)
}

//NewDB opening a db connection
func NewDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	return db, nil
}

//GetTransactions returns TX from node database
func GetTransactions(db *sql.DB) ([]*Transaction, error) {
	rows, err := db.Query("SELECT * FROM TRANSACTIONS")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]*Transaction, 0)
	for rows.Next() {
		bk := new(Transaction)
		err := rows.Scan(&bk.ID, &bk.SenderID, &bk.Amount, &bk.Timestamp)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bks, nil
}
