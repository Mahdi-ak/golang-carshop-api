package models

type Role struct {
	BaseModel
	Name      string `gorm:"notnull;size:10;type:string;unique"`
	UserRoles *[]UserRole
}
