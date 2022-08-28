package cmdVersion

import (
	"GoWhen/Unify/Only"
	"fmt"
	"github.com/kardianos/osext"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"github.com/spf13/cobra"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// type SelfUpdateGetter interface {
// }

type SelfUpdateArgs struct {
	owner      *string
	name       *string
	version    *string
	sourceRepo *string
	binaryRepo *string

	logging *bool
}

//type MyValidator struct {
//}
//func (v *MyValidator) Validate(release, asset []byte) error {
//	calculatedHash := fmt.Sprintf("%x", sha256.Sum256(release))
//	hash := fmt.Sprintf("%s", asset[:sha256.BlockSize])
//	if calculatedHash != hash {
//		return fmt.Errorf("sha2: validation failed: hash mismatch: expected=%q, got=%q", calculatedHash, hash)
//	}
//	return nil
//}
//func (v *MyValidator) Suffix() string {
//	return ".gz"
//}

//goland:noinspection ALL
type Version struct {
	ExecName       string   `json:"cmd_name" mapstructure:"cmd_name"`
	ExecVersion    string   `json:"cmd_version" mapstructure:"cmd_version"`
	ExecSourceRepo UrlValue `json:"cmd_source_repo" mapstructure:"cmd_source_repo"`
	ExecBinaryRepo UrlValue `json:"cmd_binary_repo" mapstructure:"cmd_binary_repo"`

	Cmd     string `json:"cmd" mapstructure:"cmd"`
	CmdDir  string `json:"cmd_dir" mapstructure:"cmd_dir"`
	CmdFile string `json:"cmd_file" mapstructure:"cmd_file"`

	WorkingDir Path `json:"working_dir" mapstructure:"working_dir"`
	BaseDir    Path `json:"base_dir" mapstructure:"base_dir"`
	BinDir     Path `json:"bin_dir" mapstructure:"bin_dir"`
	ConfigDir  Path `json:"config_dir" mapstructure:"config_dir"`
	CacheDir   Path `json:"cache_dir" mapstructure:"cache_dir"`
	TempDir    Path `json:"temp_dir" mapstructure:"temp_dir"`

	FullArgs ExecArgs `json:"full_args" mapstructure:"full_args"`
	Args     ExecArgs `json:"args" mapstructure:"args"`
	ArgFiles ExecArgs `json:"arg_files" mapstructure:"arg_files"`

	Env    ExecEnv     `json:"env" mapstructure:"env"`
	EnvMap Environment `json:"env_map" mapstructure:"env_map"`

	TimeStamp time.Time     `json:"timestamp" mapstructure:"timestamp"`
	Timeout   time.Duration `json:"timeout" mapstructure:"timeout"`

	GoRuntime GoRuntime `json:"go_runtime" mapstructure:"go_runtime"`

	User User `json:"user" mapstructure:"user"`

	Debug   bool  `json:"debug" mapstructure:"debug"`
	Verbose bool  `json:"verbose" mapstructure:"verbose"`
	State   State `json:"state" mapstructure:"state"`

	useRepo       *UrlValue
	OldVersion    *VersionValue
	TargetBinary  string
	RuntimeBinary string
	AutoExec      bool

	logging *FlagValue
	config  *selfupdate.Config
	ref     *selfupdate.Updater

	cmd     *cobra.Command
	SelfCmd *cobra.Command
}

func (v *Version) GetCmd() *cobra.Command {
	return v.SelfCmd
}

func (v *Version) IsValid() bool {
	var ok bool
	for range Only.Once {
		if v.useRepo.Owner == "" {
			v.State.SetWarning("rep owner is not defined - selfupdate disabled")
			break
		}

		if v.useRepo.Name == "" {
			v.State.SetWarning("repo name is not defined - selfupdate disabled")
			break
		}

		// Refer to binary repo definition first.
		if v.ExecBinaryRepo.IsValid() {
			v.useRepo = &v.ExecBinaryRepo
			v.State.SetOk("")
			ok = true
			break
		}

		// If binary repo is not set, use source repo.
		if v.ExecSourceRepo.IsValid() {
			v.useRepo = &v.ExecSourceRepo
			v.State.SetOk("")
			ok = true
			break
		}

		v.State.SetWarning(errorNoRepo)
	}

	return ok
}
func (v *Version) IsNotValid() bool {
	return !v.IsValid()
}

func (v *Version) getRepo() string {
	var ret string

	for range Only.Once {
		if v.ExecBinaryRepo.IsValid() {
			ret = v.ExecBinaryRepo.String()
			break
		}
		if v.ExecSourceRepo.IsValid() {
			ret = v.ExecSourceRepo.String()
			break
		}
	}

	return ret
}

func (v *Version) GetSemVer() *VersionValue {
	// v := semver.MustParse(r.CmdVersion)
	// return semver.Version(v.String())
	return GetSemVer(v.ExecVersion)
}

func (v *Version) PrintNameVersion() {
	fmt.Printf("%s ", v.ExecName)
	fmt.Printf("v%s", v.ExecVersion)
}

func (v *Version) TimeStampString() string {
	return v.TimeStamp.Format("2006-01-02T15:04:05-0700")
}

func (v *Version) TimeStampEpoch() int64 {
	return v.TimeStamp.Unix()
}

func (v *Version) GetTimeout() string {
	//d := r.Timeout.Round(time.Second)
	d := v.Timeout
	h := d / time.Hour
	d -= h * time.Hour

	m := d / time.Minute

	s := m / time.Second

	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

func (v *Version) GetEnvMap() *Environment {
	return &v.EnvMap
}

func (v *Version) SetRepos(source string, binary string) State {
	// if state := ux.IfNilReturnError(r); state.IsError() {
	// 	return state
	// }
	v.ExecSourceRepo = toUrlValue(source)
	v.ExecBinaryRepo = toUrlValue(binary)

	return v.State
}

func (v *Version) EnsureNotNil() *Version {
	if v == nil {
		return New("binary", "version", false)
	}
	return v
}

func (v *Version) IsWindows() bool {
	var ok bool
	if v.GoRuntime.Os == "windows" {
		ok = true
	}
	return ok
}

func (v *Version) IsMac() bool {
	var ok bool
	if v.GoRuntime.Os == "darwin" {
		ok = true
	}
	return ok
}

func (v *Version) IsOsx() bool {
	var ok bool
	if v.GoRuntime.Os == "darwin" {
		ok = true
	}
	return ok
}

type ExecEnv []string
type Environment map[string]string
type GoRuntime struct {
	Os       string
	Arch     string
	Root     string
	Version  string
	Compiler string
	NumCpus  int
}

type User struct {
	*user.User
}

// Instead of creating every time, let's cache the initial result in a global variable.
var globalRuntime *Version

func New(binary string, version string, debugFlag bool) *Version {
	var ret *Version

	for range Only.Once {
		if globalRuntime != nil {
			// Instead of creating every time, let's cache the initial result in a global variable.
			//globalRuntime.TimeStamp = time.Now()
			ret = globalRuntime
			break
		}

		ret = &Version{
			ExecName:    binary,
			ExecVersion: version,

			Cmd:     "",
			CmdDir:  "",
			CmdFile: "",

			WorkingDir: ".",
			BaseDir:    ".",
			BinDir:     ".",
			ConfigDir:  ".",
			CacheDir:   ".",
			TempDir:    ".",

			FullArgs: os.Args,
			Args:     os.Args[1:],
			ArgFiles: []string{},

			Env:    os.Environ(),
			EnvMap: make(Environment),

			TimeStamp: time.Now(),

			GoRuntime: GoRuntime{
				Os:       runtime.GOOS,
				Arch:     runtime.GOARCH,
				Root:     runtime.GOROOT(),
				Version:  runtime.Version(),
				Compiler: runtime.Compiler,
				NumCpus:  runtime.NumCPU(),
			},

			Debug:   debugFlag,
			Verbose: false,
			State:   State{},
		}

		for _, item := range os.Environ() {
			s := strings.SplitN(item, "=", 2)
			ret.EnvMap[s[0]] = s[1]
		}

		var err error
		var exe string
		var p string
		//ret.Cmd, err = os.Executable()
		//if err != nil {
		//	ret.State.SetError(err)
		//	break
		//}
		//ret.Cmd, err = filepath.Abs(ret.Cmd)
		//if err != nil {
		//	ret.State.SetError(err)
		//	break
		//}
		exe, err = osext.Executable()
		if err != nil {
			ret.State.SetError(err.Error())
			break
		}
		//if ret.GoRuntime.Os == "windows" {
		//	exe = strings.TrimSuffix(exe,".exe")
		//}
		ret.Cmd = exe
		ret.CmdDir = filepath.Dir(exe)
		ret.CmdFile = filepath.Base(exe)

		ret.User.User, err = user.Current()
		if err != nil {
			ret.State.SetError(err.Error())
			break
		}

		p, err = os.Getwd()
		if err != nil {
			ret.State.SetError(err.Error())
			break
		}
		ret.WorkingDir.Set(p)

		//if runtime.GOOS == "windows" {
		//	ret.BaseDir = ""
		//} else {
		ret.BaseDir.Set(ret.User.HomeDir, "."+ret.ExecName)
		//}

		//if runtime.GOOS == "windows" {
		//	ret.BinDir = ""
		//} else {
		ret.BinDir = ret.BaseDir.Join("bin")
		//}

		p, err = os.UserConfigDir()
		if err != nil {
			if runtime.GOOS == "windows" {
				ret.ConfigDir = ""
			} else {
				ret.ConfigDir = "."
			}
		} else {
			ret.ConfigDir = ret.BaseDir.Join("etc")
		}
		//ret.ConfigDir = filepath.Join(ret.ConfigDir, ret.CmdName)

		p, err = os.UserCacheDir()
		if err != nil {
			if runtime.GOOS == "windows" {
				ret.CacheDir = ""
			} else {
				ret.CacheDir = "."
			}
		} else {
			ret.CacheDir = ret.BaseDir.Join("cache")
		}
		//ret.CacheDir = filepath.Join(ret.CacheDir, ret.CmdName)

		p = os.TempDir()
		if ret.TempDir == "" {
			if runtime.GOOS == "windows" {
				ret.TempDir = "C:\\tmp"
			} else {
				ret.TempDir = "/tmp"
			}
		} else {
			ret.TempDir = ret.BaseDir.Join("tmp")
		}

		// ******************************************************************************** //
		ret.TargetBinary = ret.Cmd
		ret.RuntimeBinary = ResolveFile(ret.Cmd)
		ret.AutoExec = false

		ret.logging = toBoolValue(ret.Debug)
		ret.config = &selfupdate.Config{
			APIToken:            "",
			EnterpriseBaseURL:   "",
			EnterpriseUploadURL: "",
			Validator:           nil, // &MyValidator{},
			Filters:             []string{},
		}

		ret.useRepo = &ret.ExecBinaryRepo

		// Workaround for selfupdate not being flexible enough to support variable asset names
		// Should enable a template similar to GoReleaser.
		// EG: {{ .ProjectName }}-{{ .Os }}_{{ .Arch }}
		//var asset string
		//asset, te.State = toolGhr.GetAsset(rt.CmdBinaryRepo, "latest")
		//te.config.Filters = append(te.config.Filters, asset)

		// Ignore the above and just make sure all filenames are lowercase.
		ret.config.Filters = append(ret.config.Filters, addFilters(ret.CmdFile, runtime.GOOS, runtime.GOARCH)...)
		ret.ref, _ = selfupdate.NewUpdater(*ret.config)
		if *ret.logging {
			selfupdate.EnableLog()
		}
		// ******************************************************************************** //

		// Instead of creating every time, let's cache the initial result in a global variable.
		globalRuntime = ret
	}

	return ret
}
