package challenge

import (
	"aoc/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	challenges = map[string]func(data string) string{
		"0101": Challenge0101,
		"0102": Challenge0102,
		"0201": Challenge0201,
		"0202": Challenge0202,
	}

	// Used for flags.
	challenge string
	data      string

	rootCmd = &cobra.Command{
		Use:   "aoc",
		Short: "Advent of Code solver",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 && challenge == "" {
				fmt.Printf("You need to provide the challenge")
				return
			}

			if challenge == "" && len(args) == 1 {
				challenge = args[0]
			}

			if action := challenges[challenge]; action != nil {
				fmt.Printf("%s\n", action(data))
			}

		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&challenge, "challenge", "", "challenge")
	rootCmd.PersistentFlags().StringVar(&data, "dataFile", "d", "data file to load")
	hasInput := utils.HasStdinInput()
	if hasInput {
		data = utils.GetStdinString()
	}

}
