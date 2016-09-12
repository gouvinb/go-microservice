package shared

import (
	"flag"
	"os"
	"strconv"
)

var (
	flagHostname  = flag.String("server-hostname", "", "hostname of microservice")
	flagUseHTTP   = flag.Bool("server-use-http", true, "enable http protocol")
	flagUseHTTPS  = flag.Bool("server-use-https", false, "enable https protocol")
	flagHTTPPort  = flag.Int("server-http-port", -1, "http port of microservice")
	flagHTTPSPort = flag.Int("server-https-port", -1, "https port of microservice")
	flagCertFile  = flag.String("server-cert-file", "", "path of certificate file")
	flagKeyFile   = flag.String("server-key-file", "", "key of certificate")
)

func GetServerHostname(s Server) string {
	if *flagHostname != "" {
		return *flagHostname
	} else if os.Getenv("SERVER_HOSTNAME") != "" {
		return os.Getenv("SERVER_HOSTNAME")
	} else if s.Hostname != "" {
		return s.Hostname
	} else {
		// TODO: replace this if you want replace by a default name
		return ""
	}
}

func GetServerUseHTTP(s Server) bool {
	if *flagUseHTTP != true {
		return *flagUseHTTP
	} else if value, err := strconv.ParseBool(os.Getenv("SERVER_USE_HTTP")); err == nil {
		return value
	} else if s.UseHTTP != false {
		return s.UseHTTP
	} else {
		// TODO: replace this if you want replace by a default name
		return true
	}
}

func GetServerUseHTTPS(s Server) bool {
	if *flagUseHTTPS != false {
		return *flagUseHTTPS
	} else if value, err := strconv.ParseBool(os.Getenv("SERVER_USE_HTTPS")); err == nil {
		return value
	} else if s.UseHTTPS != false {
		return s.UseHTTPS
	} else {
		// TODO: replace this if you want replace by a default name
		return false
	}
}

func GetServerHTTPPort(s Server) int {
	if *flagHTTPPort != -1 {
		return *flagHTTPPort
	} else if value, err := strconv.Atoi(os.Getenv("SERVER_HTTP_PORT")); err == nil {
		return value
	} else if s.HTTPPort != -1 {
		return s.HTTPPort
	} else {
		// TODO: replace this if you want replace by a default name
		return 8000
	}
}

func GetServerHTTPSPort(s Server) int {
	if *flagHTTPSPort != -1 {
		return *flagHTTPSPort
	} else if value, err := strconv.Atoi(os.Getenv("SERVER_HTTPS_PORT")); err == nil {
		return value
	} else if s.HTTPSPort != -1 {
		return s.HTTPSPort
	} else {
		// TODO: replace this if you want replace by a default name
		return 443
	}
}

func GetServerCertFile(s Server) string {
	if *flagCertFile != "" {
		return *flagCertFile
	} else if os.Getenv("SERVER_CERT_FILE") != "" {
		return os.Getenv("SERVER_CERT_FILE")
	} else if s.CertFile != "" {
		return s.CertFile
	} else {
		// TODO: replace this if you want replace by a default name
		return ""
	}
}

func GetServerKeyFile(s Server) string {
	if *flagKeyFile != "" {
		return *flagKeyFile
	} else if os.Getenv("SERVER_KEY_FILE") != "" {
		return os.Getenv("SERVER_KEY_FILE")
	} else if s.KeyFile != "" {
		return s.KeyFile
	} else {
		// TODO: replace this if you want replace by a default name
		return ""
	}
}
