// Copyright 2016 gouvinb. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.md file.

package shared

import (
	"fmt"
	"log"
	"net/http"
)

// Server stores the hostname and port number.
type Server struct {
	Hostname  string `json:"Hostname"`  // Server name
	UseHTTP   bool   `json:"UseHTTP"`   // Listen on HTTP
	UseHTTPS  bool   `json:"UseHTTPS"`  // Listen on HTTPS
	HTTPPort  int    `json:"HTTPPort"`  // HTTP port
	HTTPSPort int    `json:"HTTPSPort"` // HTTPS port
	CertFile  string `json:"CertFile"`  // HTTPS certificate
	KeyFile   string `json:"KeyFile"`   // HTTPS private key
}

// ServerRun starts the HTTP and/or HTTPS listener.
func ServerRun(httpHandlers http.Handler, httpsHandlers http.Handler, s Server) {
	if GetServerUseHTTP(s) && GetServerUseHTTPS(s) {
		log.Println("Start https and http server")
		go func() {
			ServerStartHTTPS(httpsHandlers, s)
		}()
		ServerStartHTTP(httpHandlers, s)
	} else if GetServerUseHTTP(s) {
		log.Println("Start http server")
		ServerStartHTTP(httpHandlers, s)
	} else if GetServerUseHTTPS(s) {
		log.Println("Start https server")
		ServerStartHTTPS(httpsHandlers, s)
	} else {
		log.Fatalln("Config file does not specify a listener to start")
	}
}

// ServerStartHTTP starts the HTTP listener.
func ServerStartHTTP(handlers http.Handler, s Server) {
	log.Println("Running HTTP " + ServerHTTPAddress(s))

	// Start the HTTP listener
	log.Fatalln(http.ListenAndServe(ServerHTTPAddress(s), handlers))
}

// ServerStartHTTPS starts the HTTPS listener.
func ServerStartHTTPS(handlers http.Handler, s Server) {
	log.Println("Running HTTPS " + ServerHTTPSAddress(s))

	// Start the HTTPS listener
	log.Fatalln(http.ListenAndServeTLS(ServerHTTPSAddress(s), GetServerCertFile(s), GetServerKeyFile(s), handlers))
}

// ServerHTTPAddress returns the HTTP address.
func ServerHTTPAddress(s Server) string {
	return s.Hostname + ":" + fmt.Sprintf("%d", GetServerHTTPPort(s))
}

// ServerHTTPSAddress returns the HTTPS address.
func ServerHTTPSAddress(s Server) string {
	return s.Hostname + ":" + fmt.Sprintf("%d", GetServerHTTPSPort(s))
}
