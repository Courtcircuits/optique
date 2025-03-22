package main

import (
	"log"
	"os"
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
	cycle := NewCycle()

	if conf.Bootstrap {
		err := cycle.Setup()
		if err != nil {
			log.Fatal(err)
			cycle.Stop()
			os.Exit(1)
		}
	}

	err = cycle.Ignite()
}
