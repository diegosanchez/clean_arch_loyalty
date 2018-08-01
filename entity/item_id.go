package entity

import (
	"strconv"
	"strings"
)

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

func (this *ItemId) RenderPattern(pattern string) string {
	result := strings.Replace(pattern,
		"{site}", strings.ToLower(this.site), -1)
	result = strings.Replace(result,
		"{nro}", strconv.Itoa(this.id), -1)

	return result
}
