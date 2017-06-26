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

var db *sql.DB

//InitDB opening a db connection
func InitDB() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}
