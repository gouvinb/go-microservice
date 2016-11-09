// Copyright 2016 gouvinb. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.md file.

// Package main is the bin for launch micro service.
package main

import (
	"flag"
	"log"

	"config"
	"plugin"
	"route"
	"shared"
)

//go:generate go-bindata -pkg config -o vendor/config/bindata.go vendor/config/...

func init() {
	// Verbose logging with file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	log.Println("Start microservice")

	// parse all flag for config
	flag.Parse()

	// config the settings variable
	var cfg = config.NewConfiguration()

	log.Println("Load the configuration file")
	configFile, err := config.Asset("vendor/config/config.json")
	if err != nil {
		log.Fatalln(err)
	} else {
		config.Load(configFile, "config.json", cfg)

		log.Println("Configure the session cookie store")
		shared.SessionConfigure(cfg.Session)

		log.Println("Configure cors")
		shared.CorsConfigure(cfg.Cors)

		log.Println("Configure and connect database")
		shared.DatabaseConfigure(cfg.Database)

		// Setup the views
		shared.ViewConfigure(cfg.View)
		shared.ViewLoadTemplates(shared.GetTemplateRoot(cfg.View.Template),
			shared.GetTemplateChildren(cfg.View.Template))
		shared.ViewLoadPlugins(plugin.TagHelper(cfg.View), plugin.NoEscape())

		log.Println("Start server")
		shared.ServerRun(route.LoadHTTP(), route.LoadHTTPS(), cfg.Server)
	}
}
