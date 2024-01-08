package bookmodel

type PriceRange struct {
	Min float64 `json:"min,omitempty" form:"min"`
	Max float64 `json:"max,omitempty" form:"max"`
}

type Filter struct {
	Name  string     `json:"name,omitempty" form:"name"`
	Price PriceRange `json:"price,omitempty" form:"price"`
}
