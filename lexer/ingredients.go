package lexer

import (
	"ChefInterpreter/models"
	"fmt"
	"strings"
)

func tokenizeIngredients(input string) (ingredients []models.Ingredient, err error) {
	if !strings.HasPrefix(input, "Ingredients.") {
		return []models.Ingredient{}, fmt.Errorf("could not find Ingredients field")
	}

	ingredientsField := strings.Replace(input, "Ingredients.\n", "", 1)
	list := strings.Split(ingredientsField, "\n")
	for _, item := range list {
		fmt.Printf("Ingredient: %s\n", item)
	}

	return ingredients, nil
}
