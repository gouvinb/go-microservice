package shared

import (
	"os"
	"strconv"
)

func GetServerHostname(s Server) string {
	if os.Getenv("ServerHostname") != "" {
		return os.Getenv("ServerHostname")
	} else if s.Hostname != "" {
		return s.Hostname
	} else {
		// TODO: replace this if you want replace by a default name
		return ""
	}
}

func GetServerUseHTTP(s Server) bool {
	if value, err := strconv.ParseBool(os.Getenv("ServerUseHTTP")); err == nil {
		return value
	} else if s.UseHTTP != false {
		return s.UseHTTP
	} else {
		// TODO: replace this if you want replace by a default name
		return true
	}
}

func GetServerUseHTTPS(s Server) bool {
	if value, err := strconv.ParseBool(os.Getenv("ServerUseHTTPS")); err == nil {
		return value
	} else if s.UseHTTPS != false {
		return s.UseHTTPS
	} else {
		// TODO: replace this if you want replace by a default name
		return false
	}
}

func GetServerHTTPPort(s Server) int {
	if value, err := strconv.Atoi(os.Getenv("ServerHTTPPort")); err == nil {
		return value
	} else if s.HTTPPort != -1 {
		return s.HTTPPort
	} else {
		// TODO: replace this if you want replace by a default name
		return 8000
	}
}

func GetServerHTTPSPort(s Server) int {
	if value, err := strconv.Atoi(os.Getenv("ServerHTTPSPort")); err == nil {
		return value
	} else if s.HTTPSPort != -1 {
		return s.HTTPSPort
	} else {
		// TODO: replace this if you want replace by a default name
		return 443
	}
}

func GetServerCertFile(s Server) string {
	if os.Getenv("ServerCertFile") != "" {
		return os.Getenv("ServerCertFile")
	} else if s.CertFile != "" {
		return s.CertFile
	} else {
		// TODO: replace this if you want replace by a default name
		return ""
	}
}

func GetServerKeyFile(s Server) string {
	if os.Getenv("ServerKeyFile") != "" {
		return os.Getenv("ServerKeyFile")
	} else if s.KeyFile != "" {
		return s.KeyFile
	} else {
		// TODO: replace this if you want replace by a default name
		return ""
	}
}
