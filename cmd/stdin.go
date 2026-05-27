package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"GoWhen/cmd/cal"
)

func (cs *Cmds) HasPipedStdin() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}

	return stat.Mode()&os.ModeCharDevice == 0
}

func (cs *Cmds) ExecutePipedStdin() error {
	var err error

	args := os.Args[1:]
	if len(args) < 3 || args[0] != "parse" {
		return errors.New("parse must be the first command when reading piped stdin; use parse <format> - to read stdin")
	}

	rootCmd := cs.Unify.GetCmd()
	if args[2] != "-" {
		fmt.Fprintln(os.Stderr, "warning: piped stdin ignored because parse date argument is not '-' ")
		cs.ResetPipelineState()
		rootCmd.SetArgs(args)
		return rootCmd.Execute()
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		cs.ResetPipelineState()
		lineArgs := append([]string{}, args...)
		lineArgs[2] = line
		rootCmd.SetArgs(lineArgs)
		err = rootCmd.Execute()
		if err != nil {
			break
		}
	}

	if err != nil {
		return err
	}

	return scanner.Err()
}

func (cs *Cmds) ResetPipelineState() {
	convert := cs.Data.Convert
	formatType := cs.Data.FormatType
	goFormat := cs.Data.GoFormat
	cppFormat := cs.Data.CppFormat
	javaFormat := cs.Data.JavaFormat

	cs.reparse = false
	cs.last = false
	cs.parseRan = false
	cs.formatNoHeaders = false
	cs.Error = nil
	cs.Data = cal.Data{
		Convert:    convert,
		FormatType: formatType,
		GoFormat:   goFormat,
		CppFormat:  cppFormat,
		JavaFormat: javaFormat,
	}
	cs.Data.ClearSelectors()
}
