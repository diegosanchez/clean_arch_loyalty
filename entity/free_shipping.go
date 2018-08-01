package entity

type FreeShipping struct {
	item *Item
}

func NewFreeShipping(item *Item) *FreeShipping {
	result := new(FreeShipping)

	result.item = item

	return result
}

func (this *FreeShipping) AsMapForReview(res *map[string]interface{}) *map[string]interface{} {
	return this.PopulateMapForReview(res)
}

func (this *FreeShipping) PopulateMapForReview(res *map[string]interface{}) *map[string]interface{} {
	(*res)["review.shipping"] = "Envio gratarola!"
	return res
}
