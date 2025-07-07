package main

import (
	"github.com/Andressatass/trabalho-microinformatica-be/configuration"
	"github.com/Andressatass/trabalho-microinformatica-be/db/repository"
	"github.com/Andressatass/trabalho-microinformatica-be/handlers"
)

func main() {
	config, err := configuration.GetConfig()
	if err != nil {
		return
	}

	if config.API.Port == "" {
		config.API.Port = ":8080"
	}

	// conn, err := settings.Connect(config.Database)
	// if err != nil {
	// 	return
	// }

	//userRepository := repository.NewUserRepository(conn.DB)
	repo := repository.NewMockUserRepository()

	handlers.RegisterRoutes(repo)

	err = handlers.StartServer(config.API.Port)
	if err != nil {
		return
	}

}
