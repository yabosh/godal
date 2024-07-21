# Go Data Access Layer (godal)

Go-DAL is an intermediate data access layer that provides an abstraction between business logic and the go sql library.  Dal provides a simple facade for SQL functionality that is useful in  most sql use cases.

Go-DAL is designed to limit the need to pass active database connections through different layers of code.  Each data source can be defined and stored in the DAL cache and retrieve by name.

Go-DAL currently supports connectivity to MySQL and SQL Server databases.  Connectivity to other
dbs is supported by importing the appropriate driver.

## Example - Create a named sdatabase connection

```go
config := dalconfig.Settings{}
config.Host = "mydatabasehost"
config.Port = "1433"
config.Username = "myuser"
config.Password = "mypassword"
config.ConnectionName = "ThisConnection"
config.DBName = "sql_schema"
config.MaxConnections = 25
config.Instance = "used_for_sql_server"

// Create an instance of a SQL Server data access layer
dal := sqlserver.New(config)

// Add this connection to the DAL cache where it can be referenced by name
godal.Dal().AddConnection(config.ConnectionName, dal, config.MaxConnections)
```

## Example - Retrieve using named connection

Once a database connection has been saved in the cache, any database operations such as SELECT or UPDATE simply need to reference the desired connection by name.

```go
response = make([]models.Customers, 0)

// Retrieve rows from db and populate Customers array.
err = godal.Dal().Select(repo.connName, &response, `
    SELECT name
    FROM Customer;
`)

```

