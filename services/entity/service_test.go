package entity

import (
	"testing"
)

func TestLoadsSpeciesWithoutErrors(t *testing.T) {
	if errs := loadSpecies(); len(errs) > 0 {
		t.Error("%d errors thrown when loading species ymls", len(errs))
	}

	if len(SpeciesConfigs) == 0 {
		t.Error("no species configs were loaded")
	}
}

func TestLoadsNames(t *testing.T) {
	err := loadNames()
	if err != nil {
		t.Errorf("Error loading names: %s", err)
	}

	if len(Names) == 0 {
		t.Error("No names were loaded")
	}
}
