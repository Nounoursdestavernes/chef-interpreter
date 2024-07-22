package lexer

import (
	"github.com/Nounoursdestavernes/chef-interpreter/models"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Lexer struct {
	fields     []string
	fieldIndex int
}

func New() Lexer {
	return Lexer{}
}

func (l Lexer) Tokenize(input string) (recipe models.Recipe, err error) {
	l.fields = strings.Split(input, "\n\n")
	l.fieldIndex = 0

	recipe, err = l.tokenizeSingleRecipe()
	if err != nil {
		return models.Recipe{}, err
	}

	for l.fieldIndex < len(l.fields) {
		nextRecipe := models.Recipe{}
		nextRecipe, err = l.tokenizeSingleRecipe()
		if err != nil {
			return models.Recipe{}, err
		}

		recipe.AuxiliaryRecipes = append(recipe.AuxiliaryRecipes, nextRecipe)
	}

	return recipe, nil
}

func (l Lexer) tokenizeSingleRecipe() (recipe models.Recipe, err error) {
	// title
	recipe.Title, err = l.tokenizeTitle(l.getNextField())
	if err != nil {
		return models.Recipe{}, err
	}

	// comments
	recipe.Comments, _ = l.tokenizeComments(l.getNextField())

	// ingredients
	recipe.Ingredients, err = l.tokenizeIngredients(l.getNextField())
	if err != nil {
		return models.Recipe{}, err
	}

	// cooking time
	recipe.CookingTime, err = l.tokenizeCookingTime(l.getNextField())
	if err != nil {
		return models.Recipe{}, err
	}

	// pre-heat
	recipe.OvenTemperature, recipe.GasMark, err = l.tokenizePreheat(l.getNextField())
	if err != nil {
		return models.Recipe{}, err
	}

	// method
	recipe.Method, err = l.tokenizeMethod(l.getNextField())
	if err != nil {
		return models.Recipe{}, err
	}

	// serves
	recipe.Serves, err = l.tokenizeServes(l.getNextField())
	if err != nil {
		return models.Recipe{}, err
	}

	return recipe, nil
}

func (l Lexer) tokenizeTitle(input string) (title string, err error) {
	if !strings.HasSuffix(input, ".") {
		return "", fmt.Errorf("could not find Title field")
	}

	l.markFieldComplete()
	return input, nil
}

func (l Lexer) tokenizeComments(input string) (comments string, err error) {
	if strings.HasPrefix(input, "Ingredients.") {
		return "", nil
	}

	l.markFieldComplete()
	return input, nil
}

func (l Lexer) tokenizeCookingTime(input string) (cookingTime int, err error) {
	if !strings.HasPrefix(input, "Cooking time:") {
		return 0, nil
	}

	tokens := regexp.MustCompile("^Cooking time: ([0-9]*) (?:hour|hours|minute|minutes).$").FindStringSubmatch(input)
	if len(tokens) == 0 {
		return 0, fmt.Errorf("invalid cooking time statement: %s", input)
	}

	cookingTime, _ = strconv.Atoi(tokens[1])

	l.markFieldComplete()
	return cookingTime, nil
}

func (l Lexer) tokenizePreheat(input string) (ovenTemperature, gasMark int, err error) {
	if !strings.HasPrefix(input, "Pre-heat") {
		return 0, 0, nil
	}

	tokens := regexp.MustCompile("^Pre-heat oven to ([0-9]*) degrees (?:Celsius|Fahrenheit)(?: \\(gas mark ([0-9]*)\\)|).$").FindStringSubmatch(input)
	if len(tokens) == 0 {
		return 0, 0, fmt.Errorf("invalid preheat statement: %s", input)
	}

	ovenTemperature, _ = strconv.Atoi(tokens[1])
	ovenTemperature, _ = strconv.Atoi(tokens[2])

	l.markFieldComplete()
	return ovenTemperature, gasMark, nil
}

func (l Lexer) tokenizeServes(input string) (serves int, err error) {
	if !strings.HasPrefix(input, "Serves") {
		return 0, nil
	}

	tokens := regexp.MustCompile("^Serves ([0-9]*).$").FindStringSubmatch(input)
	if len(tokens) == 0 {
		return 0, fmt.Errorf("invalid serves statement: %s", input)
	}

	serves, _ = strconv.Atoi(tokens[1])

	l.markFieldComplete()
	return serves, nil
}

func (l Lexer) getNextField() string {
	return l.fields[l.fieldIndex]
}

func (l Lexer) markFieldComplete() {
	l.fieldIndex++
}
