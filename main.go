package main

import (
	"log"

	"github.com/scch94/Gconfiguration/config"
)

func main() {
	config, err := config.LoadConfiguration("config.json")
	if err != nil {
		log.Fatal("error al tratar de cargar la configuration ", err)
	}
	log.Printf("this is the name of the proyect: %v", config.Nombre)
}
