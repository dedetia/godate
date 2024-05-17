package config

import "github.com/spf13/viper"

type (
	MainConfig struct {
		Server       Server   `mapstructure:"server"`
		Database     Database `mapstructure:"database"`
		BaseUrlPhoto string   `mapstructure:"base_url_photo"`
		PrivateKey   string   `mapstructure:"private_key"`
	}

	Server struct {
		AppName      string `mapstructure:"app_name"`
		Port         int    `mapstructure:"port"`
		ReadTimeout  int    `mapstructure:"read_timeout"`
		WriteTimeout int    `mapstructure:"write_timeout"`
	}

	Database struct {
		Mongo Mongo `mapstructure:"mongo"`
	}

	Mongo struct {
		URI string `mapstructure:"uri"`
		DB  string `mapstructure:"db"`
	}
)

func Load() (*MainConfig, error) {
	v := viper.New()
	v.SetConfigFile("config/config.json")
	viper.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg MainConfig
	err := v.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
