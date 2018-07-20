package entity

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"github.com/ataboo/borealengine/models"
	"github.com/ataboo/borealengine/config"
)

var SpeciesConfigs map[string]models.Species
var Names []string

type Types struct {
	TypeNames []string
}

func init() {
	loadSpecies()
	loadNames()
}

func loadSpecies() (errs []string) {
	errs = []string{}
	SpeciesConfigs = map[string]models.Species{}
	types := Types{}
	LoadYamlResource("species/types", &types)

	for _, name := range types.TypeNames {
		loadedSpecies := models.Species{}
		err := LoadYamlResource("species/"+name, &loadedSpecies)
		if err != nil {
			errs = append(errs, err.Error())
			continue
		}

		SpeciesConfigs[name] = loadedSpecies
	}

	return errs
}

func loadNames() error {
	err := LoadYamlResource("species/names", &Names)

	return err
}

func LoadYamlResource(filepath string, out interface{}) error {
	file, err := ioutil.ReadFile(config.Config.ResourceRoot()+"resources/"+filepath+".yml")
	if err != nil {
		log.Printf("\nfailed to load resource: %s with %s", filepath, err)
		return err
	}

	err = yaml.UnmarshalStrict(file, out)
	if err != nil {
		log.Printf("\nfailed to unmarshal file: %s with %s", filepath, err)
		return err
	}

	return nil
}

//func NewCreature(spawnPoint models.Transform) *models.Creature {
//
//}

