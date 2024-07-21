package godalmock

// GoDALMock represents a connection to mock database
type GoDALMock struct {
	DriverName string
	DBUrl      string
}

// New creates a new, initialized instance of the mock dal.
func New(dsn string) *GoDALMock {
	return &GoDALMock{
		DriverName: "sqlmock",
		DBUrl:      dsn,
	}
}

// GetDBURL implements the ConnectionProvider.GetDBURL interface
func (m *GoDALMock) GetDBURL() string {
	return m.DBUrl
}

// GetDBDriver implements the ConnectionProvider.GetDBDriver interface
func (m *GoDALMock) GetDBDriver() string {
	return m.DriverName
}
