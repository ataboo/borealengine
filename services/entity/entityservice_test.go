package entity

import (
	"testing"
)

func TestLoadYamlResource(t *testing.T) {
	bear := SpeciesConfigs["bear"]

	t.Logf("Marshaled: \n%s", string(marshaled))
}