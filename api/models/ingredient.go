package models

type Measure int

const (
	Teaspoon Measure = iota
	Tablespoon
	Ounce
	Cup
	Pint
	Quart
	Gallon
)

type Measurement struct {
	Amount  float64
	Measure Measure
}

type Ingredient struct {
	Name     string
	Quantity Measurement
}
