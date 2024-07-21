package sqlserver

import (
	"fmt"
	"strings"

	"github.com/yabosh/godal/dalconfig"
)

// MySQLDal represents a connection to a MySQL database
type SQLServerDal struct {
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
func New(settings dalconfig.Settings) *SQLServerDal {
	var hoststring string
	if strings.TrimSpace(settings.Instance) == "" || strings.ToLower(settings.Instance) == "default" {
		hoststring = fmt.Sprintf("%s:%s", settings.Host, settings.Port)
	} else {
		hoststring = fmt.Sprintf("%s:%s/%s", settings.Host, settings.Port, settings.Instance)
	}

	url := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&connection+timeout=%d",
		settings.Username,
		settings.Password,
		hoststring,
		settings.DBName,
		30,
	)

	return &SQLServerDal{
		DBHost:     settings.Host,
		DBPort:     settings.Port,
		Username:   settings.Username,
		Password:   settings.Password,
		Database:   settings.DBName,
		DriverName: "sqlserver",
		DBUrl:      url,
	}
}

// GetDBURL implements the ConnectionProvider.GetDBURL interface
func (m *SQLServerDal) GetDBURL() string {
	return m.DBUrl
}

// GetDBDriver implements the ConnectionProvider.GetDBDriver interface
func (m *SQLServerDal) GetDBDriver() string {
	return m.DriverName
}
