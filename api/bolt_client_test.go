package api

import (
	"log"
	"testing"
)



func initDB() {
	DBClient = &BoltClient{}
	DBClient.OpenBoltDb()
	DBClient.InitializeBucket()
}

func TestReadFromBucket(t *testing.T) {

	/*initDB()


	res, err := DBClient.QueryBlock("10048800703989711791")

	if err != nil {
		log.Println(err.Error())
	}

	log.Println(t.Name(), "Found block:", res)*/

}

func TestLastBlock(t *testing.T) {
	initDB() 

	res, err := DBClient.LastBlock()

	if err != nil {
		log.Println(err.Error())
	}

	log.Println(t.Name(), "Found last block:", res)

}
