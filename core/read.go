package core

import (
	"encoding/json"
	"os"
)

type OptiqueModuleManifest struct {
	Name   string   `json:"name"`
	Type   string   `json:"type"`
	URL    string   `json:"url"`
	Ignore []string `json:"ignore"`
}

func ReadManifest() (*OptiqueModuleManifest, error) {
	optiqueManifest, err := os.ReadFile(MODULE_MANIFEST)
	if err != nil {
		return nil, err
	}
	var manifest OptiqueModuleManifest
	err = json.Unmarshal(optiqueManifest, &manifest)
	if err != nil {
		return nil, err
	}
	return &manifest, nil
}

type OptiqueProjectManifest struct {
	Name string `json:"name"` //optique name
	Module string `json:"module"` // golang module name
	Repositories []string `json:"repositories"`
	Applications []string `json:"applications"`
	Ignore []string `json:"ignore"`
}

func ReadProjectManifest() (*OptiqueProjectManifest, error) {
	return ReadProjectManifestAt(PROJECT_MANIFEST)
}

func ReadProjectManifestAt(path string) (*OptiqueProjectManifest, error) {
	optiqueManifest, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var manifest OptiqueProjectManifest
	err = json.Unmarshal(optiqueManifest, &manifest)
	if err != nil {
		return nil, err
	}
	return &manifest, nil
}
