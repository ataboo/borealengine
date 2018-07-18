package main

import (
	"github.com/ataboo/borealengine/services"
	"log"
	"github.com/op/go-logging"
	"github.com/ataboo/borealengine/config"
	"fmt"
)

var engine *services.Engine

func main() {
	fmt.Printf("\n%+v",config.Config)

	logLevel := logging.INFO
	if config.Config.General.Debug {
		logLevel = logging.DEBUG
	}
	logging.SetLevel(logLevel, config.Config.General.LogName)

	engine = services.NewEngine()

	stop, err := engine.Start()
	defer stop()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Started Boreal Engine.  Use ctrl-c to stop.")

	for {
		select {
		default:
		// Stop with ctrl-c for now
		}
	}
}