package day5

import (
	"fmt"
	"strconv"
	"strings"
)

var cheat []string

func HandleFirst(input []string) int {
	cheat = input
	ruleSet, updates := retrieveRulesAndUpdates(input)

	validUpdates := make([]int, 0)

	for i, updatePages := range updates {
		fmt.Printf("Processing Update %v (%v): ", i, updatePages)
		if CanPrintUpdate(ruleSet, updatePages) {
			validUpdates = append(validUpdates, i)
		}
		fmt.Print("\n")
	}

	sum := 0
	for i := range validUpdates {
		pages := len(updates[i])
		halfCut := pages / 2

		fmt.Printf("Retrieving page %v from update (%v)\n", updates[i][halfCut], updates[i])

		sum += updates[i][halfCut]
	}

	return sum
}

func HandleSecond(input []string) int {
	return 0
}

type ruleSetMap map[int]updatePageList

type updatePageList []int

func CanPrintUpdate(ruleSet ruleSetMap, updatePages []int) bool {
	for i, updatePage := range updatePages {
		if !ruleSet.IsAllowed(updatePage, updatePages[i+1:]) {
			return false
		}
	}

	fmt.Print("Update success!")

	return true
}

func (ruleSet ruleSetMap) IsAllowed(pageNumber int, remainingPages []int) bool {
	for _, page := range remainingPages {
		if ruleSet[page].Contains(pageNumber) {
			fmt.Printf("Update Failed! page %v cannot be before %v (%v|%v)", pageNumber, page, page, pageNumber)

			rule := fmt.Sprintf("%v|%v", page, pageNumber)

			found := false

			for _, line := range cheat {
				if rule == line {
					found = true
				}
			}

			if !found {
				fmt.Printf("\nCheat %v not found in cheat list!\n", rule)
			}

			return false
		}
	}

	return true
}

func (list updatePageList) Contains(value int) bool {
	for _, page := range list {
		if page == value {
			return true
		}
	}

	return false
}

func retrieveRulesAndUpdates(input []string) (ruleSetMap, [][]int) {
	// A white line (empty string) splits the rules and updates
	index := -1
	for i, line := range input {
		if line == "" {
			index = i
		}
	}

	return retrieveRuleSet(input[:index]), retrieveUpdates(input[index+1:])
}

func retrieveUpdates(updates []string) [][]int {
	updateList := make([][]int, len(updates))

	for i, update := range updates {
		updatePages := strings.Split(update, ",")
		updateList[i] = make([]int, len(updatePages))

		for j, updatePage := range updatePages {
			updateList[i][j], _ = strconv.Atoi(updatePage)
		}
	}

	return updateList
}

func retrieveRuleSet(rules []string) ruleSetMap {
	ruleSet := make(ruleSetMap)
	for _, rule := range rules {
		splitRules := strings.Split(rule, "|")
		ruleBefore, _ := strconv.Atoi(splitRules[0])
		ruleAfter, _ := strconv.Atoi(splitRules[1])

		ruleSet[ruleBefore] = append(ruleSet[ruleBefore], ruleAfter)
	}

	return ruleSet
}
