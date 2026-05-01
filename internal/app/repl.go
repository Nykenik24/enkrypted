package enkrypt

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type CommandHandler func(args []string, ctx *CmdContext) error

type CmdContext struct {
	App *App
}

type Command struct {
	Name  string
	Usage string
	Func  CommandHandler
}

func NewCommand(name, usage string, fn CommandHandler) *Command {
	return &Command{
		Name:  name,
		Usage: usage,
		Func:  fn,
	}
}

type REPL struct {
	Prompt   string
	Commands map[string]*Command
	Aliases  map[string]string
	stdout   *bufio.Writer
}

func NewREPL(prompt string) *REPL {
	return &REPL{
		Prompt:   prompt,
		Commands: make(map[string]*Command),
		Aliases:  make(map[string]string),
		stdout:   bufio.NewWriter(os.Stdout),
	}
}

func (r *REPL) AddCommand(cmd *Command) {
	r.Commands[cmd.Name] = cmd
}

func (r *REPL) Alias(original, alias string) {
	r.Aliases[alias] = original
}

func (r *REPL) Handle(args []string, ctx *CmdContext) error {
	if len(args) == 0 {
		return errors.New("no args")
	}

	if args[0] == "help" {
		r.help()
		return nil
	}

	name := args[0]
	if _, aliased := r.Aliases[name]; aliased {
		name = r.Aliases[name]
	}

	command, exists := r.Commands[name]
	if !exists {
		return fmt.Errorf("no handler for %s", args[0])
	}

	return command.Func(args, ctx)
}

func (r *REPL) help() {
	fmt.Println("Commands for enkrypted:")

	// might want to sort these later
	for _, cmd := range r.Commands {
		fmt.Printf("\t\x1b[1m%s\x1b[0m - \x1b[3m%s\x1b[0m\n", cmd.Name, cmd.Usage)
	}

	if len(r.Aliases) > 0 {
		fmt.Println()
		fmt.Println("Aliases for enkrypted:")
		for alias, og := range r.Aliases {
			fmt.Printf("\t\x1b[1m%s\x1b[0m - \x1b[3m%s\x1b[0m\n", alias, og)
		}
	}
}

func (r *REPL) Run(app *App) error {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Printf("\nReceived %s, exiting...\n", sig)
		os.Exit(0)
	}()

	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("\nReceived EOF, exiting...")
				return nil
			}
			return fmt.Errorf("error reading input: %v", err)
		}

		input := strings.TrimRight(line, "\r\n")
		if input == "" {
			continue
		}

		fmt.Println()
		args := strings.Split(input, " ")
		if err := r.Handle(args, &CmdContext{App: app}); err != nil {
			fmt.Printf("\n\x1b[31mError when running command:\x1b[0m \n\x1b[90m---->\x1b[0m %v\n\n", err)
		}
		fmt.Println()

		fmt.Print(r.Prompt)
	}
}
