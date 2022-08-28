package cmdCron

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdLog"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
)

func Exec(command string, args ...string) error {
	var err error

	for range Only.Once {
		cmdLog.Printf("Exec START: %s %v\n", command, args)

		cmd := exec.Command(command, args...)
		// out, err := cmd.CombinedOutput()
		// if err != nil {
		// 	break
		// }
		// LogPrintf("\n%s\n", string(out))

		var stdout io.ReadCloser
		stdout, err = cmd.StdoutPipe()
		if err != nil {
			break
		}

		// var stderr io.ReadCloser
		// stderr, err = cmd.StderrPipe()
		// if err != nil {
		// 	break
		// }

		// start the command after having set up the pipe
		err = cmd.Start()
		if err != nil {
			break
		}

		// read command's stdout line by line
		in := bufio.NewScanner(stdout)
		// inerr := bufio.NewScanner(stderr)

		// go func(){
		// 	for in.Scan() {
		// 		LogPrintf(inerr.Text()) // write each line to your log, or anything you need
		// 	}
		// }()

		for in.Scan() {
			cmdLog.Printf(in.Text()) // write each line to your log, or anything you need
		}

		err = in.Err()
		if err != nil {
			cmdLog.Printf("error: %s", err)
		}

		cmdLog.Printf("Exec STOP: %s %v\n", command, args)
	}

	return err
}

func Exec1(command string, args ...string) error {
	var err error

	for range Only.Once {
		cmdLog.Printf("Exec1 START: %s %v\n", command, args)

		cmd := exec.Command(command, args...)

		var stdoutBuf, stderrBuf bytes.Buffer
		cmd.Stdout = io.MultiWriter(&stdoutBuf) // os.Stdout, &stdoutBuf)
		cmd.Stderr = io.MultiWriter(&stderrBuf) // os.Stderr, &stderrBuf)

		err = cmd.Run()
		if err != nil {
			cmdLog.Printf("cmd.Run() failed with %s\n", err)
		}
		outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
		fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)

		// out, err := cmd.CombinedOutput()
		// if err != nil {
		// 	break
		// }
		// LogPrintf("\n%s\n", string(out))

		cmdLog.Printf("Exec STOP: %s %v\n", command, args)
	}

	return err
}

func Exec2(command string, args ...string) error {
	var err error

	for range Only.Once {
		cmdLog.Printf("Exec2 START: %s %v\n", command, args)

		cmd := exec.Command(command, args...)

		var stdout, stderr []byte
		var errStdout, errStderr error
		stdoutIn, _ := cmd.StdoutPipe()
		stderrIn, _ := cmd.StderrPipe()
		err = cmd.Start()
		if err != nil {
			cmdLog.Printf("cmd.Start() failed with '%s'\n", err)
			break
		}

		// cmd.Wait() should be called only after we finish reading
		// from stdoutIn and stderrIn.
		// wg ensures that we finish
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			stdout, errStdout = copyAndCapture(os.Stdout, stdoutIn)
			wg.Done()
		}()
		stderr, errStderr = copyAndCapture(os.Stderr, stderrIn)
		wg.Wait()

		err = cmd.Wait()
		if err != nil {
			cmdLog.Printf("cmd.Run() failed with %s\n", err)
			break
		}
		if errStdout != nil || errStderr != nil {
			cmdLog.Printf("failed to capture stdout or stderr\n")
			break
		}
		outStr, errStr := string(stdout), string(stderr)
		fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)

		cmdLog.Printf("Exec STOP: %s %v\n", command, args)
	}

	return err
}

func copyAndCapture(w io.Writer, r io.Reader) ([]byte, error) {
	var out []byte
	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			_, err := w.Write(d)
			if err != nil {
				return out, err
			}
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}
}

func Exec3(command string, args ...string) error {
	var err error

	for range Only.Once {
		cmdLog.Printf("Exec3 START: %s %v\n", command, args)

		cmd := exec.Command(command, args...)

		var errStdout, errStderr error
		stdoutIn, _ := cmd.StdoutPipe()
		stderrIn, _ := cmd.StderrPipe()
		stdout := NewCapturingPassThroughWriter(os.Stdout)
		stderr := NewCapturingPassThroughWriter(os.Stderr)
		err = cmd.Start()
		if err != nil {
			cmdLog.Printf("cmd.Start() failed with '%s'\n", err)
			break
		}

		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			_, errStdout = io.Copy(stdout, stdoutIn)
			wg.Done()
		}()

		_, errStderr = io.Copy(stderr, stderrIn)
		wg.Wait()

		err = cmd.Wait()
		if err != nil {
			cmdLog.Printf("cmd.Run() failed with %s\n", err)
			break
		}
		if errStdout != nil || errStderr != nil {
			cmdLog.Printf("failed to capture stdout or stderr\n")
			break
		}

		outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
		fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)

		cmdLog.Printf("Exec STOP: %s %v\n", command, args)
	}

	return err
}

// CapturingPassThroughWriter is a writer that remembers
// data written to it and passes it to w
type CapturingPassThroughWriter struct {
	buf bytes.Buffer
	w   io.Writer
}

// NewCapturingPassThroughWriter creates new CapturingPassThroughWriter
func NewCapturingPassThroughWriter(w io.Writer) *CapturingPassThroughWriter {
	return &CapturingPassThroughWriter{
		w: w,
	}
}

func (w *CapturingPassThroughWriter) Write(d []byte) (int, error) {
	w.buf.Write(d)
	return w.w.Write(d)
}

// Bytes returns bytes written to the writer
func (w *CapturingPassThroughWriter) Bytes() []byte {
	return w.buf.Bytes()
}
