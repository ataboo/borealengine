package entity

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"github.com/ataboo/borealengine/models"
	"github.com/ataboo/borealengine/config"
)

var SpeciesConfigs = map[string]models.Species{}

type Types struct {
	Names []string
}

func init() {
	loadSpecies()
}

func loadSpecies() {
	types := Types{}
	LoadYamlResource("species/types", &types)

	for _, name := range types.Names {
		loadedSpecies := models.Species{}
		err := LoadYamlResource("species/"+name, &loadedSpecies)
		if err == nil {
			SpeciesConfigs[name] = loadedSpecies
		}
	}
}

func LoadYamlResource(filepath string, out interface{}) error {
	file, err := ioutil.ReadFile(config.Config.ResourceRoot()+"resources/"+filepath+".yml")
	if err != nil {
		log.Printf("\nfailed to load resource: %s with %s", filepath, err)
		return err
	}

	err = yaml.Unmarshal(file, out)
	if err != nil {
		log.Printf("\nfailed to unmarshal file: %s with %s", filepath, err)
		return err
	}

	return nil
}