package repository_test

import (
	"bizCard/repository"
	"log"
	"testing"
)

func TestOpenDB(t *testing.T) {
	data := repository.OpenDB()
	log.Println(data)
	data = repository.OpenDB()
	log.Println(data)
}
