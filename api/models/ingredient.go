package models

import "math/big"

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
	Amount  big.Rat
	Measure Measure
}

type Ingredient struct {
	Name     string
	Quantity Measurement
}
