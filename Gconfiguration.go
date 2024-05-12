package Gconfiguration

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/scch94/Gconfiguration/constants"
	"github.com/scch94/ins_log"
)

//lint:ignore SA1029 "Using built-in type string as key for context value intentionally"
var ctx = context.WithValue(context.Background(), "packageName", "configuration")

// ConfigurationsModel es una interfaz para cualquier modelo de configuración.
type ConfigurationsModel interface {
	ConfigurationString() string
}

// Gconfiguration carga la configuración desde un archivo JSON y devuelve un modelo de configuración.
func GetConfig(configurationModel ConfigurationsModel) error {
	ins_log.SetService("Gconfiguration")
	ins_log.Infof(ctx, "starting gconfiguration Versión:", version())

	// Obtener la ruta del ejecutable
	exePath, err := os.Executable()
	if err != nil {
		ins_log.Fatalf(ctx, "error trying to get the file path  :", err)
		return fmt.Errorf("error trying to get the file path  : %v", err)
	}
	exeDir := filepath.Dir(exePath)

	// Construir la ruta al archivo de configuración
	configFilePath := filepath.Join(exeDir, constants.CONFIG_FILE)
	ins_log.Infof(ctx, "ubication of the configurationfile", configFilePath)

	// Leer el contenido del archivo de configuración
	configData, err := os.ReadFile(configFilePath)
	if err != nil {
		ins_log.Fatalf(ctx, "problem when we try to reed the configuration file :", err)
		return fmt.Errorf("problem when we try to reed the configuration file : %v", err)
	}

	// Decodificar el JSON en la estructura de configuración adecuada
	//var configurationStruct structs.MicropagosConfiguration
	err = json.Unmarshal(configData, configurationModel)
	if err != nil {
		ins_log.Fatalf(ctx, "the config.json file does not match the structure %s, please review the file again:", err)
		return fmt.Errorf("the config.json file does not match the structure %s, please review the file again: %v", configurationModel, err)
	}

	ins_log.Info(ctx, "ending to get the configuration")
	ins_log.Tracef(ctx, "this is the data that we have in the configuration", configurationModel.ConfigurationString())
	return nil
}

// readConfigFile lee el contenido del archivo de configuración.

// version devuelve la versión actual de la librería.
func version() string {
	return "1.0.0"
}
