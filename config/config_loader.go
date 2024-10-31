package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ConfigHolder struct{
	DBURL string `mapstructure:"db_url"`
	Port string `mapstructure:"port"`
	JWTSignKey string `mapstructure:"jwt_sign_key"`
	Env string `mapstructure:"env"`
}

func NewConfigHolder () ConfigHolder{
	var config ConfigHolder

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.SetEnvPrefix("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil{
		log.Fatal("Unable to read config.yaml from filesystem: ", err)
	}

	if err:= viper.Unmarshal(&config); err != nil{
		log.Fatal("Unable to unmarshall config to struct ", err)
	}
	log.Info("Running in mode ", config.Env)
	return config
}