// Copyright 2016 gouvinb. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package shared

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/jmoiron/sqlx"
	"gopkg.in/mgo.v2"
)

// DatabaseInfo is the details for the database connection
type DatabaseInfo struct {
	Type           string
	Name           string
	BoltDirectory  string
	URL            string
	Username       string
	Password       string
	Port           int
	MySQLParameter string
}

const (
	// TypeBolt is BoltDB
	TypeBolt = "Bolt"
	// TypeMongoDB is MongoDB
	TypeMongoDB = "MongoDB"
	// TypeMySQL is MySQL
	TypeMySQL = "MySQL"
)

var (
	// BoltDB wrapper
	BoltDB *bolt.DB
	// Mongo wrapper
	Mongo *mgo.Session
	// SQL wrapper
	SQL *sqlx.DB
)

// Connect to the database
func Connect(d DatabaseInfo) {
	var err error

	// Store the config
	databases := d

	switch GetDatabaseType(databases) {
	case TypeMySQL:
		// Connect to MySQL
		if SQL, err = sqlx.Connect("mysql", DSN(databases)); err != nil {
			log.Fatalln("SQL Driver Error", err)
		}
		// Check if is alive
		if err = SQL.Ping(); err != nil {
			log.Fatalln("Database Error", err)
		}
	case TypeBolt:
		// Connect to Bolt
		if BoltDB, err = bolt.Open(GetDatabaseDirectory(databases)+GetDatabaseName(databases)+".db", 0600, nil); err != nil {
			log.Fatalln("Bolt Driver Error", err)
		}
	case TypeMongoDB:
		// Connect to MongoDB
		if Mongo, err = mgo.DialWithTimeout(GetDatabaseURL(databases), 5*time.Second); err != nil {
			log.Fatalln("MongoDB Driver Error", err)
			return
		}
		// Prevents these errors: read tcp 127.0.0.1:27017: i/o timeout
		Mongo.SetSocketTimeout(1 * time.Second)
		// Check if is alive
		if err = Mongo.Ping(); err != nil {
			log.Fatalln("Database Error", err)
		}
	default:
		log.Fatalln("No registered database in config")
	}
}

// DSN returns the Data Source Name
func DSN(d DatabaseInfo) string {
	// Example: root:@tcp(localhost:3306)/test
	return GetDatabaseUsername(d) +
		":" +
		GetDatabasePassword(d) +
		"@tcp(" +
		GetDatabaseURL(d) +
		":" +
		fmt.Sprintf("%d", GetDatabasePort(d)) +
		")/" +
		GetDatabaseName(d) +
		GetDatabaseParameters(d)
}

// Update makes a modification to Bolt
func Update(bucketName string, key string, dataStruct interface{}) error {
	err := BoltDB.Update(func(tx *bolt.Tx) error {
		// Create the bucket
		bucket, e := tx.CreateBucketIfNotExists([]byte(bucketName))
		if e != nil {
			return e
		}

		// Encode the record
		encodedRecord, e := json.Marshal(dataStruct)
		if e != nil {
			return e
		}

		// Store the record
		if e = bucket.Put([]byte(key), encodedRecord); e != nil {
			return e
		}
		return nil
	})
	return err
}

// View retrieves a record in Bolt
func View(bucketName string, key string, dataStruct interface{}) error {
	err := BoltDB.View(func(tx *bolt.Tx) error {
		// Get the bucket
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return bolt.ErrBucketNotFound
		}

		// Retrieve the record
		v := b.Get([]byte(key))
		if len(v) < 1 {
			return bolt.ErrInvalid
		}

		// Decode the record
		e := json.Unmarshal(v, &dataStruct)
		if e != nil {
			return e
		}

		return nil
	})

	return err
}

// Delete removes a record from Bolt
func Delete(bucketName string, key string) error {
	err := BoltDB.Update(func(tx *bolt.Tx) error {
		// Get the bucket
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return bolt.ErrBucketNotFound
		}

		return b.Delete([]byte(key))
	})
	return err
}

// CheckConnection returns true if MongoDB is available
func CheckConnection(databases DatabaseInfo) bool {
	if Mongo == nil {
		Connect(databases)
	}

	if Mongo != nil {
		return true
	}

	return false
}
