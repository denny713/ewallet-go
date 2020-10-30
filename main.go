package main

import (
	"fmt"
	"github.com/denny713/ewallet-go/route"
	"github.com/denny713/ewallet-go/util"
	"github.com/spf13/viper"
	"strconv"
)

func main() {
	util.DatabaseConnect()
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error Reading Configuration File, %s", err)
	}
	r := route.RouteSetup()
	r.Run(":" + strconv.Itoa(viper.GetInt("server.port")))
}
