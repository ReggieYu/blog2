package database

import (
	"blog/config"
	"blog/model"
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init(cfg *config.Config) error {
	db, err := ConnectAndMigrate(cfg)
	if err != nil {
		return err
	}

	DB = db
	log.Printf("database connection success: %s", cfg.DBDriver)
	return nil
}

func ConnectAndMigrate(cfg *config.Config) (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
	)

	if cfg.DBDriver == "sqlite" {
		sqliteConfig := &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info), //显示sql日志
		}
		db, err = gorm.Open(sqlite.Open(cfg.SQLitePath), sqliteConfig)
		if err != nil {
			return nil, fmt.Errorf("sqlite connnect failed: %v", err)
		}

		log.Printf("sqllite connect success: %s", cfg.SQLitePath)
	} else if cfg.DBDriver == "mysql" {
		if cfg.MySQLDSN == "" {
			return nil, errors.New("mysqldsn is empty")
		}

		db, err = gorm.Open(mysql.Open(cfg.MySQLDSN), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("mysql connect failed: %v", err)
		}

		log.Printf("mysql connect success: %s", cfg.MySQLDSN)
	} else {
		return nil, fmt.Errorf("unsupported db driver: %s", cfg.DBDriver)
	}

	// ConnectAndMigrate
	log.Printf("\nstart to migrate....")
	if err := db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{}); err != nil {
		return nil, fmt.Errorf("migrate database failed: %v", err)
	}

	log.Printf("\nmigrate database success")
	return db, nil
}
