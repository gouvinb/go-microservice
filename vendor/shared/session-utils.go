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
	flagEnableSession = flag.Bool("session-enable", false, "enable session")
	flagSecretKey     = flag.String("session-secret-key", "",
		"secret key for session")
	flagSessionName = flag.String("session-name", "", "name of session")
	flagOptionPath  = flag.String("session-option-path", "",
		"path used for session")
	flagOptionDomain = flag.String("session-option-domain", "",
		"domain for session")
	flagOptionMaxAge = flag.Int("session-option-max-age", -1, "age of session")
	flagOptionSecure = flag.Bool("session-option-secure", false,
		"enable secure session")
	flagOptionHTTPOnly = flag.Bool("session-option-disable-http-only", true,
		"disable HTTP only")
)

// TODO: replace defaults returns with your defaults configurations

// IsSessionEnabled return true if use session.
func IsSessionEnabled(s Session) bool {
	value, err := strconv.ParseBool(os.Getenv("SESSION_ENABLE"))
	if *flagEnableSession != false {
		return *flagEnableSession
	} else if err == nil {
		return value
	} else if s.EnableSession != true {
		return s.EnableSession
	}
	return true
}

// GetSessionSecretKey return the secret key of session.
func GetSessionSecretKey(s Session) string {
	if *flagSecretKey != "" {
		return *flagSecretKey
	} else if os.Getenv("SESSION_SECRET_KEY") != "" {
		return os.Getenv("SESSION_SECRET_KEY")
	} else if s.SecretKey != "" {
		return s.SecretKey
	}
	return "go-microservice-secret-key-default"
}

// GetSessionName return the session name.
func GetSessionName(s Session) string {
	if *flagSessionName != "" {
		return *flagSessionName
	} else if os.Getenv("SESSION_NAME") != "" {
		return os.Getenv("SESSION_NAME")
	} else if s.Name != "" {
		return s.Name
	}
	return "go-microservice-default"
}

// GetSessionOptionPath return the session option path.
func GetSessionOptionPath(s Session) string {
	if *flagOptionPath != "" {
		return *flagOptionPath
	} else if os.Getenv("SESSION_OPTION_PATH") != "" {
		return os.Getenv("SESSION_OPTION_PATH")
	} else if s.Options.Path != "" {
		return s.Options.Path
	}
	return "/"
}

// GetSessionOptionDomain return the session option domain.
func GetSessionOptionDomain(s Session) string {
	if *flagOptionDomain != "" {
		return *flagOptionDomain
	} else if os.Getenv("SESSION_OPTION_DOMAIN") != "" {
		return os.Getenv("SESSION_OPTION_DOMAIN")
	} else if s.Options.Domain != "" {
		return s.Options.Domain
	}
	return ""
}

// GetSessionOptionMaxAge return the session max age.
func GetSessionOptionMaxAge(s Session) int {
	value, err := strconv.Atoi(os.Getenv("SESSION_OPTION_MAXAGE"))
	if *flagOptionMaxAge != -1 {
		return *flagOptionMaxAge
	} else if err == nil {
		return value
	} else if s.Options.MaxAge != -1 {
		return s.Options.MaxAge
	}
	return 28800
}

// GetSessionOptionSecure return the session secure status.
func GetSessionOptionSecure(s Session) bool {
	value, err := strconv.ParseBool(os.Getenv("SESSION_OPTION_SECURE"))
	if *flagOptionSecure != false {
		return *flagOptionSecure
	} else if err == nil {
		return value
	} else if s.Options.Secure != false {
		return s.Options.Secure
	}
	return false
}

// GetSessionOptionHTTPOnly return true if you use http only.
func GetSessionOptionHTTPOnly(s Session) bool {
	value, err := strconv.ParseBool(os.Getenv("SESSION_OPTION_HTTP_ONLY"))
	if *flagOptionHTTPOnly != true {
		return *flagOptionHTTPOnly
	} else if err == nil {
		return value
	} else if s.Options.HttpOnly != true {
		return s.Options.HttpOnly
	}
	return true
}
