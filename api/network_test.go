package api

import (
	"log"
	"testing"
)

func TestNewArkClient(t *testing.T) {
	arkapi := NewArkClient(nil)

	if arkapi == nil {
		t.Error("Error creating client")
	}
	log.Println(t.Name(), "Success")
}
