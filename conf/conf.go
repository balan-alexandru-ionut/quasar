package conf

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig `mapstructure:"server"`
	Database Database     `mapstructure:"database"`
}

var C Config

// Load reads the configuration file and unmarshals its contents into the global Config object.
// It searches for the config file in the working directory and the user's home configuration path.
// Logs a panic and terminates if the config file cannot be read or if unmarshaling fails.
func Load() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.config/quasar")

	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("Unable to read config file: %v", err)
	}

	if err := viper.Unmarshal(&C); err != nil {
		log.Panicf("Unable to unmarshal config file: %v", err)
	}
}
