package entity

type CustomShipping struct {
	item *Item
}

func NewCustomShipping(item *Item) *CustomShipping {
	result := new(CustomShipping)

	result.item = item

	return result
}

func (this *CustomShipping) AsMapForReview(res *Databag) *Databag {
	shipping := this.item.ShippingBasedOnPrice(this)
	return shipping.PopulateMapForReview(res)
}

func (this *CustomShipping) PopulateMapForReview(res *Databag) *Databag {
	return res.set("review.shipping", "Envio GARPO!")
}
