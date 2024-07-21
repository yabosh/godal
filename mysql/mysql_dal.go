package mysql

import (
	"fmt"

	"github.com/yabosh/godal/dalconfig"
)

// MySQLDal represents a connection to a MySQL database
type MySQLDal struct {
	DBHost     string
	DBPort     string
	Username   string
	Password   string
	Database   string
	DriverName string
	DBUrl      string
}

// NewMySQLDal creates a new, initialized instance of MySQLDal that is intended to be used
// by the Dal (data access layer) component.
func New(settings dalconfig.Settings) *MySQLDal {
	return &MySQLDal{
		DBHost:     settings.Host,
		DBPort:     settings.Port,
		Username:   settings.Username,
		Password:   settings.Password,
		Database:   settings.DBName,
		DriverName: "mysql",
		DBUrl: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=%t&parseTime=true",
			settings.Username,
			settings.Password,
			settings.Host,
			settings.Port,
			settings.DBName,
			settings.AllowMultipleStatements),
	}
}

// GetDBURL implements the ConnectionProvider.GetDBURL interface
func (m *MySQLDal) GetDBURL() string {
	return m.DBUrl
}

// GetDBDriver implements the ConnectionProvider.GetDBDriver interface
func (m *MySQLDal) GetDBDriver() string {
	return m.DriverName
}
