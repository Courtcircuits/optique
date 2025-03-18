package main

import (
	"github.com/Courtcircuits/optique/config"
	"github.com/Courtcircuits/optique/http"
)

// @title Optique application TO CHANGE
// @version 1.0
// @description This is a sample application
// @contact.name Courtcircuits
// @contact.url https://github.com/Courtcircuits
// @contact.email tristan-mihai.radulescu@etu.umontpellier.fr
func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		config.HandleError(err)
	}

	hello := http.NewHelloController()
	server := http.NewServer(conf.HTTP.ListenAddr)
	server.Setup(
		hello,
	).Ignite()
}
