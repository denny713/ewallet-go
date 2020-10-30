package util

import (
	"fmt"
	"github.com/denny713/ewallet-go/config"
	"github.com/denny713/ewallet-go/model"
	"github.com/jinzhu/gorm"
)

var err error

func DatabaseConnect() {
	config.DB, err = gorm.Open("postgres", config.BuildUrlDb(config.BuildDatabase()))
	fmt.Println("Database Configuration : ", config.BuildUrlDb(config.BuildDatabase()))
	if err != nil {
		fmt.Println("Error : ", err)
	}
	config.DB.AutoMigrate(&model.Users{}, &model.BalanceBank{}, &model.UserBalance{}, &model.UserBalanceHistory{}, &model.BalanceBankHistory{})
}
