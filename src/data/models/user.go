package models

type User struct {
	BaseModel
	Username     string `gorm:"type:string;size:20;not null;unique"`
	FirstName    string `gorm:"null;size:15;type:string"`
	LastName     string `gorm:"null;size:25;type:string"`
	Email        string `gorm:"null;size:64;type:string;unique;default:null"`
	Password     string `gorm:"null;size:64;type:string;unique;default:null"`
	MobileNumber string `gorm:"notnull;size:11;type:string"`
	Enabled      bool   `gorm:"default:true"`
	UserRoles    *[]UserRole
}
