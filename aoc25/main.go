package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof" // Import for pprof HTTP handlers
	"os"

	"github.com/briancheldelin/go-advent-of-code/aoc25/day1"
	"github.com/briancheldelin/go-advent-of-code/aoc25/day2"
	"github.com/briancheldelin/go-advent-of-code/aoc25/day3"

	"github.com/briancheldelin/go-advent-of-code/utilities/cmd"
)

func main() {

	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	var challenges = map[int]func(input string, part int){}

	challenges[1] = day1.RunChallange
	challenges[2] = day2.RunChallange
	challenges[3] = day3.RunChallange

	cmd := cmd.NewAOCCmd("aoc25", challenges)

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
