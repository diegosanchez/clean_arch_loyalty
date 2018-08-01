package Loyalty

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

import "github.com/diegosanchez/clean_arch_loyalty/adapter"
import "github.com/diegosanchez/clean_arch_loyalty/entity"

/*
Action items:
	Investigar ventajas y desventajas de exponer las entidades a los Controller. Hacer pruebas en código los siguientes casos de uso:

		- Tenemos un usuario con loyalty nivel 2 y un item de precio mayor a $1200 el envío es gratis.
		- Tenemos un usuario con loyalty nivel 2 y un item de precio menor a $1200 el envío no es gratis.

Lo planteamos desde el contexto de un backend en tecnologías Java y NodeJS.
*/

func Test_caso_uso_item_loyalty_2_item_mayor_a_1200_envio_gratis(t *testing.T) {
	uc := NewLoyaltyFreeShipping(adapter.NewTestItemGateway(),
		adapter.NewTestShippingOptionGateway(),
		adapter.NewTestLoyaltyDM())

	req := make(map[string]interface{})
	res := make(map[string]interface{})

	req["item_id"] = entity.NewItemId("MLA", 653840397)
	req["user_id"] = entity.NewUserId(879)

	uc.doWork(req, res)

	assert.Equal(t, "Envio gratarola!",
		res["review.shipping"].(string),
		"choOptions - free")
}

func Test_caso_uso_item_loyalty_4_item_menor_a_1200_envio_gratis(t *testing.T) {
	uc := NewLoyaltyFreeShipping(adapter.NewTestItemGateway(),
		adapter.NewTestShippingOptionGateway(),
		adapter.NewTestLoyaltyDM())

	req := make(map[string]interface{})
	res := make(map[string]interface{})

	req["item_id"] = entity.NewItemId("MLA", 653840300)
	req["user_id"] = entity.NewUserId(879)

	uc.doWork(req, res)

	assert.Equal(t, "Envio GARPO!",
		res["review.shipping"].(string),
		"choOptions - free")
}
