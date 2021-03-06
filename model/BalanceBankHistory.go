package model

type BalanceBankHistory struct {
	Id            int     `json:"id" gorm:"type:numeric"`
	BalanceBankId int     `json:"balanceBankId" gorm:"type:numeric"`
	BalanceBefore float64 `json:"balanceBefore" gorm:"type:numeric"`
	BalanceAfter  float64 `json:"balanceAfter" gorm:"type:numeric"`
	Activity      string  `json:"activity" gorm:"type:varchar(10)"`
	Type          string  `json:"type" gorm:"type:varchar(6)"`
	Ip            string  `json:"type" gorm:"type:varchar(20)"`
	Location      string  `json:"location" gorm:"type:varchar(25)"`
	UserAgent     string  `json:"userAgent" gorm:"type:varchar(25)"`
	Author        string  `json:"author" gorm:"type:varchar(10)"`
	IdHistory     string  `json:"idHistory" gorm:"type:varchar(10)"`
}

func (balanceBankHistory *BalanceBankHistory) TableName() string {
	return "balance_bank_history"
}
