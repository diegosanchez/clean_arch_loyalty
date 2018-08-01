package entity

// @TODO
// Esta es la responsabilidad de nuestras clases?
type Review interface {
	AsMapForReview(bag *Databag) *Databag
	PopulateMapForReview(bag *Databag) *Databag
}
