package lexer

import (
	"ChefInterpreter/models"
	"strings"
)

func Tokenize(input string) (recipe models.Recipe, err error) {
	fields := strings.Split(input, "\n\n")

	recipe.Title = fields[0]

	for _, field := range fields[1:] {
		switch {
		case strings.HasPrefix(field, "Ingredients."):
			ingredientsField := strings.Replace(field, "Ingredients.\n", "", 1)
			ingredients, err := tokenizeIngredients(ingredientsField)
			if err != nil {
				return models.Recipe{}, nil
			}
			recipe.Ingredients = ingredients

		case strings.HasPrefix(field, "Cooking time:"):

		case strings.HasPrefix(field, "Pre-heat"):

		case strings.HasPrefix(field, "Method."):
			methodField := strings.Replace(field, "Method.\n", "", 1)
			method, err := tokenizeMethod(methodField)
			if err != nil {
				return models.Recipe{}, nil
			}
			recipe.Method = method

		case strings.HasPrefix(field, "Serves"):

		case strings.HasPrefix(field, "Cooking time: "):

		case len(recipe.Method) == 0:
			recipe.Comments = field

		case strings.HasSuffix(field, "."):
			// starts an auxiliary recipe
		}
	}

	return recipe, nil
}
