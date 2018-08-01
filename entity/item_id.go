package entity

type ItemId struct {
	site string
	id   int
}

func NewItemId(site string, id int) *ItemId {
	result := new(ItemId)

	result.site = site
	result.id = id

	return result
}
