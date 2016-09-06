package shared

import (
	"os"
	"strconv"
)

func GetSessionSecretKey(s Session) string {
	if os.Getenv("SESSION_SECRET_KEY") != "" {
		return os.Getenv("SESSION_SECRET_KEY")
	} else if s.SecretKey != "" {
		return s.SecretKey
	} else {
		return "go-microservice-secret-key-default"
	}

}
func GetSessionName(s Session) string {
	if os.Getenv("SESSION_NAME") != "" {
		return os.Getenv("SESSION_NAME")
	} else if s.Name != "" {
		return s.Name
	} else {
		return "go-microservice-default"
	}
}

func GetSessionOptionPath(s Session) string {
	if os.Getenv("SessionOptionPath") != "" {
		return os.Getenv("SessionOptionPath")
	} else if s.Options.Path != "" {
		return s.Options.Path
	} else {
		return "/"
	}
}

func GetSessionOptionDomain(s Session) string {
	if os.Getenv("SessionOptionDomain") != "" {
		return os.Getenv("SessionOptionDomain")
	} else if s.Options.Domain != "" {
		return s.Options.Domain
	} else {
		return ""
	}
}

func GetSessionOptionMaxAge(s Session) int {
	if value, err := strconv.Atoi(os.Getenv("SessionOptionMaxAge")); err == nil {
		return value
	} else if s.Options.MaxAge != -1 {
		return s.Options.MaxAge
	} else {
		// TODO: replace this if you want replace by a default name
		return 28800
	}
}

func GetSessionOptionSecure(s Session) bool {
	if value, err := strconv.ParseBool(os.Getenv("SessionOptionSecure")); err == nil {
		return value
	} else if s.Options.Secure != false {
		return s.Options.Secure
	} else {
		// TODO: replace this if you want replace by a default name
		return false
	}
}

func GetSessionOptionHttpOnly(s Session) bool {
	if value, err := strconv.ParseBool(os.Getenv("SessionOptionHttpOnly")); err == nil {
		return value
	} else if s.Options.HttpOnly != true {
		return s.Options.HttpOnly
	} else {
		// TODO: replace this if you want replace by a default name
		return true
	}
}
