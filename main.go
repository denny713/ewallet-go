package main

import (
	"fmt"
	"github.com/denny713/ewallet-go/config"
	"github.com/denny713/ewallet-go/route"
	"strconv"
)

func main() {
	config.DatabaseConnect()
	server := config.Data.Server
	r := route.RouteSetup()
	fmt.Println("Starting Web Server In Port : " + strconv.Itoa(int(server.Port)))
	_ = r.Run(":" + strconv.Itoa(int(server.Port)))
}
