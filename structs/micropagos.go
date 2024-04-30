package structs

import (
	"encoding/json"
	"fmt"
)

type MicropagosConfiguration struct {
	URLSMSGateway    string         `json:"url_smsgateway"`
	UTLPortabilidad  string         `json:"utl_portabilidad"`
	MaxMessageLength int            `json:"max_message_length"`
	MobilRegex       string         `json:"mobil_regex"`
	Raven            []RavenService `json:"raven"`
}

type RavenService struct {
	Name        string `json:"name"`
	SendMail    bool   `json:"sendMail"`
	ShortNumber int    `json:"shortNumber"`
}

func (m MicropagosConfiguration) ConfigurationString() string {
	configJSON, err := json.Marshal(m)
	if err != nil {
		return fmt.Sprintf("Error al convertir la configuración a JSON: %v", err)
	}
	return string(configJSON)
}
