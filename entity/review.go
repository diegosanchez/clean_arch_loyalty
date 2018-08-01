package entity

type Review interface {
	AsMapForReview(bag *Databag) *Databag
	PopulateMapForReview(bag *Databag) *Databag
}
