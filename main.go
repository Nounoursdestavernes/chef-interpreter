package main

import (
	"ChefInterpreter/lexer"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	args := os.Args

	input, err := getInputFromFile(args[1])
	if err != nil {
		log.Fatal(err)
	}

	recipe, err := lexer.Tokenize(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Title: %s\nComments: %s\nIngredients: %v\nCooking time: %v\nOven temperature: %v\nGas mark: %v\nMethod: %v\nServes: %v\n\n",
		recipe.Title,
		recipe.Comments,
		recipe.Ingredients,
		recipe.CookingTime,
		recipe.OvenTemperature,
		recipe.GasMark,
		recipe.Method,
		recipe.Serves)

	for _, auxiliaryRecipe := range recipe.AuxiliaryRecipes {
		fmt.Printf("Title: %s\nComments: %s\nIngredients: %v\nCooking time: %v\nOven temperature: %v\nGas mark: %v\nMethod: %v\nServes: %v\n\n",
			auxiliaryRecipe.Title,
			auxiliaryRecipe.Comments,
			auxiliaryRecipe.Ingredients,
			auxiliaryRecipe.CookingTime,
			auxiliaryRecipe.OvenTemperature,
			auxiliaryRecipe.GasMark,
			auxiliaryRecipe.Method,
			auxiliaryRecipe.Serves)
	}
}

func getInputFromFile(filename string) (string, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
