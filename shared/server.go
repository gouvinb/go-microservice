// Copyright 2016 gouvinb. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package shared

import (
	"fmt"
	"log"
	"net/http"
)

// Server stores the hostname and port number
type Server struct {
	Hostname  string `json:"Hostname"`  // Server name
	UseHTTP   bool   `json:"UseHTTP"`   // Listen on HTTP
	UseHTTPS  bool   `json:"UseHTTPS"`  // Listen on HTTPS
	HTTPPort  int    `json:"HTTPPort"`  // HTTP port
	HTTPSPort int    `json:"HTTPSPort"` // HTTPS port
	CertFile  string `json:"CertFile"`  // HTTPS certificate
	KeyFile   string `json:"KeyFile"`   // HTTPS private key
}

// Run starts the HTTP and/or HTTPS listener
func Run(httpHandlers http.Handler, httpsHandlers http.Handler, s Server) {
	if GetServerUseHTTP(s) && GetServerUseHTTPS(s) {
		log.Println("Start https and http server")
		go func() {
			startHTTPS(httpsHandlers, s)
		}()
		startHTTP(httpHandlers, s)
	} else if GetServerUseHTTP(s) {
		log.Println("Start http server")
		startHTTP(httpHandlers, s)
	} else if GetServerUseHTTPS(s) {
		log.Println("Start https server")
		startHTTPS(httpsHandlers, s)
	} else {
		log.Fatalln("Config file does not specify a listener to start")
	}
}

// startHTTP starts the HTTP listener
func startHTTP(handlers http.Handler, s Server) {
	log.Println("Running HTTP " + httpAddress(s))

	// Start the HTTP listener
	log.Fatalln(http.ListenAndServe(httpAddress(s), handlers))
}

// startHTTPs starts the HTTPS listener
func startHTTPS(handlers http.Handler, s Server) {
	log.Println("Running HTTPS " + httpsAddress(s))

	// Start the HTTPS listener
	log.Fatalln(http.ListenAndServeTLS(httpsAddress(s), GetServerCertFile(s), GetServerKeyFile(s), handlers))
}

// httpAddress returns the HTTP address
func httpAddress(s Server) string {
	return s.Hostname + ":" + fmt.Sprintf("%d", GetServerHTTPPort(s))
}

// httpsAddress returns the HTTPS address
func httpsAddress(s Server) string {
	return s.Hostname + ":" + fmt.Sprintf("%d", GetServerHTTPSPort(s))
}
