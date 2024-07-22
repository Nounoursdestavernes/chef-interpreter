package lexer

import (
	"github.com/Nounoursdestavernes/chef-interpreter/models"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func (l Lexer) tokenizeIngredients(input string) (ingredients []models.Ingredient, err error) {
	if !strings.HasPrefix(input, "Ingredients.") {
		return []models.Ingredient{}, fmt.Errorf("could not find Ingredients field")
	}

	input = strings.Replace(input, "Ingredients.\n", "", 1)
	list := strings.Split(input, "\n")
	for _, item := range list {
		tokens := ingredientRegexp.FindStringSubmatch(item)
		if len(tokens) == 0 {
			return []models.Ingredient{}, fmt.Errorf("invalid ingredient: %s", item)
		}

		ingredient := models.Ingredient{}
		ingredient.Name = tokens[4]
		ingredient.Amount, _ = strconv.Atoi(tokens[1]) // TODO: errors possible here?

		switch tokens[3] {
		case "g", "kg", "pinch":
			ingredient.IsDry = true
		case "ml", "l", "dash":
			ingredient.IsDry = false
		case "cup", "teaspoon", "tablespoon":
			if tokens[2] == "heaped" || tokens[2] == "level" {
				ingredient.IsDry = true
			} else {
				ingredient.IsDry = false
			}
		}

		ingredients = append(ingredients, ingredient)
	}

	l.markFieldComplete()
	return ingredients, nil
}

var ingredientRegexp = regexp.MustCompile("^([0-9]*)(?: |)(heaped|level|) (g|kg|pinch(?:es|)|ml|l|dash(?:es|)|cup(?:s|)|teaspoon(?:s|)|tablespoon(?:s|)) ([A-Za-z ]*)$")
