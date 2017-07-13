package api

import (
	"log"
	"testing"

	"github.com/kristjank/goark-node/api/model"
)

func initDB() {
	DBClient = &BoltClient{}
	DBClient.OpenBoltDb()
	DBClient.InitializeBucket()
}

func TestReadFromBucket(t *testing.T) {

	initDB()

	res, err := DBClient.QueryBlock("12345679")

	if err != nil {
		log.Println(err.Error())
	}

	log.Println(t.Name(), "Found block:", res)

	DBClient.Close()
}

func TestLastBlock(t *testing.T) {
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
