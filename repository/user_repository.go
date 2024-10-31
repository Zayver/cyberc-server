package repository

import (
	"github.com/zayver/cybercomplaint-server/config"
	"github.com/zayver/cybercomplaint-server/model"
)

type UserRepository struct{
	db config.DB 
}

func NewUserRepository(db config.DB) UserRepository{
	return UserRepository{
		db: db,
	}
}

func(u *UserRepository) GetUserByUsername(username string) (model.User, error){
	var user model.User
	if err:= u.db.DB.First(&user, "username = ?", username).Error; err!= nil{
		return user, err
	}
	return user, nil
}

func(u *UserRepository) CreateUser(user model.User){
	if err:= u.db.DB.Save(&user).Error; err!= nil{
	}
}