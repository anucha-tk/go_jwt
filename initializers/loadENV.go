package initializers

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Port             string `mapstructure:"PORT"`
	DatabaseHost     string `mapstructure:"DATABASE_HOST"`
	DatabaseUsername string `mapstructure:"DATABASE_USERNAME"`
	DatabasePassword string `mapstructure:"DATABASE_PASSWORD"`
	DatabaseName     string `mapstructure:"DATABASE_NAME"`

	// JwtSecret    string        `mapstructure:"JWT_SECRET"`
	// JwtExpiresIn time.Duration `mapstructure:"JWT_EXPIRED_IN"`
	// JwtMaxAge    int           `mapstructure:"JWT_MAXAGE"`
	//
	// ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
}

func LoadENV() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
