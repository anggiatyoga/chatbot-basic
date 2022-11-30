package config

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	AppConfig AppConfig
	DB        MongoDB
	LogInfo   LogConfig
	LogError  LogConfig
}

type AppConfig struct {
	Address         string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
}

type MongoDB struct {
	UriServer      string
	AuthMechanism  string
	Username       string
	Password       string
	AuthSource     string
	CredentialName string
}

type LogConfig struct {
	Path   string
	Size   int
	Backup int
}

func GetConfig() (Config, error) {
	var conf Config
	viper.AddConfigPath(".")
	viper.SetConfigFile("app.env")

	if err := viper.ReadInConfig(); nil != err {
		log.Error().Str("config", "read config").Msg(err.Error())
		return conf, err
	}

	conf.AppConfig = AppConfig{
		Address:         viper.GetString("APP_ADDRESS"),
		WriteTimeout:    viper.GetDuration("APP_WRITE_TIMEOUT"),
		ReadTimeout:     viper.GetDuration("APP_READ_TIMEOUT"),
		ShutdownTimeout: viper.GetDuration("APP_SHUTDOWN_TIMEOUT"),
	}

	conf.DB = MongoDB{
		UriServer:      viper.GetString("URI_SERVER"),
		AuthMechanism:  viper.GetString("CREDENTIAL_AUTHMECHANISM"),
		Username:       viper.GetString("CREDENTIAL_USERNAME"),
		Password:       viper.GetString("CREDENTIAL_PASSWORD"),
		AuthSource:     viper.GetString("CREDENTIAL_AUTHSOURCE"),
		CredentialName: viper.GetString("CREDENTIAL_NAME"),
	}

	conf.LogInfo = LogConfig{
		Path:   viper.GetString("LOG_INFO_FILE"),
		Size:   viper.GetInt("LOG_INFO_SIZE"),
		Backup: viper.GetInt("LOG_INFO_BACKUP"),
	}

	conf.LogError = LogConfig{
		Path:   viper.GetString("LOG_ERROR_FILE"),
		Size:   viper.GetInt("LOG_ERROR_SIZE"),
		Backup: viper.GetInt("LOG_ERROR_BACKUP"),
	}

	return conf, nil
}

func (m *MongoDB) ToString() string {
	return fmt.Sprintf("URI_SERVER: %s\nCREDENTIAL_AUTHMECHANISM: %s\nCREDENTIAL_USERNAME: %s\nCREDENTIAL_PASSWORD: %s\nCREDENTIAL_AUTHSOURCE: %s\nCREDENTIAL_NAME: %s\n",
		m.UriServer, m.AuthMechanism, m.Username, m.Password, m.AuthSource, m.Username)
}
