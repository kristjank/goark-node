package api

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/kristjank/goark-node/api/model"
)

type IBoltClient interface {
	OpenBoltDb()
	QueryBlock(blockId string) (model.Block, error)
	SaveBlock(block model.Block)
	Check() bool
	InitializeBucket()
	LastBlock() (model.Block, error)
}

// Real implementation
type BoltClient struct {
	boltDB *bolt.DB
}

func (bc *BoltClient) OpenBoltDb() {
	var err error
	bc.boltDB, err = bolt.Open("goark-node.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (bc *BoltClient) QueryBlock(blockID string) (model.Block, error) {
	// Allocate an empty Account instance we'll let json.Unmarhal populate for us in a bit.
	block := model.Block{}

	// Read an object from the bucket using boltDB.View
	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		// Read the bucket from the DB
		b := tx.Bucket([]byte("BlockBucket"))

		// Read the value identified by our accountId supplied as []byte
		blockBytes := b.Get([]byte(blockID))
		if blockBytes == nil {
			return fmt.Errorf("No block found for " + blockID)
		}
		// Unmarshal the returned bytes into the account struct we created at
		// the top of the function
		json.Unmarshal(blockBytes, &block)

		// Return nil to indicate nothing went wrong, e.g no error
		return nil
	})
	// If there were an error, return the error
	if err != nil {
		return model.Block{}, err
	}
	// Return the Account struct and nil as error.
	return block, nil
}

// Naive healthcheck, just makes sure the DB connection has been initialized.
func (bc *BoltClient) Check() bool {
	return bc.boltDB != nil
}

// Creates an "AccountBucket" in our BoltDB. It will overwrite any existing bucket of the same name.
func (bc *BoltClient) InitializeBucket() {
	bc.boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("BlockBucket"))
		if err != nil {
			return fmt.Errorf("create bucket failed: %s", err)
		}
		return nil
	})
}

// Seed (n) make-believe account objects into the AcountBucket bucket.
func (bc *BoltClient) SaveBlock(block model.Block) {
	// Serialize the struct to JSON
	jsonBytes, _ := json.Marshal(block)

	// Write the data to the AccountBucket
	bc.boltDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("BlockBucket"))
		err := b.Put([]byte(block.ID), jsonBytes)
		return err
	})
}

func (bc *BoltClient) LastBlock() (model.Block, error) {
	// Allocate an empty Account instance we'll let json.Unmarhal populate for us in a bit.
	block := model.Block{}

	// Read an object from the bucket using boltDB.View
	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		// Read the bucket from the DB
		b := tx.Bucket([]byte("BlockBucket"))

		_, val := b.Cursor().Last()
		// Read the value identified by our accountId supplied as []byte
		if val == nil {
			return fmt.Errorf("No last block found...")
		}
		// Unmarshal the returned bytes into the account struct we created at
		// the top of the function
		json.Unmarshal(val, &block)

		// Return nil to indicate nothing went wrong, e.g no error
		return nil
	})
	// If there were an error, return the error
	if err != nil {
		return model.Block{}, err
	}
	// Return the Account struct and nil as error.
	return block, nil
}
