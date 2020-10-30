package config

import (
	"fmt"
	"github.com/denny713/ewallet-go/dto"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func BuildDatabase() *dto.Database {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error Reading Configuration File, %s", err)
	}
	dbConf := dto.Database{
		Host: viper.GetString("datasource.host"),
		Port: viper.GetInt("datasource.port"),
		Name: viper.GetString("datasource.name"),
		User: viper.GetString("datasource.username"),
		Pass: viper.GetString("datasource.password"),
	}
	return &dbConf
}

func BuildUrlDb(dbConf *dto.Database) string {
	return fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable",
		dbConf.User,
		dbConf.Pass,
		dbConf.Host,
		dbConf.Port,
		dbConf.Name,
	)
}
