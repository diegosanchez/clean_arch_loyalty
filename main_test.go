package Loyalty

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
Action items:
	Investigar ventajas y desventajas de exponer las entidades a los Controller. Hacer pruebas en código los siguientes casos de uso:

		- Tenemos un usuario con loyalty nivel 2 y un item de precio mayor a $1200 el envío es gratis.
		- Tenemos un usuario con loyalty nivel 2 y un item de precio menor a $1200 el envío no es gratis.

Lo planteamos desde el contexto de un backend en tecnologías Java y NodeJS.
*/

func Test_dummy(t *testing.T) {
	assert.Equal(t, true, true, "dummy")
}
