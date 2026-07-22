package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

// var DbPath string
// type Storage struct {
// 	Driver string `mapstructure:"driver"`
// 	Path string `mapst`
// }

func SetupConfig() error{
	viper.SetDefault("logging.gorm_log_level", "warn")

	viper.SetConfigFile("taskman.yaml")
	viper.SetConfigName("taskman")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		var notFound viper.ConfigFileNotFoundError
		if !errors.As(err, &notFound) {
			return fmt.Errorf("reading config: %w", err)
		}
		// No config file is fine — defaults + env + flags still apply.
	}
	return nil
}
