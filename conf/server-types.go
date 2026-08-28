package conf

type ServerConfig struct {
	Generic  ServerGenericConfig  `mapstructure:"generic"`
	Instance ServerInstanceConfig `mapstructure:"instance"`
	Listen   ServerListenConfig   `mapstructure:"listen"`
}

type ServerGenericConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ServerInstanceConfig struct {
	Name          string `mapstructure:"name"`
	CaseSensitive bool   `mapstructure:"caseSensitive"`
	StrictRouting bool   `mapstructure:"strictRouting"`
}

type ServerListenConfig struct {
	EnablePrefork     bool `mapstructure:"enablePrefork"`
	EnablePrintRoutes bool `mapstructure:"enablePrintRoutes"`
}
