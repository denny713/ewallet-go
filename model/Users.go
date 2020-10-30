package model

type Users struct {
	Id       int    `json:"id" gorm:"type:numeric;PRIMARY_KEY"`
	Username string `json:"username" gorm:"type:varchar(15)"`
	Password string `json:"password" gorm:"type:varchar(64)"`
	Email    string `json:"email" gorm:"type:varchar(30)"`
	Status   string `json:"status" gorm:"type:varchar(5)"`
}

func (users *Users) TableName() string {
	return "users"
}
