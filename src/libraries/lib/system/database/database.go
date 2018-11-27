package database

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"libraries/lib/system/config"

	"github.com/boltdb/bolt"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/jmoiron/sqlx"
	mgo "gopkg.in/mgo.v2"
)

// DSN returns the Data Source Name
func DSN(ci config.MySQLInfo) string {
	// Example: root:@tcp(localhost:3306)/test
	return ci.Username +
		":" +
		ci.Password +
		"@tcp(" +
		ci.Hostname +
		":" +
		fmt.Sprintf("%d", ci.Port) +
		")/" +
		ci.Name + ci.Parameter
}

var databases config.DbInfo

// Connect to the database
func Connect() {
	fmt.Println(config.Cfg.Database)
	var d = config.Cfg.Database
	var err error

	// Store the config
	databases = d

	switch d.Type {
	case config.TypeMySQL:
		// Connect to MySQL
		if config.SQL, err = sqlx.Connect("mysql", DSN(d.MySQL)); err != nil {
			log.Println("SQL Driver Error", err)
		}

		// Check if is alive
		if err = config.SQL.Ping(); err != nil {
			log.Println("Database Error", err)
		}
	case config.TypeBolt:
		// Connect to Bolt
		if config.BoltDB, err = bolt.Open(d.Bolt.Path, 0600, nil); err != nil {
			log.Println("Bolt Driver Error", err)
		}
	case config.TypeMongoDB:
		// Connect to MongoDB
		if config.Mongo, err = mgo.DialWithTimeout(d.MongoDB.URL, 5*time.Second); err != nil {
			log.Println("MongoDB Driver Error", err)
			return
		}

		// Prevents these errors: read tcp 127.0.0.1:27017: i/o timeout
		config.Mongo.SetSocketTimeout(1 * time.Second)

		// Check if is alive
		if err = config.Mongo.Ping(); err != nil {
			log.Println("Database Error", err)
		}
	default:
		log.Println("No registered database in config")
	}
}

// Update makes a modification to Bolt
func Update(bucketName string, key string, dataStruct interface{}) error {
	err := config.BoltDB.Update(func(tx *bolt.Tx) error {
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
	err := config.BoltDB.View(func(tx *bolt.Tx) error {
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
	err := config.BoltDB.Update(func(tx *bolt.Tx) error {
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
func CheckConnection() bool {
	if config.Mongo == nil {
		Connect()
	}

	if config.Mongo != nil {
		return true
	}

	return false
}

// ReadConfig returns the database information
func ReadConfig() config.DbInfo {
	return databases
}
