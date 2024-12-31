package main

import (
	"fmt"
	"log/slog"
	"os"
	"utility"

	"github.com/spf13/cobra"

	"github.com/bcheldelin/go-advent-of-code/aoc24/day7"
)

var (
	// Used for flags.
	day     int
	part    int
	example bool

	rootCmd = &cobra.Command{
		Use: "aoc24",
		Run: RunChallange,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&day, "day", "d", 0, "specify challange to run")
	rootCmd.PersistentFlags().IntVarP(&part, "part", "p", 1, "Specify challange part")
	rootCmd.PersistentFlags().BoolVarP(&example, "example", "e", false, "use example input")
}

func RunChallange(cmd *cobra.Command, args []string) {

	var filename string
	if example {
		filename = "input.txt"
	} else {
		filename = "input-example.txt"
	}

	input := utility.GetInputStringByPath(fmt.Sprintf("../day%d/%s", day, filename))

	if input == "" {
		slog.Error("Somthing went wrong, we didn't get anything for input.")
	}

	switch day {
	case 7:
		day7.Challange(input, part)
	}

}

func main() {
	if err := Execute(); err != nil {
		os.Exit(1)
	}
}
