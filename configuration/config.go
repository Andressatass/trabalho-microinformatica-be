package configuration

import (
	"fmt"

	"github.com/Andressatass/trabalho-microinformatica-be/trabalho-microinformatica-be/app"
	"github.com/iamolegga/enviper"
	"github.com/spf13/viper"
)

type Configuration struct {
	Database Database `env:"MICROINFO_DATABASE"`
	API      API      `env:"MICROINFO_API"`
}

type Database struct {
	Port     string `env:"MICROINFO_DATABASE_PORT"`
	Host     string `env:"MICROINFO_DATABASE_HOST"`
	User     string `env:"MICROINFO_DATABASE_USER"`
	Password string `env:"MICROINFO_DATABASE_PASSWORD"`
	Name     string `env:"MICROINFO_DATABASE_NAME"`
}

type API struct {
	Port string `env:"MICROINFO_API_PORT"`
}

func GetConfig() (Configuration, error) {
	var config Configuration

	e := enviper.New(viper.New())

	e.SetEnvPrefix(app.EnvPrefix)
	e.SetConfigFile(app.ConfigFilePath)
	e.SetConfigName(app.ConfigName)

	err := e.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("could not unmarshal config into the struct: %s", err)
	}

	return config, nil
}
