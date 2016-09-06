package shared

import (
	"os"
	"strconv"
)

func GetDatabaseType(d DatabaseInfo) string {
	if os.Getenv("DATABASE_TYPE") != "" {
		return os.Getenv("DATABASE_TYPE")
	} else if d.Type != "" {
		return d.Type
	} else {
		return "Bolt"
	}
}

func GetDatabaseDirectory(d DatabaseInfo) string {
	if os.Getenv("DATABASE_BOLTDIRECTORY") != "" {
		return os.Getenv("DATABASE_BOLTDIRECTORY")
	} else if d.BoltDirectory != "" {
		return d.BoltDirectory
	} else {
		// TODO: replace this if you want replace by a default BoltDirectory
		return "./"
	}
}

func GetDatabaseName(d DatabaseInfo) string {
	if os.Getenv("DATABASE_NAME") != "" {
		return os.Getenv("DATABASE_NAME")
	} else if d.Name != "" {
		return d.Name
	} else {
		// TODO: replace this if you want replace by a default name
		return "go-microservice"
	}
}

func GetDatabaseUsername(d DatabaseInfo) string {
	if os.Getenv("DATABASE_USERNAME") != "" {
		return os.Getenv("DATABASE_USERNAME")
	} else if d.Username != "" {
		return d.Username
	} else {
		// TODO: replace this if you want replace by a default Username
		return ""
	}
}

func GetDatabasePassword(d DatabaseInfo) string {
	if os.Getenv("DATABASE_PASSWORD") != "" {
		return os.Getenv("DATABASE_PASSWORD")
	} else if d.Password != "" {
		return d.Password
	} else {
		// TODO: replace this if you want replace by a default Password
		return ""
	}
}

func GetDatabaseURL(d DatabaseInfo) string {
	if os.Getenv("DATABASE_URL") != "" {
		return os.Getenv("DATABASE_URL")
	} else if d.URL != "" {
		return d.URL
	} else {
		// TODO: replace this if you want replace by a default URL
		return "127.0.0.1"
	}
}

func GetDatabasePort(d DatabaseInfo) int {
	if value, err := strconv.Atoi(os.Getenv("DATABASE_PORT")); err == nil {
		return value
	} else if d.Port != -1 {
		return d.Port
	} else {
		// TODO: replace this if you want replace by a default Port
		return 8080
	}
}

func GetDatabaseParameters(d DatabaseInfo) string {
	if os.Getenv("DATABASE_MYSQLPARAMETER") != "" {
		return os.Getenv("DATABASE_MYSQLPARAMETER")
	} else if d.MySQLParameter != "" {
		return d.MySQLParameter
	} else {
		// TODO: replace this if you want replace by a default name
		return "?parseTime=true"
	}
}
