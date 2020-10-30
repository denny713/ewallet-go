package service

import (
	"github.com/denny713/ewallet-go/config"
	"github.com/denny713/ewallet-go/model"
)

func GetAllUsers(users *[]model.Users) (err error) {
	if err := config.DB.Find(users).Error; err != nil {
		return err
	}
	return nil
}
