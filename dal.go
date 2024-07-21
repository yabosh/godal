package godal

/*
  This file contains a basic data access layer (Dal) that is a global value that can be used to
  reference database objects.  This can simplify application structure because the
  connection strings can be built at application startup and then repositories only need to
  know the name of the connection they wish to use without needing all of the details about the connection.
*/

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	// dal is a singleton that contains all of the named database connections available
	dal = &DataAccessLayer{}
)

func init() {
	dal = &DataAccessLayer{
		connMapString: make(map[string]*ConnInfo),
	}
}

// Dal provides access to the global data access layer connection map.
func Dal() *DataAccessLayer {
	return dal
}

// DataAccessLayer is a map of connection information that can be accessed by name
type DataAccessLayer struct {
	connMapString map[string]*ConnInfo
}

// ConnInfo is the metadata used to create a database connection
type ConnInfo struct {
	DBDriverName     string
	ConnectionString string
	db               *sqlx.DB
}

// ConnectionProvider represents a structure that provides connection details for a database platform.
type ConnectionProvider interface {
	GetDBURL() string
	GetDBDriver() string
}

// AddConnection will add a new, named connection to the data access layer.
func (d *DataAccessLayer) AddConnection(connName string, provider ConnectionProvider, maxConnections int) (err error) {
	connInfo := &ConnInfo{
		DBDriverName:     provider.GetDBDriver(),
		ConnectionString: provider.GetDBURL(),
	}

	connInfo.db, err = sqlx.Open(connInfo.DBDriverName, connInfo.ConnectionString)

	if err != nil {
		panic(err)
	}

	connInfo.db.SetMaxOpenConns(maxConnections)

	if err != nil {
		return err
	}

	d.connMapString[connName] = connInfo

	return nil
}

// Retrieves statistics about the database environment
func (d *DataAccessLayer) GetStats(connName string) sql.DBStats {
	db, _ := d.GetDB(connName)
	return db.Stats()
}

// GetDB will return an opened database from the sql subsystem.
// The database is referenced by name.
func (d *DataAccessLayer) GetDB(connName string) (*sqlx.DB, error) {
	info := d.connMapString[connName]

	if info == nil {
		return nil, fmt.Errorf("Cannot find named database connection")
	}

	return info.db, nil
}

// getErrorNumber interrogates an error and, if it is a MySQL error, it
// returns the error code and message
func (db *DataAccessLayer) getErrorNumber(err error) (errCode uint16, errMsg string) {
	sqlerr, ok := err.(*mysql.MySQLError)

	if ok {
		return sqlerr.Number, sqlerr.Message
	}

	return uint16(1), err.Error()
}

// RunWithDb retrieves a named database and passes it to the provided function.
func (db *DataAccessLayer) RunWithDb(connName string, op func(conndb *sqlx.DB)) {
	condb, errdb := db.GetDB(connName)

	if errdb != nil {
		fmt.Println("dal getdb:", errdb.Error())
		panic(errdb)
	}

	op(condb)
}

// Select will run a query against a database and return the results in the supplied interface
func (db *DataAccessLayer) Select(connName string, dest interface{}, query string, args ...interface{}) (err error) {
	db.RunWithDb(connName, func(condb *sqlx.DB) {
		err = condb.Select(dest, query, args...)
	})
	return err
}

// Get will run a query against a database and return a single row in the supplied interface
func (db *DataAccessLayer) Get(connName string, dest interface{}, query string, args ...interface{}) (err error) {
	db.RunWithDb(connName, func(condb *sqlx.DB) {
		err = condb.Get(dest, query, args...)
	})
	return err
}

// Exec will run a statement against a database and return a sql.Result object
func (db *DataAccessLayer) Exec(connName string, query string, args ...interface{}) (result sql.Result, err error) {
	db.RunWithDb(connName, func(condb *sqlx.DB) {
		result, err = condb.Exec(query, args...)
	})

	return result, err
}

func (db *DataAccessLayer) NamedExec(connName string, query string, arg interface{}) (result sql.Result, err error) {
	db.RunWithDb(connName, func(condb *sqlx.DB) {
		result, err = condb.NamedExec(query, arg)
	})

	return result, err
}
