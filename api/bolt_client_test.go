package api

import (
	"log"
	"testing"
)

func TestReadFromBucket(t *testing.T) {
	DBClient = &BoltClient{}
	DBClient.OpenBoltDb()
	DBClient.InitializeBucket()

	res, err := DBClient.QueryBlock("10048800703989711791")

	if err != nil {
		log.Println(err.Error())
	}

	log.Println(t.Name(), "Found block:", res)

}
