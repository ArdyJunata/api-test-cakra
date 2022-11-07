package database

import (
	"time"

	"github.com/ArdyJunata/api-test-cakra/config"
	"github.com/ArdyJunata/api-test-cakra/pkg/db"
	"github.com/ArdyJunata/api-test-cakra/server/models"
	"gorm.io/gorm"
)

func ConnectPostgres() *gorm.DB {
	gormPostgres := db.NewGormPostgres(
		config.GetString(config.POSTGRES_HOST),
		config.GetString(config.POSTGRES_PORT),
		config.GetString(config.POSTGRES_USER),
		config.GetString(config.POSTGRES_PASS),
		config.GetString(config.POSTGRES_DBNAME),
		config.GetString(config.POSTGRES_SSLMODE),
	)

	err := gormPostgres.Connect()
	if err != nil {
		panic(err)
	}

	err = gormPostgres.SetConnectionPool(
		int(config.GetInt8(config.POSTGRES_MAX_OPEN_CONNS)),
		int(config.GetInt8(config.POSTGRES_MAX_OPEN_CONNS)),
		time.Duration(config.GetInt8(config.POSTGRES_LIFETIME_IDLE_CONNS)),
		time.Duration(config.GetInt8(config.POSTGRES_LIFETIME_OPEN_CONNS)),
	)

	if err != nil {
		panic(err)
	}

	db := gormPostgres.(*db.GormPostgresDB).DB

	db.AutoMigrate(
		models.Car{},
		models.Club{},
	)

	return db
}
