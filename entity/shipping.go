package entity

type Shipping interface {
	AsMapForReview(res *map[string]interface{}) *map[string]interface{}
	PopulateMapForReview(res *map[string]interface{}) *map[string]interface{}
}
