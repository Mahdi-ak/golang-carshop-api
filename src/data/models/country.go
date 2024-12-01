package models

type Country struct {
	BaseModel
	Name   string `gorm:"notnull;size:15;type:string;"`
	Cities *[]City
}
