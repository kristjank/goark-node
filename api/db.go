package api

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" //hidden
	"github.com/spf13/viper"
)

var db *sql.DB

//InitDB opening a db connection
func InitDB(v *viper.Viper) {
	var err error
	log.Println(v.GetString("host"))
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", v.GetString("host"), v.GetInt("port"), v.GetString("user"), v.GetString("password"), v.GetString("database"))

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}
