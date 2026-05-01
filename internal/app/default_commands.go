package enkrypt

import (
	"fmt"
	"os"
	"strings"

	"github.com/Nykenik24/enkrypted/internal/config"
	"github.com/Nykenik24/enkrypted/internal/util"
)

func getterCommand(name, desc, valueName string, value func(ctx *CmdContext) any) *Command {
	return NewCommand(
		name,
		desc,
		func(args []string, ctx *CmdContext) error {
			fmt.Printf("\x1b[34m%s: \x1b[0m%s\n", valueName, value(ctx))
			return nil
		},
	)
}

var defaultCommands = []*Command{
	NewCommand(
		"quit",
		"quit the program",
		func(args []string, ctx *CmdContext) error {
			msgs := []string{
				"See you later!",
				"Goodbye!",
				"Bye bye!",
				"Hope you come back!",
			}
			msg := util.Choice(msgs)
			fmt.Printf("\"%s\"\x1b[90;3m\nAtt: Enkrypted\x1b[0m\n", msg)
			os.Exit(0)
			return nil
		},
	),
	NewCommand(
		"routes",
		"get all routes",
		func(args []string, ctx *CmdContext) error {
			routes := ctx.App.fiber.GetRoutes()
			for _, route := range routes {
				path := route.Path
				path = strings.ReplaceAll(path, "/", "\x1b[90m/\x1b[0m")
				path = strings.ReplaceAll(path, ":", "\x1b[33m:\x1b[0m")

				fmt.Printf("\x1b[1mRoute\x1b[0m %s\n", path)
				prefix := "└── "
				if len(route.Params) > 0 {
					prefix = "├── "
				}
				fmt.Printf("%s\x1b[90mMethod\x1b[0m: \x1b[34m%s\x1b[0m\n", prefix, route.Method)
				if len(route.Params) > 0 {
					prefix = "└── "
					fmt.Printf("%s\x1b[90mParams\x1b[0m: %s\n", prefix, strings.Join(route.Params, ", "))
				}
				fmt.Println()
			}

			return nil
		},
	),
	getterCommand("addr", "get the server's address", "Addresss", func(_ *CmdContext) any {
		return config.ADDR
	}),
	getterCommand("version", "get the server's version", "Version", func(_ *CmdContext) any {
		return config.VERSION
	}),
}

var defaultAlias = map[string]string{
	"q": "quit",
	"v": "version",
}

func (r *REPL) RegisterDefault() {
	for _, cmd := range defaultCommands {
		r.AddCommand(cmd)
	}

	for alias, og := range defaultAlias {
		r.Alias(og, alias)
	}
}
