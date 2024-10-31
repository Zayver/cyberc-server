package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"github.com/zayver/cybercomplaint-server/config"
)

type JwtService struct{
	config config.ConfigHolder
}

func NewJwtService(conf config.ConfigHolder) JwtService{
	return JwtService{
		config: conf,
	}
}

func (s *JwtService) GenerateJWT(username string) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"iss": "cyberc",
		"sub": username,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(60 * time.Minute).Unix(),
	})
	return token.SignedString([]byte(s.config.JWTSignKey))
}

func (s *JwtService) CheckToken(tokenString string) (bool, error){
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.config.JWTSignKey), nil
	}, jwt.WithExpirationRequired())
	if err != nil {
		return false, err
	}
	return token.Valid, nil
}