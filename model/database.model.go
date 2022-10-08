package model

import (
	"eCommerce/config"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DBInfo struct {
	Host       string
	Port       string
	Name       string
	User       string
	Pass       string
	SearchPath string
	SSLMode    string
}

var (
	dbUrlFormat = "host=%s port=%s user=%s password=%s dbname=%s search_path=%s sslmode=%s"
	dbConfig    = config.DatabaseConfig{}
	_           = env.Parse(&dbConfig)
	dbInfo      = DBInfo{
		Host:       dbConfig.DBHost,
		Port:       dbConfig.DBPort,
		Name:       dbConfig.DBName,
		User:       dbConfig.DBUser,
		Pass:       dbConfig.DBPass,
		SearchPath: dbConfig.DBSchema,
		SSLMode:    dbConfig.SSLMode,
	}
	dbUrl = fmt.Sprintf(dbUrlFormat, dbInfo.Host, dbInfo.Port, dbInfo.User, dbInfo.Pass, dbInfo.Name, dbInfo.SearchPath, dbInfo.SSLMode)
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(dbConfig.DBType, dbUrl)

	// SetMaxIdleConn sets the maximum number of connections in the idle connection pool.
	db.DB().SetMaxIdleConns(2)

	// SetMaxOpenConn sets the maximum number of open connections to the database.
	db.DB().SetMaxOpenConns(2)
	AddUUIDGenerateExtension(db)
	return db, err
}

func AddUUIDGenerateExtension(db *gorm.DB) {
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
}
