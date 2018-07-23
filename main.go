package main

import (
	"github.com/ataboo/borealengine/services"
	"log"
	"github.com/ataboo/borealengine/config"
	"fmt"
)

var engine *services.Engine

func main() {
	fmt.Printf("\n%+v",config.Config)

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