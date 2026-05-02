package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

type command struct {
	name    string
	summary string
	usage   string
	run     func(args []string) error
}

var commands = map[string]command{}

func registerCommand(commandConfig command) {
	if commandConfig.name == "" {
		panic("command name is required")
	}
	if commandConfig.run == nil {
		panic(fmt.Sprintf("command %q requires a run function", commandConfig.name))
	}
	if _, exists := commands[commandConfig.name]; exists {
		panic(fmt.Sprintf("command %q is already registered", commandConfig.name))
	}
	commands[commandConfig.name] = commandConfig
}

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run(args []string, writer io.Writer) error {
	if len(args) == 0 || args[0] == "--help" || args[0] == "help" {
		printUsage(writer)
		return nil
	}

	selectedCommand, exists := commands[args[0]]
	if !exists {
		return fmt.Errorf("unknown command %q", args[0])
	}

	return selectedCommand.run(args[1:])
}

func printUsage(writer io.Writer) {
	fmt.Fprintln(writer, "canonical-calc runs small calculator workflows")
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "Usage:")
	fmt.Fprintln(writer, "  canonical-calc <command> [args]")
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "Commands:")

	names := make([]string, 0, len(commands))
	for name := range commands {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		commandConfig := commands[name]
		fmt.Fprintf(writer, "  %-16s %s\n", commandConfig.usage, commandConfig.summary)
	}
}
