package discounter

type DiscountType string

const (
	Category DiscountType = "Category"
	Item     DiscountType = "Item"
	All      DiscountType = "-"
)

type Discounts []*Discount

type Discount struct {
	Type   DiscountType
	Value  string
	Amount float64
}
