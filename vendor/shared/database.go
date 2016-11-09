// Copyright 2016 gouvinb. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.md file.

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

// DatabaseInfo is the details for the database connection.
type DatabaseInfo struct {
	Type          string `json:"Type"`
	Name          string `json:"Name"`
	BoltDirectory string `json:"Bolt-directory"`
	URL           string `json:"URL"`
	Username      string `json:"Username"`
	Password      string `json:"Password"`
	Port          int    `json:"Port"`
	Parameter     string `json:"MySQL-parameter"`
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

// DatabaseConfigure to the database.
func DatabaseConfigure(d DatabaseInfo) {
	var err error

	// Store the config
	databases := d

	switch GetDatabaseType(databases) {
	case TypeMySQL:
		// Connect to MySQL
		if SQL, err = sqlx.Connect("mysql", DatabaseDNS(databases)); err != nil {
			log.Fatalln("SQL Driver Error", err)
		}
		// Check if is alive
		if err = SQL.Ping(); err != nil {
			log.Fatalln("Database Error", err)
		}
	case TypeBolt:
		// Connect to Bolt
		BoltDB, err = bolt.Open(GetDatabaseDirectory(databases)+
			GetDatabaseName(databases)+".db", 0600, nil)
		if err != nil {
			log.Fatalln("Bolt Driver Error", err)
		}
	case TypeMongoDB:
		// Connect to MongoDB
		Mongo, err = mgo.DialWithTimeout(GetDatabaseURL(databases), 5*time.Second)
		if err != nil {
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

// DatabaseDNS returns the Data Source Name.
func DatabaseDNS(d DatabaseInfo) string {
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

// DatabaseBoltUpdate makes a modification to Bolt.
func DatabaseBoltUpdate(bucketName string, key string,
	dataStruct interface{}) error {
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

// DatabaseBoltView retrieves a record in Bolt.
func DatabaseBoltView(bucketName string, key string,
	dataStruct interface{}) error {
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

// DatabaseBoltDelete removes a record from Bolt.
func DatabaseBoltDelete(bucketName string, key string) error {
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

// DatabaseMongoCheckConnection returns true if MongoDB is available.
func DatabaseMongoCheckConnection(databases DatabaseInfo) bool {
	if Mongo == nil {
		DatabaseConfigure(databases)
	}

	if Mongo != nil {
		return true
	}

	return false
}
