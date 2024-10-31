package config

import (
	"github.com/zayver/cybercomplaint-server/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	log "github.com/sirupsen/logrus"
)

type DB struct{
	DB *gorm.DB
}

func NewDatabase(config ConfigHolder) DB{
	db := connectDB(config.DBURL)
	return DB{
		DB: db,
	}
}
func connectDB(url string) *gorm.DB{
	var err error

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database. Error: ", err)
	}
	log.Info("Connected to db", url)
	db.AutoMigrate(&model.Complaint{})
	db.AutoMigrate(&model.User{})
	return db
}
