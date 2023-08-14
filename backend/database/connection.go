package database

import (
	"fmt"
	"net/url"

	"invhub/config"
	"invhub/helper"
	"invhub/model"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB
var dbLogLevel logger.LogLevel

func MakeConnection() {
	var db *gorm.DB
	var err error

	connection := config.Database.Connection
	helper.Logger.Infof("Initializing database connection " + connection)

	if config.App.Env != "production" || config.App.Debug {
		dbLogLevel = logger.Info
	} else {
		dbLogLevel = logger.Warn
	}

	switch connection {
	case "postgres":
		db, err = pg()
	case "mysql":
		fallthrough
	default:
		db, err = my()
	}

	if err != nil {
		helper.Logger.Error(err)
		panic("DB Connection Failed.")
	}

	Db = db

	db.AutoMigrate(
		&model.User{},
		&model.AccessToken{},
		&model.RefreshToken{},
		&model.Product{},
		&model.Inventory{},
	)
}

func my() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name, url.QueryEscape(config.App.Timezone))
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(dbLogLevel),
	})
}

func pg() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", config.Database.Host, config.Database.User, config.Database.Password, config.Database.Name, config.Database.Port, config.App.Timezone)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(dbLogLevel),
	})
}
