package log // import "github.com/Zenika/goru/log"

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func Init() error {
	logLevel := viper.GetString("logLevel")

	level, err := log.ParseLevel(logLevel)
	if err != nil {
		return errors.Wrap(err, "Error while parsing log level")
	}

	log.SetLevel(level)
	log.SetOutput(os.Stdout)

	log.Debug("Log level ", level)

	return nil
}
