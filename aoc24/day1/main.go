package main

import (
	"bufio"
	"io"
	"log/slog"
	"math"
	"sort"
	"strconv"
	"strings"

	"utility"
)

func parse(dataInput io.Reader) (array1 []int, array2 []int) {
	scanner := bufio.NewScanner(dataInput)

	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, "   ")

		// handle array1 element
		intElem1, err := strconv.Atoi(lineSplit[0])
		if err != nil {
			slog.Error("Unable to convert string to int for element1", "error", err)
		}
		array1 = append(array1, intElem1)

		// handle array2 element
		intElem2, err := strconv.Atoi(lineSplit[1])
		if err != nil {
			slog.Error("Unable to convert string to int for element2", "error", err)
		}
		array2 = append(array2, intElem2)
	}

	return
}

func sortLocations(locations []int) {
	sort.Ints(locations)
}

func distance(locations1 []int, locations2 []int) (distance int) {
	distance = 0

	len1 := len(locations1)

	for i := 0; i < len1; i++ {
		f1 := float64(locations1[i])
		f2 := float64(locations2[i])

		r1 := f1 - f2

		d1 := math.Abs(r1)

		distance += int(d1)
	}

	return
}

func similarity(locations1 []int, locations2 []int) (score int) {
	score = 0
	len1 := len(locations1)

	for i := 0; i < len1; i++ {
		element := locations1[i]
		score += element * countMatches(element, locations2)
	}
	return
}

func countMatches(match int, locations []int) (count int) {
	count = 0
	for i := 0; i < len(locations); i++ {
		if locations[i] == match {
			count++
		}
	}
	return
}

func main() {
	inputReader := utility.GetInputReader()
	defer inputReader.Close()

	locations1, locations2 := parse(inputReader)

	if locations1 == nil || locations2 == nil {
		slog.Error("locations data is empty")
	}

	sortLocations(locations1)
	sortLocations(locations2)

	distance := distance(locations1, locations2)
	slog.Info("total distance calculated", "distance", distance)

	score := similarity(locations1, locations2)
	slog.Info("similarity score calculated", "score", score)
}
