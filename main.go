package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/xyproto/ask"
	"github.com/xyproto/files"
	"github.com/xyproto/vt"
)

func compile(sources, compilerPath string) (bool, error) {
	err := os.WriteFile("/tmp/source", []byte(sources), 0o644)
	if err != nil {
		return false, err
	}
	if files.Run(compilerPath+" /tmp/source -o /tmp/exe") == nil {
		return true, nil
	}
	return false, errors.New("could not compile")
}

func main() {
	compiler := "gcc"
	args := os.Args
	if len(args) > 1 {
		compiler = strings.TrimSpace(args[1])
	}

	compilerPath := files.Which(compiler)
	if compilerPath == "" {
		fmt.Fprintf(os.Stderr, "error: could not find "+compiler+" in the PATH")
		os.Exit(1)
	}

	o := vt.New()
	o.Printf("<blue>REPL</blue> <white>for</white> <green>%s</green>\n", compiler)
	o.Printf("<white>%s</white>\n", strings.Repeat("-", 80))

	var sources string

OUT:
	for {
		input := ask.Ask("> ")

		switch strings.TrimSpace(input) {
		case "q", "quit", "exit":
			break OUT
		}

		fmt.Println("current sources:")
		sources += input + "\n"

		if ok, err := compile(sources, compilerPath); ok && err == nil {
			fmt.Println("success, compiled")
		} else if err != nil {
			fmt.Println("error: " + err.Error())
			fmt.Println(sources)
		} else {
			fmt.Println("does not compile yet:")
			fmt.Println(sources)
		}
	}
}
