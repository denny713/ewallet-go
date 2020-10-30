package route

import (
	"github.com/denny713/ewallet-go/controller"
	"github.com/gin-gonic/gin"
)

func RouteSetup() *gin.Engine {
	r := gin.Default()
	usr := r.Group("/user")
	{
		usr.GET("", controller.GetAllUsers)
	}
	return r
}
