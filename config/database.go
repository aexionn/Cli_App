package config

import (
	"fmt"
	"os"
	"strings"

	// "github.com/aexionn/Cli_App/database/model"
	"github.com/glebarez/sqlite"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func parseLogLevel(s string) (logger.LogLevel, error) {
	switch strings.ToLower(s) {
	case "silent":
		return logger.Silent, nil
	case "error":
		return logger.Error, nil
	case "warn":
		return logger.Warn, nil
	case "info":
		return logger.Info, nil
	default:
		return 0, fmt.Errorf("invalid gorm_log_level %q: valid options are silent, error, warn, info", s)
	}
}

func Connection() {
	var err error

	// Read and validate the log level from config.
	level, levelErr := parseLogLevel(viper.GetString("logging.gorm_log_level"))
	if levelErr != nil {
		fmt.Println(levelErr)
		os.Exit(1)
	}

	DB, err = gorm.Open(sqlite.Open(viper.GetString("storage.path")), &gorm.Config{
		Logger: logger.Default.LogMode(level),
	})

	if err != nil {
		panic("Cannot connect to database")
	}
}