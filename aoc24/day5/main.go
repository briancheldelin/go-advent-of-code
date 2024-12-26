package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"utility"
)

type UpdateRules map[int][]int

func inputSplit(input string) (pageRules, pageUpdates string) {
	r := strings.Split(input, "\n\n")
	return r[0], r[1]
}

func newRules(input string) UpdateRules {

	rulesStrings := strings.Split(input, "\n")
	rules := make(UpdateRules)

	for _, ruleString := range rulesStrings {
		rulePair := strings.Split(ruleString, "|")
		pageBefore, _ := strconv.Atoi(rulePair[0])
		pageAfter, _ := strconv.Atoi(rulePair[1])
		if rules[pageBefore] == nil {
			rules[pageBefore] = make([]int, 0, 1)
		}
		rules[pageBefore] = append(rules[pageBefore], pageAfter)
	}

	return rules
}

func newUpdates(input string) (result [][]int) {
	lines := strings.Split(input, "\n")

	result = make([][]int, len(lines))

	for i, line := range lines {
		elements := strings.Split(line, ",")
		result[i] = make([]int, len(elements))
		for k, element := range elements {
			elementInt, _ := strconv.Atoi(element)
			result[i][k] = elementInt
		}
	}

	return result
}

func fixUpdate(rules UpdateRules, update []int) []int {

	for true {
		result, updateElement, rule := checkUpdate(rules, update)
		if result == true {
			break
		}
		elementIndex := slices.Index(update, updateElement)
		ruleIndex := slices.Index(update, rule)

		update[elementIndex] = rule
		update[ruleIndex] = updateElement
	}

	return update

}

func checkUpdate(rules UpdateRules, update []int) (bool, int, int) {
	for updateIndex, updateElement := range update {
		for _, rule := range rules[updateElement] {
			if slices.Contains(update[:updateIndex], rule) {
				return false, rule, updateElement
			}
		}
	}
	return true, 0, 0
}

func filterUpdates(rules UpdateRules, updates [][]int, want bool) [][]int {
	filteredUpdates := make([][]int, 0, len(updates)/2) // start with half of len()

	for _, update := range updates {
		if result, _, _ := checkUpdate(rules, update); result == want {
			filteredUpdates = append(filteredUpdates, update)
		}
	}

	return filteredUpdates
}

func sumMedians(updates [][]int) (sum int) {
	for _, update := range updates {
		sum += update[len(update)/2]
	}
	return sum
}

func fixUpdates(rules UpdateRules, updates [][]int) [][]int {
	for i, update := range updates {
		updates[i] = fixUpdate(rules, update)
	}
	return updates
}

func main() {

	part := os.Args[1]

	input := utility.InputString()

	inputRules, inputUpdates := inputSplit(string(input))

	rules := newRules(inputRules)
	updates := newUpdates(inputUpdates)

	if part == "part1" {
		updates = filterUpdates(rules, updates, true)
	} else if part == "part2" {
		updates = filterUpdates(rules, updates, false)
		updates = fixUpdates(rules, updates)
	}
	sum := sumMedians(updates)

	fmt.Printf("Sum for %s update medieums: sum=%d\n", part, sum)
}
