package lexer

import (
	"ChefInterpreter/models"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func tokenizeIngredients(input string) (ingredients []models.Ingredient, err error) {
	if !strings.HasPrefix(input, "Ingredients.") {
		return []models.Ingredient{}, fmt.Errorf("could not find Ingredients field")
	}

	input = strings.Replace(input, "Ingredients.\n", "", 1)
	list := strings.Split(input, "\n")
	for _, item := range list {
		result := ingredientRegexp.FindStringSubmatch(item)
		if len(result) == 0 {
			return []models.Ingredient{}, fmt.Errorf("invalid ingredient: %s", item)
		}

		ingredient := models.Ingredient{}
		ingredient.Name = result[4]
		ingredient.Amount, _ = strconv.Atoi(result[1]) // TODO: errors possible here?

		switch result[3] {
		case "g", "kg", "pinch":
			ingredient.IsDry = true
		case "ml", "l", "dash":
			ingredient.IsDry = false
		case "cup", "teaspoon", "tablespoon":
			if result[2] == "heaped" || result[2] == "level" {
				ingredient.IsDry = true
			} else {
				ingredient.IsDry = false
			}
		}

		ingredients = append(ingredients, ingredient)
	}

	return ingredients, nil
}

var ingredientRegexp = regexp.MustCompile("^([0-9]*)(?: |)(heaped|level|) (g|kg|pinch(?:es|)|ml|l|dash(?:es|)|cup(?:s|)|teaspoon(?:s|)|tablespoon(?:s|)) ([A-Za-z ]*)$")
