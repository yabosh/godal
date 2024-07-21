package dalconfig

// DalConfig contains the values used to configure
// a new Dal object.
type Settings struct {
	ConnectionName          string
	Host                    string
	Port                    string
	Instance                string
	Username                string
	Password                string
	DBName                  string
	MaxConnections          int
	AllowMultipleStatements bool
}
