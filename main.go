package main

import (
	"github.com/go-clinic/appointments/persistance"
	"github.com/go-clinic/common"
	_ "github.com/go-clinic/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api
func main() {
	config := ReadConfiguration()
	logger := CreateLogger(config)

	db := InitializeDatabase(config)
	defer db.Close()

	router := CreateGinServerInstance(logger)

	RegisterModules(router, config, db)

	router.Run(config.Server.FullAddress())
}

func InitializeDatabase(config common.Configuration) *gorm.DB {
	db, error := gorm.Open("postgres", config.Database.ConnectionString())

	if error != nil {
		panic(error.Error())
	}

	persistance.MigrateDatabase(db)

	return db
}
