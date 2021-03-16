package controller

import (
	"github.com/denny713/ewallet-go/model"
	"github.com/denny713/ewallet-go/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var err error
var users []model.Users

func GetAllUsers(c *gin.Context) {
	err = service.GetAllUsers(&users)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, users)
	}
}
