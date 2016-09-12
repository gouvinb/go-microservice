// Copyright 2016 gouvinb. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package shared

import (
	"flag"
	"os"
	"strconv"
)

var (
	flagType           = flag.String("database-type", "", "type of db (ex: bolt or mysql)")
	flagName           = flag.String("database-name", "", "name of database")
	flagBoltDirectory  = flag.String("database-bolt-directory", "", "path of database file for BoltDB")
	flagURL            = flag.String("database-url", "", "url of database")
	flagUsername       = flag.String("database-username", "", "login of database")
	flagPassword       = flag.String("database-password", "", "password of database if needed")
	flagPort           = flag.Int("database-port", -1, "port used by database")
	flagMySQLParameter = flag.String("database-mysql-parameter", "", "parameter for mysql database")
)

// TODO: replace defaults returns with your defaults configurations

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

func GetDatabasePort(d DatabaseInfo) int {
	if *flagPort != -1 {
		return *flagPort
	} else if value, err := strconv.Atoi(os.Getenv("DATABASE_PORT")); err == nil {
		return value
	} else if d.Port != -1 {
		return d.Port
	}
	return 8080
}

func GetDatabaseParameters(d DatabaseInfo) string {
	if *flagMySQLParameter != "" {
		return *flagMySQLParameter
	} else if os.Getenv("DATABASE_MYSQLPARAMETER") != "" {
		return os.Getenv("DATABASE_MYSQLPARAMETER")
	} else if d.MySQLParameter != "" {
		return d.MySQLParameter
	}
	return "?parseTime=true"
}
