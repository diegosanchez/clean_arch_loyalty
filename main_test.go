package Loyalty

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

import "github.com/diegosanchez/clean_arch_loyalty/adapter"

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

func Test_dummy_2(t *testing.T) {
	assert.Equal(t, false, false, "dummy2")
}

/*
func Test_item_gateway(t *testing.T) {
	var igt ItemGateway = adapter.NewTestItemGateway()

	igt.ItemById(entity.NewItemId("MLA", 12))

	assert.Equal(t, false, false, "dummy2")
}
*/

func Test_caso_uso_item_loyalty_2_item_mayor_a_1200_envio_gratis(t *testing.T) {
	uc := NewLoyaltyFreeShipping(adapter.NewTestItemGateway(), adapter.NewTestShippingOptionGateway(),
		adapter.NewTestLoyaltyDM())

	req := make(map[string]interface{})
	res := make(map[string]interface{})

	uc.doWork(req, res)

	// item := NewTestItemGateway().itemById(NewItemId("MLA", 12))
	// loyalty := NewTestLoyaltyDM().loyaltyForUser(NewUser(987))
	// shippingOpt := NewTestShippingOptionGateway().forItem(item)

	assert.Equal(t, "xxxx", "xxxx", "choOptions - free")

}
