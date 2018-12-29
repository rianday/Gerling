package config

import (
	"github.com/boltdb/bolt"
	"github.com/jmoiron/sqlx"
	mgo "gopkg.in/mgo.v2"
)

var (
	// BoltDB wrapper
	BoltDB *bolt.DB
	// Mongo wrapper
	Mongo *mgo.Session
	// SQL wrapper
	SQL *sqlx.DB
	// Database info
	databases DbInfo
)

// Type is the type of database from a Type* constant
type Type string

const (
	// TypeBolt is BoltDB
	TypeBolt Type = "Bolt"
	// TypeMongoDB is MongoDB
	TypeMongoDB Type = "MongoDB"
	// TypeMySQL is MySQL
	TypeMySQL Type = "MySQL"
)

// Info contains the database configurations
type DbInfo struct {
	// Database type
	Type Type
	// MySQL info if used
	MySQL     MySQLInfo
	MySQLTest MySQLInfo
	// Bolt info if used
	Bolt BoltInfo
	// MongoDB info if used
	MongoDB MongoDBInfo
}

// MySQLInfo is the details for the database connection
type MySQLInfo struct {
	Username  string
	Password  string
	Name      string
	Hostname  string
	Port      int
	Parameter string
}

// // MySQLInfoTest is the details for the database connection
// type MySQLInfoTest struct {
// 	Username  string
// 	Password  string
// 	Name      string
// 	Hostname  string
// 	Port      int
// 	Parameter string
// }

// BoltInfo is the details for the database connection
type BoltInfo struct {
	Path string
}

// MongoDBInfo is the details for the database connection
type MongoDBInfo struct {
	URL      string
	Database string
}
