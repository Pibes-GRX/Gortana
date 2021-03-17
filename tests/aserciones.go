package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDBconectionIP(t *testing.T, a string) {

	var b string = "localhost"

	assert.Equal(t, a, b, "La conexión con la base de datos no es el localhost")

}
