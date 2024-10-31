package service

import (
	log "github.com/sirupsen/logrus"
	"github.com/zayver/cybercomplaint-server/dto/request"
	"github.com/zayver/cybercomplaint-server/dto/response"
	"github.com/zayver/cybercomplaint-server/model"
	"github.com/zayver/cybercomplaint-server/repository"
	"golang.org/x/crypto/bcrypt"
)


type LoginService struct{
	userRepository repository.UserRepository
	jwtService JwtService
}

func NewLoginService(userRepo repository.UserRepository, jwt JwtService) LoginService{
	return LoginService{
		userRepository: userRepo,
		jwtService: jwt,
	}
}

func (l *LoginService) Login(request request.LoginRequest) (response.LoginResponse, error){
	user, err := l.userRepository.GetUserByUsername(request.Username)
	if err != nil{
		return response.LoginResponse{}, err
	}

	if !l.checkPasswordHash(request.Password, user.Password){
		return response.LoginResponse{}, bcrypt.ErrMismatchedHashAndPassword
	}
	token, err := l.jwtService.GenerateJWT(user.Username)
	if err != nil{
		log.Error("Error creating jwt: ", err)
		return response.LoginResponse{}, err
	}
	return response.LoginResponse{Token: token}, nil
}

func(l *LoginService) Signup(){
	pass, _ := bcrypt.GenerateFromPassword([]byte("pass"), bcrypt.DefaultCost)
	user := model.User{
		Username: "test",
		Password: string(pass),
	}
	l.userRepository.CreateUser(user)
}

func(l *LoginService) checkPasswordHash(password, hash string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return  err == nil
}