package cmd

import (
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

type CMD struct {
	day     int
	part    int
	example bool
	root    *cobra.Command
}

func NewAOCCmd(year string, challanges map[int]func(input string, part int)) *CMD {

	var AocCMD = CMD{}

	root := &cobra.Command{
		Use: year,
		Run: func(cmd *cobra.Command, args []string) {

			var filename string
			if AocCMD.example {
				slog.Info("using example input")
				filename = "input-example.txt"
			} else {
				filename = "input.txt"
			}

			input := getInputStringByPath(fmt.Sprintf("./day%d/%s", AocCMD.day, filename))

			if input == "" {
				slog.Error("Somthing went wrong, we didn't get anything for input.")
				os.Exit(1)
			}

			challanges[AocCMD.day](input, AocCMD.part)

		},
	}

	root.PersistentFlags().IntVarP(&AocCMD.day, "day", "d", 0, "specify challange to run")
	root.PersistentFlags().IntVarP(&AocCMD.part, "part", "p", 1, "Specify challange part")
	root.PersistentFlags().BoolVarP(&AocCMD.example, "example", "e", false, "use example input")

	AocCMD.root = root

	return &AocCMD
}

func (c *CMD) Execute() error {
	return c.root.Execute()
}

func newInputReader(filename string) *os.File {
	inputReader, err := os.Open(filename)

	if err != nil {
		slog.Error("got an error while trying to open input file", "Error", err)
	}

	return inputReader
}

func getInputStringByPath(path string) string {
	inputFile := newInputReader(path)
	defer inputFile.Close()

	if b, err := io.ReadAll(inputFile); err == nil {
		return string(b)
	}

	return ""
}
