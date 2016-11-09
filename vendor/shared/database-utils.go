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
	flagType = flag.String("database-type", "",
		"type of db (ex: bolt or mysql)")
	flagName          = flag.String("database-name", "", "name of database")
	flagBoltDirectory = flag.String("database-bolt-directory", "",
		"path of database file for BoltDB")
	flagURL      = flag.String("database-url", "", "url of database")
	flagUsername = flag.String("database-username", "", "login of database")
	flagPassword = flag.String("database-password", "",
		"password of database if needed")
	flagPort      = flag.Int("database-port", -1, "port used by database")
	flagParameter = flag.String("database-mysql-parameter", "",
		"parameter for mysql database")
)

// TODO: replace defaults returns with your defaults configurations

// GetDatabaseType return the database type.
func GetDatabaseType(d DatabaseInfo) string {
	if *flagType != "" {
		return *flagType
	} else if os.Getenv("DATABASE_TYPE") != "" {
		return os.Getenv("DATABASE_TYPE")
	} else if d.Type != "" {
		return d.Type
	}
	return "Bolt"
}

// GetDatabaseDirectory return the database path for BoltDB.
func GetDatabaseDirectory(d DatabaseInfo) string {
	if *flagBoltDirectory != "" {
		return *flagBoltDirectory
	} else if os.Getenv("DATABASE_BOLTDIRECTORY") != "" {
		return os.Getenv("DATABASE_BOLTDIRECTORY")
	} else if d.BoltDirectory != "" {
		return d.BoltDirectory
	}
	return "./"
}

// GetDatabaseName return the database name.
func GetDatabaseName(d DatabaseInfo) string {
	if *flagName != "" {
		return *flagName
	} else if os.Getenv("DATABASE_NAME") != "" {
		return os.Getenv("DATABASE_NAME")
	} else if d.Name != "" {
		return d.Name
	}
	return "go-microservice"
}

// GetDatabaseUsername return the username for login to database.
func GetDatabaseUsername(d DatabaseInfo) string {
	if *flagUsername != "" {
		return *flagUsername
	} else if os.Getenv("DATABASE_USERNAME") != "" {
		return os.Getenv("DATABASE_USERNAME")
	} else if d.Username != "" {
		return d.Username
	}
	return ""
}

// GetDatabasePassword return the password for login to database.
func GetDatabasePassword(d DatabaseInfo) string {
	if *flagPassword != "" {
		return *flagPassword
	} else if os.Getenv("DATABASE_PASSWORD") != "" {
		return os.Getenv("DATABASE_PASSWORD")
	} else if d.Password != "" {
		return d.Password
	}
	return ""
}

// GetDatabaseURL return database URL.
func GetDatabaseURL(d DatabaseInfo) string {
	if *flagURL != "" {
		return *flagURL
	} else if os.Getenv("DATABASE_URL") != "" {
		return os.Getenv("DATABASE_URL")
	} else if d.URL != "" {
		return d.URL
	}
	return "127.0.0.1"
}

// GetDatabasePort return the database port if you used an URL.
func GetDatabasePort(d DatabaseInfo) int {
	value, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if *flagPort != -1 {
		return *flagPort
	} else if err == nil {
		return value
	} else if d.Port != -1 {
		return d.Port
	}
	return 8080
}

// GetDatabaseParameters return a string of all parameters for mysql link.
func GetDatabaseParameters(d DatabaseInfo) string {
	if *flagParameter != "" {
		return *flagParameter
	} else if os.Getenv("DATABASE_PARAMETER") != "" {
		return os.Getenv("DATABASE_PARAMETER")
	} else if d.Parameter != "" {
		return d.Parameter
	}
	return "?parseTime=true"
}
