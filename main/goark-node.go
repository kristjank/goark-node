package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gopkg.in/gin-gonic/gin.v1"
)

var errorlog *os.File
var logger *log.Logger
var router *gin.Engine
var env Env

type Env struct {
	db *sql.DB
}

//INIT PART
func init() {

	//initializeRoutes()

}

func initLogger() {
	errorlog, err := os.OpenFile("log/goark-node.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		os.Exit(1)
	}

	logger = log.New(errorlog, "ark-go: ", log.Lshortfile|log.LstdFlags)
}

///////////////////////////
func main() {
	logger.Println("GOArk-Node starting")
	initLogger()
	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Initialize the routes
	initializeRoutes()

	//environment init db and logger
	db, err := models.NewDB()
	if err != nil {
		log.Panic(err)
	}
	env = &Env{db: db}

	// Start serving the application
	router.Run("localhost:4002")

}
