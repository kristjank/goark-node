package api

import (
	"log"
	"os"
	"testing"

	"github.com/asdine/storm"
	"github.com/kristjank/goark-node/api/model"
)

var testNodeDB *storm.DB

func initDB() {
	var err error
	//opening DB from GO-ARK settings folder
	testNodeDB, err = storm.Open("../cmd/goark-node/db/ark-node.db")

	if err != nil {
		log.Fatal(err.Error())
		panic(err.Error())
	}

	log.Println("Storm DB Opened at:", testNodeDB.Path)
	log.SetOutput(os.Stdout)
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

	//err := testNodeDB.All(&results)
	err := testNodeDB.AllByIndex("PkIndex", &results, storm.Limit(50), storm.Skip(5), storm.Reverse())
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

	lastBlock = results[0]
	log.Println(lastBlock.Height)

	testNodeDB.Close()
}

func TestGetBlockByID(t *testing.T) {
	initDB()
	var block model.Block
	err := testNodeDB.One("ID", "1844069042066945391", &block)

	if err != nil {
		log.Fatal(t.Name(), err.Error())
	}

	log.Println(block)
}

func TestGetTransByID(t *testing.T) {
	initDB()
	var trans model.Transaction
	//44d9d0a3093232b9368a24af90577741df8340b93732db23b90d44f6590d3e42
	//cbd6862966bb1b03ba742397b7e5a88d6eefb393a362ead0d605723b840db2af
	//b2ef0adc90e3cf4af5d221350d79c2f2712378e0ef5a71244eecaca4afdc7140  -- negative timestamp... int representation
	err := testNodeDB.One("ID", "b2ef0adc90e3cf4af5d221350d79c2f2712378e0ef5a71244eecaca4afdc7140", &trans)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(trans)
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
