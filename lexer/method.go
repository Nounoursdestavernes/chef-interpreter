package lexer

import (
	"github.com/Nounoursdestavernes/chef-interpreter/models"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func (l Lexer) tokenizeMethod(input string) (method []models.MethodStatement, err error) {
	if !strings.HasPrefix(input, "Method") {
		fmt.Println(input)
		return []models.MethodStatement{}, fmt.Errorf("could not find Method field")
	}

	input = strings.Replace(input, "Method.\n", "", 1)
	statements := strings.SplitAfter(input, ".")

	for _, statement := range statements {
		statement = strings.TrimSpace(statement)
		if statement == "" {
			continue
		}

		isValid := false

		for _, command := range commandOptions {
			tokens := command.regexp.FindStringSubmatch(statement)

			if len(tokens) > 0 {
				isValid = true

				methodStatement, err := populateMethodStatement(command.command, command.regexp.SubexpNames()[1:], tokens[1:])
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

	l.markFieldComplete()
	return method, nil
}

func printMethod(method []models.MethodStatement) {
	for _, statement := range method {
		fmt.Printf("OPCODE: %d\n\tINGREDIENT: %s\n\tMIXING BOWL: %d\n\tBAKING DISH: %d\n\tMINUTES: %d\n\tHOURS: %d\n\tVERB: %s\n\tRECIPE: %s\n\n",
			statement.Command,
			statement.Ingredient,
			statement.MixingBowl,
			statement.BakingDish,
			statement.Minutes,
			statement.Hours,
			statement.Verb,
			statement.Recipe,
		)
	}
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
		case "verb":
			methodStatement.Verb = fieldValues[field]
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
		regexp:  regexp.MustCompile("^Take (?P<ingredient>[A-Za-z ]*?) from(?: the|) refrigerator.$"),
		command: models.CommandTake,
	},
	{
		regexp:  regexp.MustCompile("^Put (?P<ingredient>[A-Za-z ]*) into(?: the|)(?: (?P<bowl>[0-9]*)(?:st|nd|rd|th)|) mixing bowl.$"),
		command: models.CommandPut,
	},
	{
		regexp:  regexp.MustCompile("^Fold (?P<ingredient>[A-Za-z ]*) into(?: the|)(?: (?P<bowl>[0-9]*)(?:st|nd|rd|th)|) mixing bowl.$"),
		command: models.CommandFold,
	},
	{
		regexp:  regexp.MustCompile("^Add dry ingredients(?: to(?: the|)(?: (?P<bowl>[0-9]*)(?:st|nd|rd|th)|) mixing bowl|).$"),
		command: models.CommandAddDry,
	},
	{
		regexp:  regexp.MustCompile("^Add (?P<ingredient>[A-Za-z ]*) to(?: the|)(?: (?P<bowl>[0-9]*)(?:st|nd|rd|th)|) mixing bowl.$"),
		command: models.CommandAdd,
	},
	{
		regexp:  regexp.MustCompile("^Add (?P<ingredient>[A-Za-z ]*).$"),
		command: models.CommandAdd,
	},
	{
		regexp:  regexp.MustCompile("^Remove (?P<ingredient>[A-Za-z ]*) from(?: the|)(?: (?P<bowl>[0-9]*)(?:st|nd|rd|th)|) mixing bowl.$"),
		command: models.CommandRemove,
	},
	{
		regexp:  regexp.MustCompile("^Remove (?P<ingredient>[A-Za-z ]*).$"),
		command: models.CommandRemove,
	},
	{
		regexp:  regexp.MustCompile("^Combine (?P<ingredient>[A-Za-z ]*) into(?: the|)(?: (?P<bowl>[0-9]*)(?:st|nd|rd|th)|) mixing bowl.$"),
		command: models.CommandCombine,
	},
	{
		regexp:  regexp.MustCompile("^Combine (?P<ingredient>[A-Za-z ]*).$"),
		command: models.CommandCombine,
	},
	{
		regexp:  regexp.MustCompile("^Divide (?P<ingredient>[A-Za-z ]*) into(?: the|)(?: (?P<bowl>[0-9]*)(?:st|nd|rd|th)|) mixing bowl.$"),
		command: models.CommandDivide,
	},
	{
		regexp:  regexp.MustCompile("^Divide (?P<ingredient>[A-Za-z ]*).$"),
		command: models.CommandDivide,
	},
	{
		regexp:  regexp.MustCompile("^Liqu[ei]fy(?: the|) contents of the(?: (?P<bowl>[0-9]*)(?:st|nd|rd|th)|) mixing bowl.$"),
		command: models.CommandLiquefyBowl,
	},
	{
		regexp:  regexp.MustCompile("^Liqu[ei]fy(?: the|) (?P<ingredient>[A-Za-z ]*).$"),
		command: models.CommandLiquefyIngredient,
	},
	{
		regexp:  regexp.MustCompile("^Stir(?: the(?: (?P<bowl>[0-9]*)(?:st|nd|rd|th)|) mixing bowl|) for (?P<minutes>[0-9]*) minute(?:s|).$"),
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
		regexp:  regexp.MustCompile("^Clean(?: the|)(?: (?P<bowl>[0-9]*)(?:st|nd|rd|th)|) mixing bowl.$"),
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
		regexp:  regexp.MustCompile("^Refrigerate(?: for (?P<hours>[0-9]*) hour(?:s|)|).$"),
		command: models.CommandRefrigerate,
	},
	{
		regexp:  regexp.MustCompile("^(?:[A-Za-z ]*?)(?: the (?P<ingredient>[A-Za-z ]*)|) until (?P<verb>[A-Za-z ]*?)(?:ied|ed).$"),
		command: models.CommandVerbEnd,
	},
	{
		regexp:  regexp.MustCompile("^(?P<verb>[A-Za-z ]*?)(?:e|y|) the (?P<ingredient>[A-Za-z ]*).$"),
		command: models.CommandVerbStart,
	},
}
