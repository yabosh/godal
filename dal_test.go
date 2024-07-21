package godal_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/yabosh/godal"
	"github.com/yabosh/godal/godalmock"
)

type TestRecord struct {
	ID string `db:"id"`
}

func Test_add_connection_to_dal(t *testing.T) {
	const connName = "test_conn_name"

	// Given SQL Mock
	_, mock, _ := sqlmock.NewWithDSN(connName)

	// And a set of expected rows
	rows := sqlmock.NewRows([]string{"id"}).
		AddRow("1").
		AddRow("2")
	_ = mock.ExpectQuery("^SELECT").WillReturnRows(rows)

	// And a DAL connection that uses the mock
	cp := godalmock.New(connName)
	godal.Dal().AddConnection(connName, cp, 10)

	// When a SQL query is executed
	response := make([]TestRecord, 0)
	err := godal.Dal().Select(connName,
		&response,
		`
		SELECT id
		FROM table;
		`)

	// Then a result should be returned
	assert.Nil(t, err)
	assert.Len(t, response, 2)
	assert.Equal(t, "1", response[0].ID)
}
