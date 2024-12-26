package main

import (
	"slices"
	"testing"
	"utility"
)

const EXAMPLE = `
47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func (r UpdateRules) Equal(o UpdateRules) bool {
	if len(r) != len(o) {
		return false
	}

	for i := range r {
		if !slices.Equal(r[i], o[i]) {
			return false
		}
	}

	return true
}

func TestInputSplit(t *testing.T) {
	testInput := "47|53\n97|13\n\n75,47,61,53,29\n97,61,53,29,13"
	expectedRules := "47|53\n97|13"
	expectedUpdates := "75,47,61,53,29\n97,61,53,29,13"

	resultRules, resultUpdates := inputSplit(testInput)

	if resultRules != expectedRules {
		t.Errorf("Did not get expected rules: got %s, expected %s", resultRules, expectedRules)
	}

	if resultUpdates != expectedUpdates {
		t.Errorf("Did not get expected updates: got %s, expected %s", resultUpdates, expectedUpdates)
	}
}

func TestNewRules(t *testing.T) {
	testInput := "47|53\n47|60\n97|13"

	expectedRules := UpdateRules{
		47: []int{53, 60},
		97: []int{13},
	}

	resultRules := newRules(testInput)

	if !resultRules.Equal(expectedRules) {
		t.Errorf("Result rules is not what we expected")
	}
}

func TestNewUpdates(t *testing.T) {
	inputUpdates := "75,47,61,53,29\n97,61,53,29,13"
	expectedUpdates := [][]int{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
	}

	resultUpdates := newUpdates(inputUpdates)

	if len(expectedUpdates) != len(resultUpdates) {
		t.Error("Did not get the expected result")
	}
}

func TwoDSliceEqual(a [][]int, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, _ := range a {
		for j, _ := range b {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

type CheckUpdatesTest struct {
	rules  UpdateRules
	update []int
	expect bool
}

func TestCheckUpdate(t *testing.T) {
	tests := map[string]CheckUpdatesTest{
		"Good": {
			UpdateRules{47: {22, 53}},
			[]int{47, 22, 53},
			true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if result, _, _ := checkUpdate(test.rules, test.update); result != test.expect {
				t.Errorf("check update got unexpected result: got (%t) expected (%t)", result, test.expect)
			}
		})
	}
}

func BenchmarkCheckUpdate(b *testing.B) {
	input := utility.InputString()

	inputRules, updateInput := inputSplit(string(input))

	rules := newRules(inputRules)
	updates := newUpdates(updateInput)

	// b.Run("Benchmark Without Pointer", func(b *testing.B) {
	// 	for _, update := range updates {
	// 		checkUpdateNoPointer(rules, update)
	// 	}
	// })
	b.Run("Benchmark Pointer", func(b *testing.B) {
		for _, update := range updates {
			checkUpdate(rules, update)
		}
	})
}

func TestFilterUpdates(t *testing.T) {
	testInput := "47|53\n97|13\n\n75,47,61,53,29\n97,61,53,29,13\n53,47,29"
	expectedUpdates := [][]int{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
	}

	inputRules, inputUpdates := inputSplit(testInput)
	rules := newRules(inputRules)
	updates := newUpdates(inputUpdates)
	resultUpdates := filterUpdates(rules, updates, true)

	if !TwoDSliceEqual(resultUpdates, expectedUpdates) {
		t.Error("Did not get what we expected!")
	}

}

func TestSumMedians(t *testing.T) {
	updates := [][]int{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
	}

	expectedSum := 114

	if resultSum := sumMedians(updates); resultSum != expectedSum {
		t.Errorf("Incorrect expected sum: expected %d got %d", expectedSum, resultSum)
	}
}

func TestFixUpdate(t *testing.T) {

	rulesInput := "47|53\n97|13\n97|61\n97|47\n75|29\n61|13\n75|53\n29|13\n97|29\n53|29\n61|53\n97|53\n61|29\n47|13\n75|47\n97|75\n47|61\n75|61\n47|29\n75|13\n53|13"
	rules := newRules(rulesInput)

	tests := []struct {
		name   string
		update []int
		expect []int
	}{
		{"example 1", []int{75, 97, 47, 61, 53}, []int{97, 75, 47, 61, 53}},
		{"example 1", []int{61, 13, 29}, []int{61, 29, 13}},
		{"example 1", []int{97, 13, 75, 29, 47}, []int{97, 75, 47, 29, 13}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if result := fixUpdate(rules, test.update); !slices.Equal(result, test.expect) {
				t.Errorf("result was not expected: expected=%v got=%v", test.expect, result)
			}
		})
	}

}
