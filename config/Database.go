package config

import (
	"fmt"
	"github.com/denny713/ewallet-go/dto"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Data = ReadConfig()

func BuildDatabase() *dto.Database {
	if Data != nil {
		database := dto.Database{}
		database = Data.Database
		return &database
	}
	return nil
}

func BuildUrlDb(dbConf *dto.Database) string {
	if dbConf != nil {
		return fmt.Sprintf(
			"postgres://%v:%v@%v:%v/%v?sslmode=disable",
			dbConf.User,
			dbConf.Pass,
			dbConf.Host,
			dbConf.Port,
			dbConf.Name,
		)
	}
	return "-"
}
