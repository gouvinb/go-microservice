// Copyright 2016 gouvinb. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.md file.

// Package config load a default config file throught generate go-bindata.
package config

import (
	"encoding/json"
	"log"

	"shared"
)

// Configuration contains the application settings.
type Configuration struct {
	Database shared.DatabaseInfo `json:"Database"`
	Cors     shared.Cors         `json:"Cors"`
	Server   shared.Server       `json:"Server"`
	Session  shared.Session      `json:"Session"`
}

// NewConfiguration export interface.
func NewConfiguration() *Configuration {
	log.Println("New configuration instanciate")
	return &Configuration{}
}

// ParseJSON unmarshals bytes to structs.
func (c *Configuration) ParseJSON(b []byte) error {
	log.Println("Unmarshals bytes to structs")
	return json.Unmarshal(b, &c)
}

// Parser must implement ParseJSON.
type Parser interface {
	ParseJSON([]byte) error
}

// Load the JSON config file.
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
