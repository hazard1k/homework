package models

type ItemPrice struct {
	Base       float64
	Discounted float64
}

type Item struct {
	Id          string
	Name        string
	Article     string
	Category    string
	Description string
	Price       ItemPrice
}
