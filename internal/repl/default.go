package repl

import (
	"fmt"
	"os"

	"github.com/Nykenik24/enkrypted/internal/util"
)

var defaultCommands = []*Command{
	NewCommand(
		"quit",
		"quit the program",
		func(args []string) error {
			msgs := []string{
				"See you later!",
				"Goodbye!",
				"Bye bye!",
				"Hope you come back!",
			}
			msg := util.Choice(msgs)
			fmt.Println()
			fmt.Printf("\"%s\"\x1b[90;3m\nAtt: Enkrypted\x1b[0m\n", msg)
			os.Exit(0)
			return nil
		},
	),
}

func (r *REPL) RegisterDefault() {
	for _, cmd := range defaultCommands {
		r.AddCommand(cmd)
	}
}
