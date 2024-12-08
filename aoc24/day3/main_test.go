package main

import (
	"strings"
	"testing"
)

const TEST_STRING = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

func TestFind(t *testing.T) {
	inputReader := strings.NewReader(TEST_STRING)

	foundOperations := findOps(inputReader, FIND_OPS_PART1)

	if len(foundOperations) != 4 {
		t.Error("did not get any hits when we expected some")
	}

}

func TestParse(t *testing.T) {
	inputReader := strings.NewReader(TEST_STRING)
	foundOperations := findOps(inputReader, FIND_OPS_PART1)
	multiplicationSets := parseMultiply(foundOperations)

	if len(multiplicationSets) != 4 {
		t.Error("did not get any multiplication sets")
	}
}

func TestMultiplySum(t *testing.T) {
	inputReader := strings.NewReader(TEST_STRING)
	foundOperations := findOps(inputReader, FIND_OPS_PART1)
	multiplicationSets := parseMultiply(foundOperations)

	multplicationSum := multiplySum(multiplicationSets)

	if multplicationSum == 0 {
		t.Error("did not get an expected result")
	}
}

func TestMultiply(t *testing.T) {
	testOp := []int{1, 2}
	expected := 2

	result := multiply(testOp)

	if result != expected {
		t.Errorf("did not get expected result: result=%d, expected=%d", result, expected)
	}
}
