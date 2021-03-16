package config

import (
	"fmt"
	"github.com/denny713/ewallet-go/dto"
	"github.com/spf13/viper"
)

func ReadConfig() *dto.Config {
	data := dto.Config{}
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error Reading Configuration File, %s", err)
		return nil
	}
	database := dto.Database{
		Host: viper.GetString("datasource.host"),
		Port: viper.GetInt("datasource.port"),
		Name: viper.GetString("datasource.name"),
		User: viper.GetString("datasource.user"),
		Pass: viper.GetString("datasource.pass"),
	}
	data.Database = database
	server := dto.Server{
		Port: viper.GetInt64("server.port"),
	}
	data.Server = server
	return &data
}
