package entity

type CustomShipping struct {
	item *Item
}

func NewCustomShipping(item *Item) *CustomShipping {
	result := new(CustomShipping)

	result.item = item

	return result
}

func (this *CustomShipping) AsMapForReview(res *map[string]interface{}) *map[string]interface{} {
	shipping := this.item.ShippingBasedOnPrice(this)
	return shipping.PopulateMapForReview(res)
}

func (this *CustomShipping) PopulateMapForReview(res *map[string]interface{}) *map[string]interface{} {
	(*res)["review.shipping"] = "Envio GARPO!"
	return res
}
