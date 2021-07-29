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
	Errors
	Additions
}

type Answers struct {
	StartWithGroup       string `mapstructure:"start_with_group"`
	StartWithoutGroup    string `mapstructure:"start_without_group"`
	ChangeGroup          string `mapstructure:"change_group"`
	Back                 string `mapstructure:"back"`
	ScheduleWithoutGroup string `mapstructure:"schedule_without_group"`
	ScheduleWithGroup    string `mapstructure:"schedule_with_group"`
	Thanks               string `mapstructure:"thanks"`
}

type Additions struct {
	ThanksNotSubscribed   string `mapstructure:"thanks_not_subscribed"`
	ScheduleNotSubscribed string `mapstructure:"schedule_not_subscribed"`
	ChangesInKeiSchedule  string `mapstructure:"changes_in_kei_schedule"`
}

type Errors struct {
	IncorrectGroupName   string `mapstructure:"incorrect_group_name"`
	GroupNotSelected     string `mapstructure:"group_not_selected"`
	StudentNotSubscribed string `mapstructure:"student_not_subscribed"`
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

	err = viper.UnmarshalKey("messages.additions", &cfg.Messages.Additions)
	if err != nil {
		return err
	}

	err = viper.UnmarshalKey("messages.errors", &cfg.Messages.Errors)
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
