package models

type LOGISTIC struct {
	LOGISTIC_ID      uint64 `gorm:"primary_key:auto_increment;column:LOGISTIC_ID;type:int(10);not null" json:"logistic_id"`
	LOGISTIC_NAME    string `gorm:"column:LOGISTIC_NAME;type:varchar(255);not null" json:"logistic_name"`
	AMOUNT           string `gorm:"column:AMOUNT;type:int(10)" json:"amount"`
	DESTINATION_NAME string `gorm:"column:DESTINATION_NAME;type:varchar(255)" json:"destination_name"`
	ORIGIN_NAME      string `gorm:"column:ORIGIN_NAME;type:varchar(255)" json:"origin_name"`
	DURATION         string `gorm:"column:DURATION;type:varchar(255)" json:"duration"`
	COMMENT          string `gorm:"column:COMMENT;type:varchar(255)" json:"comment"`
	ACTIVE           string `gorm:"column:ACTIVE;type:char(1)" json:"active"`
	CREATE_USER      uint64 `gorm:"column:CREATE_USER;type:int(9)" json:"create_user"`
	CREATE_DATE      string `gorm:"column:CREATE_DATE;type:datetime" json:"create_date"`
	UPDATE_USER      uint64 `gorm:"column:UPDATE_USER;type:int(9)" json:"update_user"`
	UPDATE_DATE      string `gorm:"column:UPDATE_DATE;type:datetime" json:"update_date"`
}
