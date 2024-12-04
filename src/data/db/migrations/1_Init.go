package migrations

import (
	"github.com/Mahdi-ak/golang-carshop-api/src/config"
	"github.com/Mahdi-ak/golang-carshop-api/src/constants"
	"github.com/Mahdi-ak/golang-carshop-api/src/data/db"
	"github.com/Mahdi-ak/golang-carshop-api/src/data/models"
	"github.com/Mahdi-ak/golang-carshop-api/src/pkg/logging"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var logger = logging.NewLogger(config.GetConfig())

// migrates the database to the initial state by creating tables and adding default information.
func Up_1() {
	// Get the database connection
	database := db.GetDb()

	// Create tables in the database
	createTables(database)

	// Add default information to the database
	createDefaultInformation(database)
}

// createTables creates tables in the database for the predefined models.
func createTables(database *gorm.DB) {
	// Initialize an empty list of tables
	tables := []interface{}{}

	// Define the models for which tables need to be created
	country := models.Country{}
	city := models.City{}
	user := models.User{}
	role := models.Role{}
	userRole := models.UserRole{}

	// Add each model to the list of tables to be created
	tables = addNewTable(database, country, tables)
	tables = addNewTable(database, city, tables)
	tables = addNewTable(database, user, tables)
	tables = addNewTable(database, role, tables)
	tables = addNewTable(database, userRole, tables)

	// Create tables in the database
	database.Migrator().CreateTable(tables...)
	logger.Info(logging.Postgres, logging.Migration, "Tables created", nil)
}

// addNewTable checks if a table for the given model exists in the database.
// If the table does not exist, it appends the model to the list of tables to be created.
// Parameters:
// - database: The database connection.
// - model: The model struct to check for table existence.
// - tables: The list of tables to be created.
// Returns:
// - []interface{}: The updated list of tables.
func addNewTable(database *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !database.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables
}

// createDefaultInformation creates the default roles in the database if they do not already exist.
func createDefaultInformation(database *gorm.DB) {

	adminRole := models.Role{Name: constants.AdminRoleName}
	createRoleIfNotExists(database, &adminRole)

	defaultRole := models.Role{Name: constants.AdminRoleName}
	createRoleIfNotExists(database, &defaultRole)

	u := models.User{Username: "admin", MobileNumber: "09141234567", Email: "Mahdiii.ak4@gmail.com"}
	pass := "12345678"
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	u.Password = string(hashPassword)
	createAdminUserIfNotExists(database, &u, adminRole.Id)

}

// createRoleIfNotExists checks if a role exists in the database, creates the role if it doesn't exist.
// Parameters:
// - database: The database connection.
// - r: The role to check and create if not exists.
func createRoleIfNotExists(database *gorm.DB, r *models.Role) {
	exists := 0
	database.
		Model(&models.Role{}).
		Select("1").
		Where("name = ?", r.Name).
		First(&exists)
	if exists == 0 {
		database.Create(r)
	}
}

func createAdminUserIfNotExists(database *gorm.DB, u *models.User, roleId int) {
	exists := 0
	database.
		Model(&models.User{}).
		Select("1").
		Where("username = ?", u.Username).
		First(&exists)
	if exists == 0 {
		database.Create(u)
		ur := models.UserRole{UserId: u.Id, RoleId: roleId}
		database.Create(&ur)
	}
}

func Down_1() {

}
