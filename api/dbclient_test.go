package api

import (
	"log"
	"testing"

	"github.com/asdine/storm/q"

	"github.com/asdine/storm"
	"github.com/kristjank/goark-node/api"
	"github.com/kristjank/goark-node/api/model"
	"github.com/spf13/viper"
)

func initDB() {
	var err error
	api.ArkNodeDB, err = storm.Open(viper.GetString("db.filename"))

	if err != nil {
		log.Error(err.Error())
		panic(err.Error())
	}

	log.Println("Storm DB Opened at:", api.ArkNodeDB.Path)
}

func TestReadFromBucket(t *testing.T) {

	initDB()

	err = db.Select(q.Eq("ID")).Find(&users)

	res, err := api.ArkNodeDB.Find("12345679")

	if err != nil {
		log.Println(err.Error())
	}

	log.Println(t.Name(), "Found block:", res)

	DBClient.Close()
}

func TestAllBlock(t *testing.T) {
	initDB()

	res, err := DBClient.LastBlock()

	if err != nil {
		log.Println(err.Error())
	}

	log.Println(t.Name(), "Found last block:", res)
	DBClient.Close()
}

func TestSaveBlock(t *testing.T) {
	initDB()

	block2Save := model.Block{ID: "12345679", GeneratorID: "kristjan"}

	err := DBClient.SaveBlock(block2Save)

	block2Save.NumberOfTransactions = 10
	err = DBClient.SaveBlock(block2Save)

	block2Save.NumberOfTransactions = 22
	err = DBClient.SaveBlock(block2Save)

	if err != nil {
		t.Errorf(err.Error())
	}

	DBClient.Close()
}

func seedData() {
	initDB()

	DBClient.SaveBlock(model.Block{ID: "123", GeneratorID: "kristjan"})
	DBClient.SaveBlock(model.Block{ID: "124", GeneratorID: "kristjan1"})
	DBClient.SaveBlock(model.Block{ID: "125", GeneratorID: "nina1"})
	DBClient.SaveBlock(model.Block{ID: "125", PreviousBlock: "124"})
	DBClient.SaveBlock(model.Block{ID: "125", PreviousBlock: "124", GeneratorID: "lovro"})

	DBClient.Close()
}

func TestBoltClient_AllBlocks(t *testing.T) {
	seedData()
	initDB()

	res, err := DBClient.GetAllBlocks()

	if err != nil {
		t.Errorf(err.Error())
	}

	for _, element := range res {
		log.Println(element)
	}
	DBClient.Close()
}

func TestBoltClient_GetAllTx(t *testing.T) {

	initDB()

	res, err := DBClient.GetAllTransactions()

	if err != nil {
		t.Errorf(err.Error())
	}

	for _, element := range res {
		log.Println(element)
	}
	DBClient.Close()
}
