package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	db  *gorm.DB
	err error
	dsn = "doadmin:AVNS_H2AHwsJhR9To34OEkeD@tcp(dbaas-db-4459780-do-user-12872560-0.b.db.ondigitalocean.com:25060)/mini_project"
)

func ConnectDB() {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxOpenConns(150)
	if err != nil {
		panic(err)
	}

	log.Println("Connected to database")
}

func GetDB() *gorm.DB {
	return db
}
