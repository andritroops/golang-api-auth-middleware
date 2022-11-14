package config

import (
	"fmt"
	"log"

	"github.com/andritroops/go-latihan/models/entity"
)

func RunMigration() {

	err := DB.AutoMigrate(&entity.User{})

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database migrated.")
}
