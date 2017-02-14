package config // import "github.com/Zenika/goru/config"

import (
	log "github.com/Sirupsen/logrus"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var File string

func Init() error {
	viper.SetConfigName("goru")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if File != "" {
		viper.SetConfigFile(File)
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, notFound := err.(viper.ConfigFileNotFoundError); notFound {
			return nil
		}
		return errors.Wrap(err, "Error while reading config file")
	}

	log.Info("Using config file ", viper.ConfigFileUsed())

	return nil
}
