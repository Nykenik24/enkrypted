package repl

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

type Handler func(args []string) error

type Command struct {
	Name  string
	Usage string
	Func  Handler
}

func NewCommand(name, usage string, fn Handler) *Command {
	return &Command{
		Name:  name,
		Usage: usage,
		Func:  fn,
	}
}

type REPL struct {
	Prompt   string
	Commands map[string]*Command
	stdout   *bufio.Writer
}

func NewREPL(prompt string) *REPL {
	return &REPL{
		Prompt:   prompt,
		Commands: make(map[string]*Command),
		stdout:   bufio.NewWriter(os.Stdout),
	}
}

func (r *REPL) AddCommand(cmd *Command) {
	r.Commands[cmd.Name] = cmd
}

func (r *REPL) Handle(args []string) error {
	if len(args) == 0 {
		return errors.New("no args")
	}

	if args[0] == "help" {
		r.help()
		return nil
	}

	command, exists := r.Commands[args[0]]
	if !exists {
		return fmt.Errorf("no handler for %s", args[0])
	}

	return command.Func(args)
}

func (r *REPL) help() {
	fmt.Println("\nCommands for enkrypted:")
	for _, cmd := range r.Commands {
		fmt.Printf("\t\x1b[1m%s\x1b[0m - \x1b[3m%s\x1b[0m\n", cmd.Name, cmd.Usage)
	}
	fmt.Println()
}

func (r *REPL) Run() error {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Printf("\nReceived %s, exiting...\n", sig)
		os.Exit(0)
	}()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(r.Prompt)

		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return errors.New("\nEOF received, exiting...")
			}
			return fmt.Errorf("error reading input: %v", err)
		}

		input := strings.TrimRight(line, "\r\n")
		if input == "" {
			continue
		}

		args := strings.Split(input, " ")
		if err := r.Handle(args); err != nil {
			fmt.Printf("\n\x1b[31mError when running command:\x1b[0m \n\x1b[90m---->\x1b[0m %v\n\n", err)
		}
	}
}
