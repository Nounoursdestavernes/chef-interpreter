package lexer

import (
	"ChefInterpreter/models"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTokenizeMethod(t *testing.T) {
	method, err := tokenizeMethod(methodValid)
	assert.Nil(t, err)

	method, err = tokenizeMethod(methodInvalidThe)
	assert.Errorf(t, err, "invalid method statement: %s", methodInvalidThe)
	printMethod(method)

	method, err = tokenizeMethod(methodInvalidA)
	assert.Errorf(t, err, "invalid method statement: %s", methodInvalidA)
	printMethod(method)

	method, err = tokenizeMethod(methodInvalidOrdinal)
	assert.Errorf(t, err, "invalid method statement: %s", methodInvalidOrdinal)
	printMethod(method)

	method, err = tokenizeMethod(methodInvalidLiquefy)
	assert.Errorf(t, err, "invalid method statement: %s", methodInvalidLiquefy)
	printMethod(method)

	method, err = tokenizeMethod(methodInvalidSingularMinute)
	assert.Errorf(t, err, "invalid method statement: %s", methodInvalidSingularMinute)
	printMethod(method)

	method, err = tokenizeMethod(methodInvalidNumericRecipe)
	assert.Errorf(t, err, "invalid method statement: %s", methodInvalidNumericRecipe)
	printMethod(method)

	method, err = tokenizeMethod(methodInvalidGarbage)
	assert.Errorf(t, err, "invalid method statement: %s", methodInvalidGarbage)
	printMethod(method)
}

func printMethod(method []models.MethodStatement) {
	for _, statement := range method {
		fmt.Printf("OPCODE: %d\n\tINGREDIENT: %s\n\tMIXING BOWL: %d\n\tBAKING DISH: %d\n\tMINUTES: %d\n\tHOURS: %d\n\tVERB START: %s\n\tVERB END: %s\n\tRECIPE: %s\n\n",
			statement.Command,
			statement.Ingredient,
			statement.MixingBowl,
			statement.BakingDish,
			statement.Minutes,
			statement.Hours,
			statement.VerbStart,
			statement.VerbEnd,
			statement.Recipe,
		)
	}
}

const methodValid = `Take cauliflower from refrigerator.
Put cauliflower into mixing bowl.
Put cauliflower into 1st mixing bowl.
Fold cauliflower into mixing bowl.
Fold cauliflower into 2nd mixing bowl.
Add cauliflower.
Add cauliflower to mixing bowl.
Add cauliflower to 3rd mixing bowl.
Remove cauliflower.
Remove cauliflower from mixing bowl.
Remove cauliflower from 4th mixing bowl.
Combine cauliflower.
Combine cauliflower into mixing bowl.
Combine cauliflower into 5th mixing bowl.
Divide cauliflower.
Divide cauliflower into mixing bowl.
Divide cauliflower into 6th mixing bowl.
Add dry ingredients to mixing bowl.
Add dry ingredients to 7th mixing bowl.
Liquefy cauliflower.
Liquefy contents of the mixing bowl.
Liquefy contents of the 8th mixing bowl.
Stir for 1 minutes.
Stir the mixing bowl for 10 minutes.
Stir the 9th mixing bowl for 25 minutes.
Stir cauliflower into the mixing bowl.
Stir cauliflower into the 10th mixing bowl.
Mix well.
Mix the mixing bowl well.
Mix the 11th mixing bowl well.
Clean mixing bowl.
Clean 12th mixing bowl.
Pour contents of the mixing bowl into the baking dish.
Pour contents of the mixing bowl into the 14th baking dish.
Pour contents of the 13th mixing bowl into the baking dish.
Pour contents of the 13th mixing bowl into the 14th baking dish.
Set aside.
Serve with caramel sauce.
Refrigerate.
Refrigerate for 4 hours.
Defenestrate the cauliflower until defenestrated.
Bloat the cauliflower until bloated.
Magnify the cauliflower until magnified.
Defenestrate the cauliflower.`
const methodInvalidThe = `Take cauliflower from the refrigerator.`
const methodInvalidA = `Put cauliflower into a mixing bowl.`
const methodInvalidOrdinal = `Add cauliflower to third mixing bowl.`
const methodInvalidLiquefy = `Liquify cauliflower.`
const methodInvalidSingularMinute = `Stir for 1 minute.`
const methodInvalidNumericRecipe = `Serve with caramel sauce 3.`
const methodInvalidGarbage = `74(&^$*^%#*^(*^$(&539$%9`
