package cmd

import (
	"bufio"
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
	scanner := bufio.NewScanner(os.Stdin)
	rootCmd := cs.Unify.GetCmd()

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		cs.ResetPipelineState()

		err = cs.Data.DateParse(".", line)
		if err != nil {
			break
		}

		if len(args) == 0 {
			cs.Data.Print()
			continue
		}

		rootCmd.SetArgs(args)
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
	cs.Error = nil
	cs.Data = cal.Data{
		Convert:    convert,
		FormatType: formatType,
		GoFormat:   goFormat,
		CppFormat:  cppFormat,
		JavaFormat: javaFormat,
	}
}
