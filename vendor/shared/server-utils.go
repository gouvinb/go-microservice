// Copyright 2016 gouvinb. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.md file.

package shared

import (
	"flag"
	"os"
	"strconv"
)

var (
	flagHostname = flag.String("server-hostname", "", "hostname of microservice")
	flagUseHTTP  = flag.Bool("server-dont-use-http", true,
		"disable http protocol")
	flagUseHTTPS  = flag.Bool("server-use-https", false, "enable https protocol")
	flagHTTPPort  = flag.Int("server-http-port", -1, "http port of microservice")
	flagHTTPSPort = flag.Int("server-https-port", -1,
		"https port of microservice")
	flagCertFile = flag.String("server-cert-file", "",
		"path of certificate file")
	flagKeyFile = flag.String("server-key-file", "", "key of certificate")
)

// TODO: replace defaults returns with your defaults configurations

// GetServerHostname return the hostname.
func GetServerHostname(s Server) string {
	if *flagHostname != "" {
		return *flagHostname
	} else if os.Getenv("SERVER_HOSTNAME") != "" {
		return os.Getenv("SERVER_HOSTNAME")
	} else if s.Hostname != "" {
		return s.Hostname
	}
	return ""
}

// GetServerUseHTTP return the value for enable http or not.
func GetServerUseHTTP(s Server) bool {
	value, err := strconv.ParseBool(os.Getenv("SERVER_USE_HTTP"))
	if *flagUseHTTP != true {
		return *flagUseHTTP
	} else if err == nil {
		return value
	} else if s.UseHTTP != false {
		return s.UseHTTP
	}
	return true
}

// GetServerUseHTTPS return the value for enable https or not.
func GetServerUseHTTPS(s Server) bool {
	value, err := strconv.ParseBool(os.Getenv("SERVER_USE_HTTPS"))
	if *flagUseHTTPS != false {
		return *flagUseHTTPS
	} else if err == nil {
		return value
	} else if s.UseHTTPS != false {
		return s.UseHTTPS
	}
	return false
}

// GetServerHTTPPort return the http port used by go-microservice.
func GetServerHTTPPort(s Server) int {
	value, err := strconv.Atoi(os.Getenv("SERVER_HTTP_PORT"))
	if *flagHTTPPort != -1 {
		return *flagHTTPPort
	} else if err == nil {
		return value
	} else if s.HTTPPort != -1 {
		return s.HTTPPort
	}
	return 8000
}

// GetServerHTTPSPort return the https port used by go-microservice.
func GetServerHTTPSPort(s Server) int {
	value, err := strconv.Atoi(os.Getenv("SERVER_HTTPS_PORT"))
	if *flagHTTPSPort != -1 {
		return *flagHTTPSPort
	} else if err == nil {
		return value
	} else if s.HTTPSPort != -1 {
		return s.HTTPSPort
	}
	return 443
}

// GetServerCertFile return certificate file path.
func GetServerCertFile(s Server) string {
	if *flagCertFile != "" {
		return *flagCertFile
	} else if os.Getenv("SERVER_CERT_FILE") != "" {
		return os.Getenv("SERVER_CERT_FILE")
	} else if s.CertFile != "" {
		return s.CertFile
	}
	return ""
}

// GetServerKeyFile return certificate file password for use the certificate.
func GetServerKeyFile(s Server) string {
	if *flagKeyFile != "" {
		return *flagKeyFile
	} else if os.Getenv("SERVER_KEY_FILE") != "" {
		return os.Getenv("SERVER_KEY_FILE")
	} else if s.KeyFile != "" {
		return s.KeyFile
	}
	return ""
}
