package entity

type FreeShipping struct {
	item *Item
}

func NewFreeShipping(item *Item) *FreeShipping {
	result := new(FreeShipping)

	result.item = item

	return result
}

func (this *FreeShipping) AsMapForReview(res *Databag) *Databag {
	return this.PopulateMapForReview(res)
}

func (this *FreeShipping) PopulateMapForReview(res *Databag) *Databag {
	return res.set("review.shipping", "Envio gratarola!")
}
