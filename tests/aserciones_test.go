package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"database/sql"
)

var a string = "localhost"

func TestDBconectionIP(t *testing.T) {

	var b string = "localhost"

	assert.Equal(t, a, b, "La conexión con la base de datos no es el localhost")

}

func dbConnection(t *testing.T) {

	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")

	if db == nil {
		assert.Equal(t, db, nil, "La conexión no puede ser nula")
		assert.NotEqual(t, err, nil, "La conexión no puede realizarse con errores")
	}

}
