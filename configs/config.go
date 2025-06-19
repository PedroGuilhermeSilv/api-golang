package configs

import (
	errors "github.com/PedroGuilhermeSilv/api-golang/pkg/erros"
	"github.com/spf13/viper"

	"github.com/go-chi/jwtauth"
)

var cfg *conf

type conf struct {
	DbDriver       string `mapstructure:"DB_DRIVER"`
	DbHost         string `mapstructure:"DB_HOST"`
	DbPort         string `mapstructure:"DB_PORT"`
	DbUser         string `mapstructure:"DB_USER"`
	DbPass         string `mapstructure:"DB_PASS"`
	DbName         string `mapstructure:"DB_NAME"`
	WebServerPort  string `mapstructure:"WEB_SERVER_PORT"`
	JwtSecretKey   string `mapstructure:"JWT_SECRET_KEY"`
	JwtExpiredTime int    `mapstructure:"JWT_EXPIRED_TIME"`
	TokenAuth      *jwtauth.JWTAuth
}

func LoadConfig(path string) (*conf, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return errors.HandleError(cfg, err)
	}

	err := viper.Unmarshal(&cfg)
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JwtSecretKey), nil)
	return errors.HandleError(cfg, err)
}
