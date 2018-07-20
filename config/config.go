package config

import (
	"log"

	"github.com/BurntSushi/toml"
	"os"
)

func init() {
	Config.Read()
}

var Config borealConfig

// Represents database server and credentials
type borealConfig struct {
	Mongo   mongo
	Ws      ws
	General general
}

type mongo struct {
	EngineDB      string
	EngineServer  string
	AccountDB     string
	AccountServer string
}

type ws struct {
	HostAddress string
	ReadBuffer int
	WriteBuffer int
	EnforceOrigin bool
}

type general struct {
	LogName    string
	Debug      bool
	DeltaMicro int
}

// Read and parse the configuration file
func (c *borealConfig) Read() {
	confPath := c.ResourceRoot()+"boreal_config.toml"

	if _, err := toml.DecodeFile(confPath, &c); err != nil {
		log.Println("Error loading boreal_config.toml, check boreal_config.toml.example")
		log.Fatal(err)
	}
}

func (c *borealConfig) ResourceRoot() string {
	resourceRoot, ok := os.LookupEnv("BOREALROOT")
	if !ok {
		resourceRoot = "./"
	}

	return resourceRoot
}
