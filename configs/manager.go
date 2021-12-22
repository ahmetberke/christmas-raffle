package configs

import (
	"log"

	"github.com/spf13/viper"
)

var Manager manager

type manager struct {
	Participants []*participant
	SmtpSettings *smtpSettings
	MailSettings *mailSettings
}

type smtpSettings struct {
	Host     string
	Port     string
	Email    string
	Password string
}

type mailSettings struct {
	Subject      string
	BodyFilePath string
}

type participant struct {
	Name  string
	Email string
}

func ReadConfigFile(path string, filename string) {
	viper.SetConfigType("yaml")
	viper.SetConfigName(filename)
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading configuration file: %v", err)
	}
	if err := viper.Unmarshal(&Manager); err != nil {
		log.Fatalf("Error decoding configuration file: %v", err)
	}
}
