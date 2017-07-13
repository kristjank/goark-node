package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/kristjank/goark-node/api/model"
)

//IBoltClient interface definition
type IBoltClient interface {
	OpenBoltDb()
	QueryBlock(blockID string) (model.Block, error)
	SaveBlock(block model.Block) error
	SaveTransaction(trans model.Transaction) (string, error)
	Check() bool
	InitializeBucket()
	LastBlock() (model.Block, error)
	GetAllBlocks() (results []model.Block, err error)
	GetAllTransactions() (results []model.Transaction, err error)
	Close()
}

var (
	ErrBucketNotFound = errors.New("Bucket not found")
	ErrKeyNotFound    = errors.New("Key not found")
	ErrDoesNotExist   = errors.New("Does not exist")
	ErrFoundIt        = errors.New("Found it")
	ErrExistsInSet    = errors.New("Element already exists in set")
	ErrInvalidID      = errors.New("Element ID can not contain \":\"")
)

//Constant names for BoltDb bucket initializations
const (
	BlockBucket       = "BlockBucket"
	TransactionBucket = "TransactionBucket"
)

//BoltClient Realimplementation
type BoltClient struct {
	boltDB *bolt.DB
}

//OpenBoltDb db opening
func (bc *BoltClient) OpenBoltDb() {
	var err error
	bc.boltDB, err = bolt.Open("ark-node.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Close the database
func (bc *BoltClient) Close() {
	bc.boltDB.Close()
}

//QueryBlock returns the block by id
func (bc *BoltClient) QueryBlock(blockID string) (model.Block, error) {
	// Allocate an empty Account instance we'll let json.Unmarhal populate for us in a bit.
	block := model.Block{}

	// Read an object from the bucket using boltDB.View
	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		// Read the bucket from the DB
		b := tx.Bucket([]byte(BlockBucket))

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

//Check Naive healthcheck, just makes sure the DB connection has been initialized.
func (bc *BoltClient) Check() bool {
	return bc.boltDB != nil
}

//InitializeBucket Creates an "BlockBucket" in our BoltDB. It will overwrite any existing bucket of the same name.
func (bc *BoltClient) InitializeBucket() {
	bc.boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(BlockBucket))
		if err != nil {
			return fmt.Errorf("create BlockBucket failed: %s", err)
		}

		_, err = tx.CreateBucket([]byte("TransactionBucket"))
		if err != nil {
			return fmt.Errorf("create TransactionBucket failed: %s", err)
		}

		return nil
	})
}

//SaveBlock to blockbucker
func (bc *BoltClient) SaveBlock(block model.Block) error {
	// Serialize the struct to JSON
	jsonBytes, _ := json.Marshal(block)

	// Write the data to the BlockBucketBlockBucket
	return bc.boltDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BlockBucket))
		err := b.Put([]byte(block.ID), jsonBytes)
		return err
	})
}

//SaveTransaction to TransactionBucket
func (bc *BoltClient) SaveTransaction(trans model.Transaction) (string, error) {
	// Serialize the struct to JSON
	jsonBytes, _ := json.Marshal(trans)

	// Write the data to the BlockBucketBlockBucket
	err := bc.boltDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TransactionBucket))
		err := b.Put([]byte(trans.ID), jsonBytes)
		return err
	})
	if err != nil {
		return "", err
	}
	return trans.ID, err

}

//LastBlock returns the last written block header
func (bc *BoltClient) LastBlock() (model.Block, error) {
	// Allocate an empty block instance we'll let json.Unmarhal populate for us in a bit.
	block := model.Block{}

	// Read an object from the bucket using boltDB.View
	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		// Read the buckDBClientet from the DB
		b := tx.Bucket([]byte(BlockBucket))

		_, val := b.Cursor().Last()
		// Read the value identified by our blockId supplied as []byte
		if val == nil {
			return fmt.Errorf("No last block found")
		}
		// Unmarshal the returned bytes into the block struct we created at
		// the top of the function
		json.Unmarshal(val, &block)

		// Return nil to indicate nothing went wrong, e.g no error
		return nil
	})
	// If there were an error, return the error
	if err != nil {
		return model.Block{}, err
	}
	// Return the model.Block struct and nil as error.
	return block, nil
}

//GetAllBlocks elements of a list
func (bc *BoltClient) GetAllBlocks() (results []model.Block, err error) {
	return results, bc.boltDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BlockBucket))
		if bucket == nil {
			return ErrBucketNotFound
		}
		return bucket.ForEach(func(_, value []byte) error {
			block := model.Block{}
			json.Unmarshal(value, &block)
			results = append(results, block)
			return nil // Continue ForEach
		})
	})
}

//GetAllTransactions elements of a list
func (bc *BoltClient) GetAllTransactions() (results []model.Transaction, err error) {
	return results, bc.boltDB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(TransactionBucket))
		if bucket == nil {
			return ErrBucketNotFound
		}
		return bucket.ForEach(func(_, value []byte) error {
			trans := model.Transaction{}
			json.Unmarshal(value, &trans)
			results = append(results, trans)
			return nil // Continue ForEach
		})
	})
}
