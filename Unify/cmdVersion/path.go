package cmdVersion

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdLog"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
)

func (v *Version) SetCmd(a ...string) error {
	var err error

	for range Only.Once {
		v.Cmd, err = filepath.Abs(filepath.Join(a...))
		if err != nil {
			break
		}

		v.CmdDir = filepath.Dir(v.Cmd)
		v.CmdFile = filepath.Base(v.Cmd)
	}

	return err
}

func (v *Version) IsBootstrapBinary() bool {
	var ok bool
	for range Only.Once {
		if v.ExecName != v.CmdFile {
			break
		}
		if v.ExecName != BootstrapBinaryName {
			break
		}
		ok = true
	}
	return ok
}

func (v *Version) AutoRun() State {
	for range Only.Once {
		if !v.AutoExec {
			break
		}

		if v.IsBootstrapBinary() {
			// Let's avoid an endless loop.
			break
		}

		if len(v.FullArgs) > 0 {
			if v.FullArgs[0] == CmdVersion {
				// Let's avoid another endless loop.
				break
			}
		}

		// @TODO - This is broken!
		cmd := exec.Command(v.TargetBinary, []string{"version", "info"}...)

		var stdout io.ReadCloser
		var err error
		stdout, err = cmd.StdoutPipe()
		if err != nil {
			break
		}

		err = cmd.Start()
		if err != nil {
			break
		}

		in := bufio.NewScanner(stdout)

		for in.Scan() {
			cmdLog.Printf(in.Text()) // write each line to your log, or anything you need
		}

		err = in.Err()
		if err != nil {
			cmdLog.Printf("error: %s", err)
			v.State.SetError(err.Error())
			break
		}

		// @TODO - This is broken!
		// fmt.Printf("Executing the real binary: '%s'\n", v.RuntimeBinary)
		// // c := exec.Command(v.TargetBinary, v.FullArgs...)
		// c := exec.Command(v.TargetBinary, []string{"version"}...)
		//
		// var stdoutBuf, stderrBuf bytes.Buffer
		// c.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
		// c.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)
		// err := c.Run()
		// waitStatus := c.ProcessState.Sys().(syscall.WaitStatus)
		// waitStatus.ExitStatus()
		//
		// if err != nil {
		// 	fmt.Printf("stdoutBuf: %s\n", stdoutBuf.String())
		// 	fmt.Printf("stderrBuf: %s\n", stderrBuf.String())
		// 	v.State.SetError(err.Error())
		// 	break
		// }

		v.State.SetOk("")
	}

	return v.State
}

func (v *Version) CreateDummyBinary() State {
	for range Only.Once {
		var err error

		result := FileStat(v.RuntimeBinary, v.TargetBinary)
		if result.CopyOfRuntime {
			v.AutoExec = true
			break
		}

		//if result.IsRuntimeBinary {
		//	// We are running as the bootstrap binary.
		//	su.State.SetOk()
		//	break
		//}

		if result.LinkToRuntime {
			err = os.Remove(v.TargetBinary)
			if err != nil {
				v.State.SetError(err.Error())
				break
			}
			result.IsMissing = true
			v.AutoExec = true
		}

		if result.IsMissing {
			err = CopyFile(v.RuntimeBinary, v.TargetBinary)
			if err != nil {
				v.State.SetError(err.Error())
				break
			}
			v.AutoExec = true
		}
	}

	return v.State
}

func (v *Version) IsRunningAs(run string) bool {
	var ok bool
	// If OK - running executable file matches the string 'run'.
	//ok, err := regexp.MatchString("^" + run, r.CmdFile)

	if v.IsWindows() {
		//fmt.Printf("DEBUG: WINDOWS!\n")
		ok = strings.HasPrefix(run, strings.TrimSuffix(v.CmdFile, ".exe"))
		//run = strings.TrimSuffix(run, ".exe")
	} else {
		ok = strings.HasPrefix(run, v.CmdFile)
	}
	//fmt.Printf("DEBUG: Cmd.Runtime.IsRunningAs?? %s\n", ok)
	//fmt.Printf("DEBUG: run: %s\n", run)
	//fmt.Printf("DEBUG: r.CmdName: %s\n", r.CmdName)
	//fmt.Printf("DEBUG: r.CmdFile: %s\n", r.CmdFile)
	return ok
}

func (v *Version) IsRunningAsFile() bool {
	// If OK - running executable file matches the application binary name.
	//ok, err := regexp.MatchString("^" + r.CmdName, r.CmdFile)
	ok := strings.HasPrefix(v.ExecName, v.CmdFile)
	return ok
}

func (v *Version) IsRunningAsLink() bool {
	return !v.IsRunningAsFile()
}

type Path string

func NewPath(path ...string) *Path {
	var ret Path
	ret.Set(path...)
	return &ret
}

func (p *Path) DirExists() bool {
	var ok bool

	for range Only.Once {
		stat, err := os.Stat(string(*p))
		if os.IsNotExist(err) {
			break
		}

		if !stat.IsDir() {
			break
		}

		ok = true
	}

	return ok
}

func (p *Path) FileExists() bool {
	var ok bool

	for range Only.Once {
		stat, err := os.Stat(string(*p))
		if os.IsNotExist(err) {
			break
		}

		if stat.IsDir() {
			break
		}

		ok = true
	}

	return ok
}

func (p *Path) Chmod(mode os.FileMode) bool {
	var ok bool

	for range Only.Once {
		err := os.Chmod(string(*p), mode)
		if err != nil {
			break
		}

		ok = true
	}

	return ok
}

func (p *Path) Set(path ...string) {

	for range Only.Once {
		dir := filepath.Join(path...)
		if strings.HasPrefix(dir, "~/") {
			u, err := user.Current()
			if err != nil {
				break
			}
			dir = strings.TrimPrefix(dir, "~/")
			dir = filepath.Join(u.HomeDir, dir)
		}

		*p = Path(dir)
	}
}

func (p *Path) String() string {
	return (string)(*p)
}

//func (p *Path) Set(elem ...string) Path {
//	return (Path)(filepath.Join(elem...))
//}

func (p *Path) Join(elem ...string) Path {
	var pa []string
	//if p == nil {
	//	*p = "/"
	//}
	pa = append(pa, (string)(*p))
	pa = append(pa, elem...)
	return (Path)(filepath.Join(pa...))
}

func (p *Path) MkdirAll() error {
	var err error

	for range Only.Once {
		if p.DirExists() {
			break
		}

		dir := string(*p)
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			break
		}

		if !p.DirExists() {
			err = errors.New("no dir")
			break
		}
	}

	return err
}

func (p *Path) Copy(fp string) error {
	var err error

	for range Only.Once {
		var stat os.FileInfo
		stat, err = os.Stat(fp)
		if os.IsNotExist(err) {
			break
		}
		if stat.IsDir() {
			err = errors.New("file is a dir")
			break
		}

		var input []byte
		input, err = ioutil.ReadFile(fp)
		if err != nil {
			break
		}

		dfp := filepath.Join(string(*p), filepath.Base(fp))
		err = ioutil.WriteFile(dfp, input, stat.Mode())
		if err != nil {
			break
		}
	}

	return err
}

func (p *Path) Move(fp string) error {
	var err error

	for range Only.Once {
		err = p.Copy(fp)
		if err != nil {
			break
		}

		err = os.Remove(fp)
		if err != nil {
			break
		}
	}

	return err
}

func (p *Path) GrepFile(search string) (int, error) {
	var line int
	var err error

	for range Only.Once {
		p.Set(string(*p))

		var f *os.File
		f, err = os.Open(string(*p))
		if err != nil {
			// Silently ignore missing files.
			err = nil
			break
		}
		//goland:noinspection GoDeferInLoop,GoUnhandledErrorResult
		defer f.Close()

		// Splits on newlines by default.
		scanner := bufio.NewScanner(f)
		line = 1
		// https://golang.org/pkg/bufio/#Scanner.Scan
		for scanner.Scan() {
			if strings.Contains(scanner.Text(), search) {
				break
			}

			line++
		}

		err = scanner.Err()
		if err != nil {
			break
		}
	}

	return line, err
}


// FileRead Retrieves data from a local file.
func (p *Path) FileRead(ref interface{}) error {
	var err error
	for range Only.Once {
		if *p == "" {
			err = errors.New("empty file")
			break
		}

		var f *os.File
		f, err = os.Open(string(*p))
		if err != nil {
			if os.IsNotExist(err) {
				err = nil
			}
			break
		}

		//goland:noinspection GoUnhandledErrorResult,GoDeferInLoop
		defer f.Close()

		err = json.NewDecoder(f).Decode(&ref)
	}

	// for range Only.Once {
	//	fn := ep.GetFilename()
	//	if err != nil {
	//		break
	//	}
	//
	//	ret, err = os.FileRead(fn)
	//	if err != nil {
	//		break
	//	}
	// }

	return err
}

// FileWrite Saves data to a file path.
func (p *Path) FileWrite(ref interface{}, perm os.FileMode) error {
	var err error
	for range Only.Once {
		if *p == "" {
			err = errors.New("empty file")
			break
		}

		var f *os.File
		f, err = os.OpenFile(string(*p), os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm)
		if err != nil {
			err = errors.New(fmt.Sprintf("Unable to write to file %s - %v", string(*p), err))
			break
		}

		//goland:noinspection GoUnhandledErrorResult,GoDeferInLoop
		defer f.Close()
		err = json.NewEncoder(f).Encode(ref)

		// fn := ep.GetFilename()
		// if err != nil {
		//	break
		// }
		//
		// err = os.FileWrite(fn, data, perm)
		// if err != nil {
		//	break
		// }
	}

	return err
}

// PlainFileRead Retrieves data from a local file.
func (p *Path) PlainFileRead() ([]byte, error) {
	var data []byte
	var err error
	for range Only.Once {
		if *p == "" {
			err = errors.New("empty file")
			break
		}

		var f *os.File
		f, err = os.Open(string(*p))
		if err != nil {
			if os.IsNotExist(err) {
				err = nil
			}
			break
		}

		//goland:noinspection GoUnhandledErrorResult,GoDeferInLoop
		defer f.Close()

		data, err = ioutil.ReadAll(f)
	}

	return data, err
}

// PlainFileWrite Saves data to a file path.
func (p *Path) PlainFileWrite(data []byte, perm os.FileMode) error {
	var err error
	for range Only.Once {
		if *p == "" {
			err = errors.New("empty file")
			break
		}

		var f *os.File
		f, err = os.OpenFile(string(*p), os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm)
		if err != nil {
			err = errors.New(fmt.Sprintf("Unable to write to file %s - %v", string(*p), err))
			break
		}
		//goland:noinspection GoUnhandledErrorResult,GoDeferInLoop
		defer f.Close()

		_, err = f.Write(data)
	}

	return err
}

// FileRemove Removes a file path.
func (p *Path) FileRemove() error {
	var err error
	for range Only.Once {
		if *p == "" {
			err = errors.New("empty file")
			break
		}

		var f os.FileInfo
		f, err = os.Stat(string(*p))
		if os.IsNotExist(err) {
			err = nil
			break
		}
		if err != nil {
			break
		}
		if f.IsDir() {
			err = errors.New("file is a directory")
			break
		}

		err = os.Remove(string(*p))
	}

	return err
}


//goland:noinspection SpellCheckingInspection
var RcFiles = []Path{
	// BASH
	"/etc/profile",
	"/etc/bashrc",
	"~/.profile",
	"~/.bash_profile",
	"~/.bashrc",
	"~/.bash_login",
	"~/.bash_logout",

	// ZSH
	"/etc/zlogin",
	"/etc/zlogout",
	"/etc/zprofile",
	"/etc/zshenv",
	"/etc/zshrc",
	"~/.zlogin",
	"~/.zlogout",
	"~/.zprofile",
	"~/.zshenv",
	"~/.zshrc",

	// CSH
	"/etc/csh.cshrc",
	"/etc/csh.login",
	"/etc/csh.logout",
	"~/.cshrc",
	"~/.login",
	"~/.logout",
}

//goland:noinspection GoUnusedExportedFunction
func GrepFiles(search string, fps ...Path) ([]string, error) {
	var files []string
	var err error

	if fps == nil {
		fps = RcFiles
	}
	if len(fps) == 0 {
		fps = RcFiles
	}

	for _, p := range fps {
		var line int
		line, err = p.GrepFile(search)
		if line > 0 {
			files = append(files, p.String()+" line:"+strconv.Itoa(line))
		}
	}

	return files, err
}

type TargetFile struct {
	IsMissing       bool
	IsRuntimeBinary bool
	FileMatches     bool
	IsSymlink       bool
	LinkTo          string
	LinkEval        string
	LinkToRuntime   bool
	CopyOfRuntime   bool

	Error error
	Info  os.FileInfo
}

func FileStat(runtimeBinary string, targetBinary string) *TargetFile {
	var targetFile TargetFile

	for range Only.Once {
		targetFile.Info, targetFile.Error = os.Stat(targetBinary)
		if os.IsNotExist(targetFile.Error) {
			targetFile.IsMissing = true
		} else {
			targetFile.IsMissing = false

			if filepath.Base(runtimeBinary) == BootstrapBinaryName {
				targetFile.IsRuntimeBinary = true
			} else if runtimeBinary == targetBinary {
				targetFile.IsRuntimeBinary = true
				targetFile.CopyOfRuntime = true
			} else {
				targetFile.IsRuntimeBinary = false

				targetFile.Error = CompareBinary(runtimeBinary, targetBinary)
				if targetFile.Error == nil {
					targetFile.FileMatches = true
				} else {
					targetFile.FileMatches = false
				}
			}
		}

		targetFile.LinkTo, targetFile.Error = os.Readlink(targetBinary)
		if targetFile.LinkTo != "" {
			targetFile.IsSymlink = true

			targetFile.LinkEval, targetFile.Error = filepath.EvalSymlinks(targetBinary)
			if targetFile.LinkEval == "" {
				targetFile.LinkToRuntime = false
			} else {
				targetFile.LinkEval, targetFile.Error = filepath.Abs(targetFile.LinkEval)
				if targetFile.LinkEval == runtimeBinary {
					targetFile.LinkToRuntime = true
				} else if filepath.Base(targetFile.LinkEval) == BootstrapBinaryName {
					targetFile.LinkToRuntime = true
				} else {
					targetFile.LinkToRuntime = false
				}
			}
		} else {
			targetFile.IsSymlink = false
		}
	}

	return &targetFile
}

func ResolveFile(file string) string {
	var result string
	var err error

	for range Only.Once {
		_, err = os.Stat(file)
		if os.IsNotExist(err) {
			break
		}

		result, err = os.Readlink(file)
		if result == "" {
			result = file
			break
		}

		result, err = filepath.EvalSymlinks(file)
		if result == "" {
			result = file
			break
		}

		result, err = filepath.Abs(result)
		if result == "" {
			result = file
			break
		}
	}

	return result
}
