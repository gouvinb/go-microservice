package config

import (
	"encoding/json"
	"log"

	"github.com/gouvinb/go-microservice/shared"
)

// configuration contains the application settings
type configuration struct {
	Database shared.DatabaseInfo `json:"Database"`
	Server   shared.Server       `json:"Server"`
	Session  shared.Session      `json:"Session"`
}

// NewConfiguration export interface
func NewConfiguration() *configuration {
	log.Println("New configuration instanciate")
	return &configuration{}
}

// ParseJSON unmarshals bytes to structs
func (c *configuration) ParseJSON(b []byte) error {
	log.Println("Unmarshals bytes to structs")
	return json.Unmarshal(b, &c)
}

// Parser must implement ParseJSON
type Parser interface {
	ParseJSON([]byte) error
}

// Load the JSON config file
func Load(configFileByte []byte, configFile string, p Parser) {
	log.Println("Parse the config")
	err := p.ParseJSON(configFileByte)
	if err != nil {
		log.Println(configFileByte)
		log.Fatalln("Could not parse %q: %v", configFile, err)
	} else {
		log.Println("Configuration loaded")
	}
}
