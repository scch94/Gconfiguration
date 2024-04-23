package config

import (
	"encoding/json"
	"log"
	"os"
)

func LoadConfiguration(filename string) (*MicropagosConfiguration, error) {
	configFile, err := os.Open(filename)
	if err != nil {
		log.Fatal("error al cargar la configuration")
		return nil, err
	}
	defer configFile.Close()
	var config MicropagosConfiguration
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		log.Fatal("error al parsear el json")
		return nil, err
	}

	return &config, nil
}
