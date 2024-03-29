package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type Database struct {
	dialector gorm.Dialector
}

func NewDatabse(dialector gorm.Dialector) *Database {
	return &Database{

		dialector: dialector,
	}
}

func (d *Database) Connect() {

	db, err := gorm.Open(d.dialector, &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic(err)
	}
	pgdb, err := db.DB()
	if err != nil {
		panic(err)
	}
	// Config connection pool
	pgdb.SetMaxIdleConns(10)
	pgdb.SetMaxOpenConns(100)
	// Set DB connection as global
	DB = db
}
