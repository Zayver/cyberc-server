package config

import (
	"crypto/rand"
	"errors"
	"math/big"

	log "github.com/sirupsen/logrus"
	"github.com/zayver/cyberc-server/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	log.Info("Connected to db")
	//err = db.AutoMigrate(&model.Complaint{}, &model.User{})
	if err != nil{
		log.Fatal("Failed to migrate models to db: ", err)
	}
	initDB(db)
	return db
}

func initDB(db *gorm.DB){
	var users = []model.User{}
	if err := db.First(&users).Error; err != nil{
		if !errors.Is(err, gorm.ErrRecordNotFound){
			log.Fatal("Error initializing db: ", err)
		}
	}
	if len(users) == 0{
		gen := generatePassword(10)
		log.Info("Creating user for first boot, pass: ", string(gen))
		pass, _ := bcrypt.GenerateFromPassword(gen, bcrypt.DefaultCost)
		user := model.User{
			Username: "malpica",
			Password: string(pass),
		}
		if err := db.Save(&user).Error; err != nil{
			log.Fatal("Error creating root user")
		}
	}
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	"0123456789" +
	"!@#$%^&*()-_=+"

func generatePassword(length int) []byte {
	password := make([]byte, length)
	for i := range password {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			log.Fatal("")
		}
		password[i] = charset[index.Int64()]
	}
	return password
}