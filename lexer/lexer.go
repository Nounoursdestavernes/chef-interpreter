package lexer

import (
	"ChefInterpreter/models"
	"fmt"
	"strings"
)

func Tokenize(input string) (recipe models.Recipe, err error) {
	recipe, remainingInput, err := tokenizeSingleRecipe(input)
	if err != nil {
		return models.Recipe{}, err
	}

	for len(remainingInput) > 0 {
		nextRecipe := models.Recipe{}
		nextRecipe, remainingInput, err = tokenizeSingleRecipe(remainingInput)
		if err != nil {
			return models.Recipe{}, err
		}

		recipe.AuxiliaryRecipes = append(recipe.AuxiliaryRecipes, nextRecipe)
	}

	return recipe, nil
}

func tokenizeTitle(input string) (title string, err error) {
	if !strings.HasSuffix(input, ".") {
		return "", fmt.Errorf("could not find Title field")
	}

	return input, nil
}

func tokenizeComments(input string) (comments string, err error) {
	return input, nil
}

func tokenizeCookingTime(input string) (cookingTime int, err error) {
	// TODO: handle
	return
}

func tokenizePreheat(input string) (ovenTemperature, gasMark int, err error) {
	// TODO: handle
	return
}

func tokenizeServes(input string) (serves int, err error) {
	// TODO: handle
	return
}

func tokenizeSingleRecipe(input string) (recipe models.Recipe, remainingInput string, err error) {
	fields := strings.Split(input, "\n\n")
	fieldIndex := 0

	// title
	recipe.Title, err = tokenizeTitle(fields[fieldIndex])
	if err != nil {
		remainingInput = strings.Join(fields[fieldIndex+1:], "\n\n")
		return models.Recipe{}, remainingInput, err
	}
	fieldIndex++

	// comments
	if !strings.HasPrefix(fields[fieldIndex], "Ingredients.") {
		recipe.Comments, _ = tokenizeComments(fields[fieldIndex])
		fieldIndex++
	}

	// ingredients
	recipe.Ingredients, err = tokenizeIngredients(fields[fieldIndex])
	if err != nil {
		remainingInput = strings.Join(fields[fieldIndex+1:], "\n\n")
		return models.Recipe{}, remainingInput, err
	}
	fieldIndex++

	// cooking time
	if strings.HasPrefix(fields[fieldIndex], "Cooking time:") {
		recipe.CookingTime, _ = tokenizeCookingTime(fields[fieldIndex])
		fieldIndex++
	}

	// pre-heat
	if strings.HasPrefix(fields[fieldIndex], "Pre-heat") {
		recipe.OvenTemperature, recipe.GasMark, _ = tokenizePreheat(fields[fieldIndex])
		fieldIndex++
	}

	// method
	recipe.Method, err = tokenizeMethod(fields[fieldIndex])
	if err != nil {
		remainingInput = strings.Join(fields[fieldIndex+1:], "\n\n")
		return models.Recipe{}, remainingInput, err
	}
	fieldIndex++

	// serves
	if fieldIndex < len(fields) && strings.HasPrefix(fields[fieldIndex], "Serves") {
		recipe.Serves, _ = tokenizeServes(fields[fieldIndex])
		fieldIndex++
	}

	remainingInput = strings.Join(fields[fieldIndex:], "\n\n")
	remainingInput = strings.TrimSpace(remainingInput)
	return recipe, remainingInput, nil
}
