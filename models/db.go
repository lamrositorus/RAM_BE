package models

import (
	"RAM/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg config.Config) (*gorm.DB, error) {
	var dsn string
	if cfg.UseURL {
		dsn = cfg.DatabaseURL
	} else {
		dsn = fmt.Sprintf(
    "host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
    cfg.DBHost, cfg.DBUser , cfg.DBPassword, cfg.DBName, cfg.DBPort,
)

	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("❌ Failed to connect DB: %v", err)
		return nil, err
	}

	// Auto migrate your models here
	err = db.AutoMigrate(&User {}, &Keuangan{}, &EstimasiKeuntungan{}, &EstimasiModalRequest{}, &SusutTimbangan{})
	if err != nil {
		return nil, err
	}

	DB = db
	log.Println("✅ Connected to PostgreSQL via GORM")
	return db, nil
}

func CloseDB() {
	sqlDB, err := DB.DB()
	if err == nil {
		sqlDB.Close()
	}
}
