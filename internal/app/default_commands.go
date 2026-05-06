package enkrypt

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/Nykenik24/enkrypted/internal/config"
	"github.com/Nykenik24/enkrypted/internal/models"
	"github.com/Nykenik24/enkrypted/internal/repository"
	"golang.org/x/crypto/ssh/terminal"
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
			os.Exit(0)
			return nil
		},
	),
	NewCommand(
		"routes",
		"get all routes",
		func(args []string, ctx *CmdContext) error {
			routes := ctx.App.fiber.GetRoutes()

			if len(routes) == 0 {
				fmt.Println("\x1b[91mNo routes\x1b[0m")
				return nil
			}

			for _, route := range routes {
				path := route.Path
				path = strings.ReplaceAll(path, "/", "\x1b[33m/\x1b[0m")
				path = strings.ReplaceAll(path, ":", "\x1b[33m:\x1b[0m")
				path = "\x1b[90m" + *config.ADDR + "\x1b[0m" + path

				fmt.Printf("\x1b[34m%s\x1b[0m %s\n", route.Method, path)

				if len(route.Params) > 0 {
					for i, param := range route.Params {
						prefix := "├──"
						if i == len(route.Params)-1 {
							prefix = "└──"
						}
						fmt.Printf("%s Params[\x1b[35m%d\x1b[0m]: \x1b[32m\"%s\"\x1b[0m\n", prefix, i, param)
					}
				}
				fmt.Println()
			}

			return nil
		},
	),
	NewCommand(
		"clients",
		"get all connected clients",
		func(args []string, ctx *CmdContext) error {
			repo := repository.GlobalClientRepo()
			clients := repo.GetAll()

			short := false
			if len(args) > 1 && args[1] == "short-id" {
				short = true
			}

			if len(clients) == 0 {
				fmt.Println("\x1b[91mNo clients\x1b[0m")
				return nil
			}

			fmt.Println("\x1b[33mConnected clients\x1b[0m")
			i := 0
			count, err := repo.Count()
			if err != nil {
				return err
			}
			for _, client := range clients {
				prefix := "├──"
				if i == count-1 {
					prefix = "└──"
				}
				id := client.ID.Bytes()
				if short {
					id = client.ID.Short()
				}
				fmt.Printf("%s \x1b[1mClient\x1b[0m \x1b[34m%s\x1b[0m\n", prefix, id)
				i++
			}
			fmt.Println()

			return nil
		},
	),
	NewCommand(
		"new-room",
		"create a new room",
		func(args []string, ctx *CmdContext) error {
			fmt.Println("\x1b[32mCreate a new room\n----------\x1b[0m")

			stars := func(in string) string {
				str := ""
				for range in {
					str = str + "*"
				}
				return str
			}
			var (
				pass []byte
				err  error
			)
			for {
				fmt.Print("\x1b[34mPassword\x1b[0m: ")
				pass, err = terminal.ReadPassword(0)
				if err != nil {
					return err
				}

				fmt.Println(stars(string(pass)))

				if len(pass) < 8 {
					fmt.Println("\x1b[31mPassword must be at least 8 characters long\x1b[0m")
					continue
				}

				break
			}

			fmt.Print("\x1b[34mConfirm password\x1b[0m: ")
			confirm, err := terminal.ReadPassword(0)
			if err != nil {
				return err
			}

			fmt.Println(stars(string(confirm)))

			if slices.Compare(confirm, pass) != 0 {
				fmt.Print("\n\x1b[31mPassword doesn't match first entry\x1b[0m")
				return nil
			}

			fmt.Println()
			fmt.Println("Creating new room...")

			repo := repository.NewRoomRepo()
			repo.Create(models.Room{
				Password: string(confirm),
			})

			return nil
		},
	),
	getterCommand("addr", "get the server's address", "Addresss", func(_ *CmdContext) any {
		return *config.ADDR
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
