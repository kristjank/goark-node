package api

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"gopkg.in/gin-gonic/gin.v1"
)

const (
	host     = "localhost"
	port     = 5432
	user     = ""
	password = "password"
	dbname   = "ark_mainnet"
)

//GetPeers Returns a list of peers to client call. Response is in JSON
func GetPeers(c *gin.Context) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
