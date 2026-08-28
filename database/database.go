package database

import (
	"log"
	"quasar/conf"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var config = &conf.C

var DB *gorm.DB

func Connect() {
	dsn := "host=" + config.Database.Host +
		" user=" + config.Database.Username +
		" password=" + config.Database.Password +
		" dbname=" + config.Database.Name +
		" port=" + strconv.Itoa(config.Database.Port)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panicf("Cannot establish database connection: %v", err)
	}
}
