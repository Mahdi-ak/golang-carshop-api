package models

type City struct {
	BaseModel
	Name      string `gorm:"notnull;size:10;type:string;"`
	CountryId int
	Country   Country `gorm:"foreignKey:CountryId"`
}
