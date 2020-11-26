package models

type Recipe struct {
	Title            string
	Comments         string
	Ingredients      []Ingredient
	CookingTime      int
	OvenTemperature  int
	GasMark          int
	Method           []MethodStatement
	Serves           int
	AuxiliaryRecipes []Recipe
}

type Ingredient struct {
	Name  string
	Value int
	IsDry bool
}

type MethodStatement struct {
	Command    Command
	Ingredient string
	MixingBowl int
	BakingDish int
	Minutes    int
	Hours      int
	Verb       string
	Recipe     string
}

type Command int

const (
	CommandTake Command = iota
	CommandPut
	CommandFold
	CommandAdd
	CommandRemove
	CommandCombine
	CommandDivide
	CommandAddDry
	CommandLiquefyIngredient
	CommandLiquefyBowl
	CommandStirBowl
	CommandStirIngredient
	CommandMix
	CommandClean
	CommandPour
	CommandVerbStart
	CommandVerbEnd
	CommandSet
	CommandServe
	CommandRefrigerate
)
