package model

type UserBalance struct {
	Id             int     `json:"id" gorm:"type:numeric;PRIMARY_KEY"`
	UserId         int     `json:"userId" gorm:"type:numeric"`
	Balance        float64 `json:"balance" gorm:"type:numeric"`
	BalanceAchieve float64 `json:"balanceAchieve" gorm:"type:numeric"`
}

func (userBalance *UserBalance) TableName() string {
	return "user_balance"
}
