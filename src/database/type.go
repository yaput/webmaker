package database

import "database/sql"

// Connection Store connection DB
type Connection struct {
	Db  *sql.DB
	cfg Config
}

// Config holds the configuration used for instantiating a new Connection.
type Config struct {
	// Address that locates our postgres instance
	Host string
	// Port to connect to
	Port string
	// User that has access to the database
	User string
	// Password so that the user can login
	Password string
	// Database to connect to (must have been created priorly)
	Database string
}

// SelectQuery handle query
type SelectQuery struct {
	Field     []string
	Table     string
	queryExec string
	conn      *sql.DB
}
