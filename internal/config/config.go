package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Token string

	Messages Messages
}

type Messages struct {
	Answers
}

type Answers struct {
	Start     string `mapstructure:"start"`
	HowAreYou string `mapstructure:"how_are_you"`
	Unknown   string `mapstructure:"unknown"`
}

func Init() (*Config, error) {
	err := setUpViper()
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	err = fromEnv(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func unmarshal(cfg *Config) error {
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return err
	}

	err = viper.UnmarshalKey("messages.answers", &cfg.Messages.Answers)
	if err != nil {
		return err
	}

	return nil
}

func fromEnv(cfg *Config) error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	cfg.Token = os.Getenv("TOKEN")
	return nil
}

func setUpViper() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	return viper.ReadInConfig()
}

//
//func fromEnvByViper(cfg *Config) error {
//	err := viper.BindEnv("token")
//	if err != nil {
//		return err
//	}
//
//	cfg.Token = viper.GetString("token")
//	return nil
//}
