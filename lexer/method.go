package lexer

import (
	"ChefInterpreter/models"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func tokenizeMethod(input string) (method []models.MethodStatement, err error) {
	statements := strings.SplitAfter(input, ".")

	for _, statement := range statements {
		statement = strings.TrimSpace(statement)
		if statement == "" {
			continue
		}

		isValid := false

		for _, command := range commandOptions {
			result := command.regexp.FindStringSubmatch(statement)

			if len(result) > 0 {
				isValid = true

				methodStatement, err := populateMethodStatement(command.command, command.regexp.SubexpNames()[1:], result[1:])
				if err != nil {
					return []models.MethodStatement{}, err
				}

				method = append(method, methodStatement)
				break
			}
		}

		if !isValid {
			return []models.MethodStatement{}, fmt.Errorf("invalid method statement: %s", statement)
		}
	}

	return method, nil
}

func populateMethodStatement(command models.Command, fieldNames []string, fieldValues []string) (methodStatement models.MethodStatement, err error) {
	methodStatement.Command = command

	for field, fieldName := range fieldNames {
		switch fieldName {
		case "ingredient":
			methodStatement.Ingredient = fieldValues[field]
		case "bowl":
			methodStatement.MixingBowl, _ = strconv.Atoi(fieldValues[field]) // TODO: errors possible here?
		case "dish":
			methodStatement.BakingDish, _ = strconv.Atoi(fieldValues[field]) // TODO: errors possible here?
		case "minutes":
			methodStatement.Minutes, _ = strconv.Atoi(fieldValues[field]) // TODO: errors possible here?
		case "hours":
			methodStatement.Hours, _ = strconv.Atoi(fieldValues[field]) // TODO: errors possible here?
		case "verbstart":
			methodStatement.VerbStart = fieldValues[field]
		case "verbend":
			methodStatement.VerbEnd = fieldValues[field]
		case "recipe":
			methodStatement.Recipe = fieldValues[field]
		}
	}

	return methodStatement, nil
}

type commandOption struct {
	regexp  *regexp.Regexp
	command models.Command
}

var commandOptions = [...]commandOption{
	{
		regexp:  regexp.MustCompile("^Take (?P<ingredient>[A-Za-z ]*?) from refrigerator.$"),
		command: models.CommandTake,
	},
	{
		regexp:  regexp.MustCompile("^Put (?P<ingredient>[A-Za-z ]*) into(?: (?P<bowl>[0-9]*)(?:st|nd|rd|th)|) mixing bowl.$"),
		command: models.CommandPut,
	},
	{
		regexp:  regexp.MustCompile("^Fold (?P<ingredient>[A-Za-z ]*) into(?: (?P<bowl>[0-9]*)(?:st|nd|rd|th)|) mixing bowl.$"),
		command: models.CommandFold,
	},
	{
		regexp:  regexp.MustCompile("^Add dry ingredients(?: to(?: (?P<bowl>[0-9]*)(?:st|nd|rd|th)|) mixing bowl|).$"),
		command: models.CommandAddDry,
	},
	{
		regexp:  regexp.MustCompile("^Add (?P<ingredient>[A-Za-z ]*) to(?: (?P<bowl>[0-9]*)(?:st|nd|rd|th)|) mixing bowl.$"),
		command: models.CommandAdd,
	},
	{
		regexp:  regexp.MustCompile("^Add (?P<ingredient>[A-Za-z ]*).$"),
		command: models.CommandAdd,
	},
	{
		regexp:  regexp.MustCompile("^Remove (?P<ingredient>[A-Za-z ]*) from(?: (?P<bowl>[0-9]*)(?:st|nd|rd|th)|) mixing bowl.$"),
		command: models.CommandRemove,
	},
	{
		regexp:  regexp.MustCompile("^Remove (?P<ingredient>[A-Za-z ]*).$"),
		command: models.CommandRemove,
	},
	{
		regexp:  regexp.MustCompile("^Combine (?P<ingredient>[A-Za-z ]*) into(?: (?P<bowl>[0-9]*)(?:st|nd|rd|th)|) mixing bowl.$"),
		command: models.CommandCombine,
	},
	{
		regexp:  regexp.MustCompile("^Combine (?P<ingredient>[A-Za-z ]*).$"),
		command: models.CommandCombine,
	},
	{
		regexp:  regexp.MustCompile("^Divide (?P<ingredient>[A-Za-z ]*) into(?: (?P<bowl>[0-9]*)(?:st|nd|rd|th)|) mixing bowl.$"),
		command: models.CommandDivide,
	},
	{
		regexp:  regexp.MustCompile("^Divide (?P<ingredient>[A-Za-z ]*).$"),
		command: models.CommandDivide,
	},
	{
		regexp:  regexp.MustCompile("^Liquefy contents of the(?: (?P<bowl>[0-9]*)(?:st|nd|rd|th)|) mixing bowl.$"),
		command: models.CommandLiquefyBowl,
	},
	{
		regexp:  regexp.MustCompile("^Liquefy (?P<ingredient>[A-Za-z ]*).$"),
		command: models.CommandLiquefyIngredient,
	},
	{
		regexp:  regexp.MustCompile("^Stir(?: the(?: (?P<bowl>[0-9]*)(?:st|nd|rd|th)|) mixing bowl|) for (?P<minutes>[0-9]*) minutes.$"),
		command: models.CommandStirBowl,
	},
	{
		regexp:  regexp.MustCompile("^Stir (?P<ingredient>[A-Za-z ]*) into the(?: (?P<bowl>[0-9]*)(?:st|nd|rd|th)|) mixing bowl.$"),
		command: models.CommandStirIngredient,
	},
	{
		regexp:  regexp.MustCompile("^Mix(?: the(?: (?P<bowl>[0-9]*)(?:st|nd|rd|th)|) mixing bowl|) well.$"),
		command: models.CommandMix,
	},
	{
		regexp:  regexp.MustCompile("^Clean(?: (?P<bowl>[0-9]*)(?:st|nd|rd|th)|) mixing bowl.$"),
		command: models.CommandClean,
	},
	{
		regexp:  regexp.MustCompile("^Pour contents of the(?: (?P<bowl>[0-9]*)(?:st|nd|rd|th)|) mixing bowl into the(?: (?P<dish>[0-9]*)(?:st|nd|rd|th)|) baking dish.$"),
		command: models.CommandPour,
	},
	{
		regexp:  regexp.MustCompile("^Set aside.$"),
		command: models.CommandSet,
	},
	{
		regexp:  regexp.MustCompile("^Serve with (?P<recipe>[A-Za-z ]*).$"),
		command: models.CommandServe,
	},
	{
		regexp:  regexp.MustCompile("^Refrigerate(?: for (?P<hours>[0-9]*) hours|).$"),
		command: models.CommandRefrigerate,
	},
	{
		regexp:  regexp.MustCompile("^(?P<verbstart>[A-Za-z ]*?)(?:e|y|) the (?P<ingredient>[A-Za-z ]*) until (?P<verbend>[A-Za-z ]*?)(?:ied|ed).$"),
		command: models.CommandVerbEnd,
	},
	{
		regexp:  regexp.MustCompile("^(?P<verbstart>[A-Za-z ]*?)(?:e|y|) the (?P<ingredient>[A-Za-z ]*).$"),
		command: models.CommandVerbStart,
	},
}
