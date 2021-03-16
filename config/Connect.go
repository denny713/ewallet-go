package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

func DatabaseConnect() {
	url := BuildUrlDb(BuildDatabase())
	if url != "-" {
		DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})
		fmt.Println("Database Configuration : ", url)
		if err != nil {
			fmt.Println("Error : ", err)
		}
	} else {
		fmt.Println("Error : Cannot Connect Database")
	}
}
