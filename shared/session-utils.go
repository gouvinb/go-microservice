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
	flagSecretKey      = flag.String("session-secret-key", "", "secret key for session")
	flagSessionName    = flag.String("session-name", "", "name of session")
	flagOptionPath     = flag.String("session-option-path", "", "path used for session")
	flagOptionDomain   = flag.String("session-option-domain", "", "domain for session")
	flagOptionMaxAge   = flag.Int("session-option-max-age", -1, "age of session")
	flagOptionSecure   = flag.Bool("session-option-secure", false, "enable secure session")
	flagOptionHTTPOnly = flag.Bool("session-option-disable-http-only", true, "disable HTTP only")
)

// TODO: replace defaults returns with your defaults configurations

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

func GetSessionOptionMaxAge(s Session) int {
	if *flagOptionMaxAge != -1 {
		return *flagOptionMaxAge
	} else if value, err := strconv.Atoi(os.Getenv("SESSION_OPTION_MAXAGE")); err == nil {
		return value
	} else if s.Options.MaxAge != -1 {
		return s.Options.MaxAge
	}
	return 28800
}

func GetSessionOptionSecure(s Session) bool {
	if *flagOptionSecure != false {
		return *flagOptionSecure
	} else if value, err := strconv.ParseBool(os.Getenv("SESSION_OPTION_SECURE")); err == nil {
		return value
	} else if s.Options.Secure != false {
		return s.Options.Secure
	}
	return false
}

func GetSessionOptionHttpOnly(s Session) bool {
	if *flagOptionHTTPOnly != true {
		return *flagOptionHTTPOnly
	} else if value, err := strconv.ParseBool(os.Getenv("SESSION_OPTION_HTTP_ONLY")); err == nil {
		return value
	} else if s.Options.HttpOnly != true {
		return s.Options.HttpOnly
	}
	return true
}
