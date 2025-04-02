package config

import (
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	dsn := os.Getenv("DATABASE_URL")
	var err error
	Cfg.GormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		println("Error connecting to database: ", err)
	}
	Cfg.GormDB.Logger.LogMode(logger.Error)

	sqlDb, err := Cfg.GormDB.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDb.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDb.SetMaxOpenConns(30) // max in db is 83, leaving extras for other systems

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDb.SetConnMaxLifetime(time.Hour)

	if err != nil {
		println("Error connecting to database: ", err)
	}
}
