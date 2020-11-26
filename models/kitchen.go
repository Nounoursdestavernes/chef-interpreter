package models

type Kitchen struct {
	MixingBowls  []MixingBowl
	BakingDishes []BakingDish
}

type MixingBowl []Ingredient
type BakingDish []Ingredient
