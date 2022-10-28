package main

import (
	"fmt"

	"github.com/rasyidev/dts-h8-assignment-ii/database"
	"github.com/rasyidev/dts-h8-assignment-ii/models/domain"
)

func main() {
	database.GetDB()
}

func createItem(item domain.Item){
	db := database.GetDB()

	Item := domain.Item{
		Model: ,
	}
}
