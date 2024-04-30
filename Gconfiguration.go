package Gconfiguration

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/scch94/Gconfiguration/constants"
)

// ConfigurationsModel es una interfaz para cualquier modelo de configuración.
type ConfigurationsModel interface {
	ConfigurationString() string
}

// Gconfiguration carga la configuración desde un archivo JSON y devuelve un modelo de configuración.
func GetConfig(configurationModel ConfigurationsModel) error {
	fmt.Println("Iniciando Gconfiguration. Versión:", version())

	// Obtener la ruta del ejecutable
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("error al obtener la ruta del ejecutable: %v", err)
	}
	exeDir := filepath.Dir(exePath)

	// Construir la ruta al archivo de configuración
	configFilePath := filepath.Join(exeDir, constants.CONFIG_FILE)
	fmt.Println("Ubicación del archivo de configuración:", configFilePath)

	// Leer el contenido del archivo de configuración
	configData, err := os.ReadFile(configFilePath)
	if err != nil {
		fmt.Println("error al leer el archivo de configuracion :", err)
		return fmt.Errorf("error al leer el archivo de configuración: %v", err)
	}

	// Decodificar el JSON en la estructura de configuración adecuada
	//var configurationStruct structs.MicropagosConfiguration
	err = json.Unmarshal(configData, configurationModel)
	if err != nil {
		return fmt.Errorf("el archivo config.json no coincide con la estructura %s, revisa el archivo nuevamente: %v", configurationModel, err)
	}

	fmt.Println("acabamos de obtener la configuracion")
	fmt.Println("este es el valor de la config", configurationModel.ConfigurationString())
	return nil
}

// readConfigFile lee el contenido del archivo de configuración.

// version devuelve la versión actual de la librería.
func version() string {
	return "1.0.0"
}
