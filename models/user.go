package models

type USER struct {
	UUID     uint64 `gorm:"column:UUID;primary_key:auto_increment;type:int(9)" json:"uuid"`
	MSISDN   uint64 `gorm:"column:MSISDN;type:bigint(50)" json:"msisdn"`
	NAME     string `gorm:"column:NAME;type:varchar(255)" json:"name"`
	USERNAME string `gorm:"column:USERNAME;type:varchar(255)" json:"username"`
	PASSWORD string `gorm:"column:PASSWORD;->;<-;not null" json:"-"`
	TOKEN    string `gorm:"-" json:"token,omitempty"`
}

type CredentialsLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterUser struct {
	MSISDN   uint64 `gorm:"column:MSISDN;type:bigint(50)" json:"msisdn"`
	NAME     string `gorm:"column:NAME;type:varchar(255)" json:"name"`
	USERNAME string `gorm:"column:USERNAME;type:varchar(255)" json:"username"`
	PASSWORD string `gorm:"column:PASSWORD;->;<-;not null" json:"password"`
}
