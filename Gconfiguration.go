package Gconfiguration

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/scch94/ins_log"
)

//lint:ignore SA1029 "Using built-in type string as key for context value intentionally"
var ctx = context.WithValue(context.Background(), "packageName", "configuration")

// ConfigurationsModel es una interfaz para cualquier modelo de configuración.
type ConfigurationsModel interface {
	ConfigurationString() string
}

// Gconfiguration carga la configuración desde un archivo JSON y devuelve un modelo de configuración.
func GetConfig(configurationModel ConfigurationsModel, configPath string, fileName string) error {
	ins_log.SetService("Gconfiguration")
	ins_log.Infof(ctx, "starting gconfiguration Versión: %s", version())

	// Obtener la ruta del ejecutable
	// Construye la ruta completa al archivo de configuración
	configFilePath := filepath.Join(configPath, fileName)
	ins_log.Infof(ctx, "Location of the configuration file: %s", configFilePath)

	// Leer el contenido del archivo de configuración
	configData, err := os.ReadFile(configFilePath)
	if err != nil {
		ins_log.Fatalf(ctx, "error reading the configuration file: %v", err)
		return fmt.Errorf("error reading the configuration file: %v", err)
	}

	// Decodificar el JSON en la estructura de configuración adecuada
	//var configurationStruct structs.MicropagosConfiguration
	err = json.Unmarshal(configData, configurationModel)
	if err != nil {
		ins_log.Fatalf(ctx, "the config.json file does not match the structure %s, please review the file again:", err)
		return fmt.Errorf("the config.json file does not match the structure %s, please review the file again: %v", configurationModel, err)
	}

	ins_log.Info(ctx, "ending to get the configuration")
	ins_log.Tracef(ctx, "this is the data that we have in the configuration %s", configurationModel.ConfigurationString())
	return nil
}

// readConfigFile lee el contenido del archivo de configuración.

// version devuelve la versión actual de la librería.
func version() string {
	return "1.0.0"
}
