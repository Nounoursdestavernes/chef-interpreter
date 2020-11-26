package lexer

import (
	"ChefInterpreter/models"
	"fmt"
	"strings"
)

func tokenizeIngredients(input string) (ingredients []models.Ingredient, err error) {
	list := strings.Split(input, "\n")
	for _, item := range list {
		fmt.Printf("Ingredient: %s\n", item)
	}
	return
}
