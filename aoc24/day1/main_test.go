package main

import (
	"slices"
	"strings"
	"testing"
)

var testData = `3   4
4   3
2   5
1   3
3   9
3   3
`

func TestParse(t *testing.T) {
	expectedArray1 := []int{3, 4, 2, 1, 3, 3}
	expectedArray2 := []int{4, 3, 5, 3, 9, 3}

	testReader := strings.NewReader(testData)

	resultArray1, resultArray2 := parse(testReader)

	if resultArray1 == nil || resultArray2 == nil {
		t.Errorf("both or a single array are nil: array1=%v, array2=%v", resultArray1, resultArray2)
	}

	if !slices.Equal(resultArray1, expectedArray1) {
		t.Errorf("resultArray1 does not contain the expected elements: result=%v, expected=%v", resultArray1, expectedArray1)
	}

	if !slices.Equal(resultArray2, expectedArray2) {
		t.Errorf("resultArray2 does not contain the expected elements: result=%v, expected=%v", resultArray2, expectedArray2)
	}
}

func TestSortLocations(t *testing.T) {
	locationsToSort := []int{5, 4, 3, 2, 1}
	expectedOrder := []int{1, 2, 3, 4, 5}

	sortLocations(locationsToSort)

	if !slices.Equal(locationsToSort, expectedOrder) {
		t.Errorf("list was not sorted as expected: sortedList=%v, expectedLists=%v", locationsToSort, expectedOrder)
	}
}

func TestDistance(t *testing.T) {
	expectedDistance := 11

	testReader := strings.NewReader(testData)

	locations1, locations2 := parse(testReader)
	sortLocations(locations1)
	sortLocations(locations2)

	resultDistance := distance(locations1, locations2)

	if resultDistance != expectedDistance {
		t.Errorf("calculated distance was not what we expected: result=%d, expected=%d", resultDistance, expectedDistance)
	}
}

func TestSimularity(t *testing.T) {
	expectedScore := 31

	testReader := strings.NewReader(testData)

	locations1, locations2 := parse(testReader)

	resultScore := similarity(locations1, locations2)

	if resultScore != expectedScore {
		t.Errorf("calculated similarity score is not expected: score=%d, expected=%d, locations1=%v, loctions2=%v", resultScore, expectedScore, locations1, locations2)
	}

}

func TestCountMatches(t *testing.T) {
	locations := []int{1, 1, 1, 2}
	expectedMatches := 3
	matchValue := 1

	resultMatches := countMatches(matchValue, locations)

	if resultMatches != expectedMatches {
		t.Errorf("count of matches is not expected: result=%d, expected=%d, matching=%d, locations=%d", resultMatches, expectedMatches, matchValue, locations)
	}
}
