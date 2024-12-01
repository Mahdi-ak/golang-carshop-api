package migrations

import (
	"github.com/Mahdi-ak/golang-carshop-api/src/config"
	"github.com/Mahdi-ak/golang-carshop-api/src/data/db"
	"github.com/Mahdi-ak/golang-carshop-api/src/data/models"
	"github.com/Mahdi-ak/golang-carshop-api/src/pkg/logging"
)

var logger = logging.NewLogger(config.GetConfig())

func Up_1() {
	database := db.GetDb()
	tables := []interface{}{}

	country := models.Country{}
	city := models.City{}

	if !database.Migrator().HasTable(country) {
		tables = append(tables, country)
	}

	if !database.Migrator().HasTable(city) {
		tables = append(tables, city)
	}
	database.Migrator().CreateTable(tables...)
	logger.Info(logging.Postgres, logging.Migration, "Tables created", nil)
}

func Down_1() {

}
