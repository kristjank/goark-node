package api

import (
	"log"
	"testing"

	"github.com/asdine/storm"
	"github.com/kristjank/goark-node/api/model"
)

var testNodeDB *storm.DB

func initDB() {
	var err error
	testNodeDB, err = storm.Open("arktest.db")

	if err != nil {
		log.Fatal(err.Error())
		panic(err.Error())
	}

	log.Println("Storm DB Opened at:", testNodeDB.Path)
}

func TestSaveTx(t *testing.T) {

	initDB()

	var tx model.Transaction

	tx.ID = "123321blkm.,m12"
	tx.Amount = 1000022

	err := testNodeDB.Save(&tx)
	if err != nil {
		log.Fatal(t.Name(), err.Error())
	}

	testNodeDB.Close()
}

func TestListTx(t *testing.T) {
	initDB()

	var results []model.Transaction
	//var query storm.Query
	//	query = testNodeDB.All()

	err := testNodeDB.All(&results)
	if err != nil {
		log.Fatal(t.Name(), err.Error())
	}

	for id, element := range results {
		log.Println(element.ID, element.Amount, element.Timestamp, element.Blockid, id)
	}

	testNodeDB.Close()
}

func TestListBlocks(t *testing.T) {
	initDB()

	var results []model.Block
	err := testNodeDB.AllByIndex("Height", &results, storm.Reverse())

	if err != nil {
		log.Fatal(t.Name(), err.Error())
	}

	for id, element := range results {
		log.Println(element.ID, element.Height, id)
	}

	testNodeDB.Close()
}

func TestLastBlockByLimit(t *testing.T) {
	initDB()

	var results []model.Block
	var lastBlock model.Block
	err := testNodeDB.AllByIndex("Height", &results, storm.Limit(1), storm.Reverse())

	if err != nil {
		log.Fatal(t.Name(), err.Error())
	}

	/*for id, element := range results {
		log.Println(element.ID, element.Height, id)
		lastBlock = element
		break
	}*/

	lastBlock = results[0]
	log.Println(lastBlock)

	testNodeDB.Close()
}

func TestLastBlock(t *testing.T) {
	initDB()

	var lastBlock model.Block
	var query storm.Query
	var err error
	query = testNodeDB.Select().OrderBy("Height").Reverse()
	err = query.First(&lastBlock)
	if err != nil {
		log.Fatal(t.Name(), err.Error())
	}

	log.Println(lastBlock)

	testNodeDB.Close()
}

/*
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
*/
