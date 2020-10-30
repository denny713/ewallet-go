package model

type BalanceBank struct {
	Id             int     `json:"id" gorm:"type:numeric;PRIMARY_KEY"`
	Balance        float64 `json:"balance" gorm:"type:numeric"`
	BalanceAchieve float64 `json:"balanceAchieve" gorm:"type:numeric"`
	Code           string  `json:"code" gorm:"type:varchar(25)"`
	Enable         bool    `json:"enable" gorm:"type:bool"`
	UserId         int     `json:"userId" gorm:"type:numeric"`
}

func (balanceBank *BalanceBank) TableName() string {
	return "balance_bank"
}
