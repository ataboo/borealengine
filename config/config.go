package config

import (
	"log"

	"github.com/BurntSushi/toml"
	"fmt"
)

func init() {
	Config.Read()
}

var Config BorealConfig

// Represents database server and credentials
type BorealConfig struct {
	Mongo mongo
	Ws ws
	General general
}

type mongo struct {
	DB   string
	Server string
}

type ws struct {
	HostAddress string
}

type general struct {
	LogName string
	Debug bool
	DeltaMicro int
}

// Read and parse the configuration file
func (c *BorealConfig) Read() {
	confPath := "boreal_config.toml"
	//flaggy.String(&confPath, "b", "bconf", "Location of a boreal_config.toml")
	//flaggy.Parse()

	if _, err := toml.DecodeFile(confPath, &c); err != nil {
		log.Println("Error loading boreal_config.toml, check boreal_config.toml.example")
		log.Fatal(err)
	}

	fmt.Printf("\nLoaded config: %+v", c)
}
