package cmdVersion

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdHelp"
	"context"
	"fmt"
	"github.com/google/go-github/v30/github"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/tcnksm/go-gitconfig"
	"os"
)

const Group = "Version"

func (v *Version) AttachCommands(cmd *cobra.Command, disableVflag bool) State {
	for range Only.Once {
		if cmd == nil {
			break
		}
		v.cmd = cmd

		v.SelfCmd = &cobra.Command{
			Use:                   CmdVersion,
			Short:                 fmt.Sprintf("Self-manage this executable."),
			Long:                  fmt.Sprintf("Self-manage this executable."),
			DisableFlagParsing:    true,
			DisableFlagsInUseLine: true,
			RunE: func(cmd *cobra.Command, args []string) error {
				v.OldVersion = toVersionValue(v.ExecVersion)
				v.State = v.CmdVersion(cmd, args...)
				return v.State.GetState()
			},
		}
		v.cmd.AddCommand(v.SelfCmd)
		v.SelfCmd.Example = cmdHelp.PrintExamples(v.SelfCmd, "")
		v.SelfCmd.Annotations = map[string]string{"group": Group}

		var selfUpdateCmd = &cobra.Command{
			Use:                   CmdSelfUpdate,
			Short:                 fmt.Sprintf("Update version of executable."),
			Long:                  fmt.Sprintf("Check and update the latest version."),
			DisableFlagParsing:    true,
			DisableFlagsInUseLine: true,
			RunE: func(cmd *cobra.Command, args []string) error {
				v.OldVersion = toVersionValue(v.ExecVersion)
				v.State = v.CmdVersionUpdate()
				return v.State.GetState()
			},
		}
		v.cmd.AddCommand(selfUpdateCmd)
		selfUpdateCmd.Example = cmdHelp.PrintExamples(v.SelfCmd, "")
		selfUpdateCmd.Annotations = map[string]string{"group": Group}

		var versionCheckCmd = &cobra.Command{
			Use:                   CmdVersionCheck,
			Short:                 fmt.Sprintf("Check and show any version updates."),
			Long:                  fmt.Sprintf("Check and show any version updates."),
			DisableFlagParsing:    true,
			DisableFlagsInUseLine: true,
			RunE: func(cmd *cobra.Command, args []string) error {
				v.OldVersion = toVersionValue(v.ExecVersion)
				v.State = v.CmdVersionCheck()
				return v.State.GetState()
			},
		}
		v.SelfCmd.AddCommand(versionCheckCmd)
		versionCheckCmd.Example = cmdHelp.PrintExamples(v.SelfCmd, "")
		versionCheckCmd.Annotations = map[string]string{"group": Group}

		var versionListCmd = &cobra.Command{
			Use:                   CmdVersionList,
			Short:                 fmt.Sprintf("List available versions."),
			Long:                  fmt.Sprintf("List available versions."),
			DisableFlagParsing:    true,
			DisableFlagsInUseLine: true,
			RunE: func(cmd *cobra.Command, args []string) error {
				v.OldVersion = toVersionValue(v.ExecVersion)
				v.State = v.CmdVersionList(args...)
				return v.State.GetState()
			},
		}
		v.SelfCmd.AddCommand(versionListCmd)
		versionListCmd.Example = cmdHelp.PrintExamples(v.SelfCmd, "")
		versionListCmd.Annotations = map[string]string{"group": Group}

		var versionInfoCmd = &cobra.Command{
			Use:                   CmdVersionInfo,
			Short:                 fmt.Sprintf("Info on current version."),
			Long:                  fmt.Sprintf("Info on current version."),
			DisableFlagParsing:    true,
			DisableFlagsInUseLine: true,
			RunE: func(cmd *cobra.Command, args []string) error {
				v.OldVersion = toVersionValue(v.ExecVersion)
				v.State = v.CmdVersionInfo(args...)
				return v.State.GetState()
			},
		}
		v.SelfCmd.AddCommand(versionInfoCmd)
		versionInfoCmd.Example = cmdHelp.PrintExamples(v.SelfCmd, "")
		versionInfoCmd.Annotations = map[string]string{"group": Group}

		var versionLatestCmd = &cobra.Command{
			Use:                   CmdVersionLatest,
			Short:                 fmt.Sprintf("Info on latest version."),
			Long:                  fmt.Sprintf("Info on latest version."),
			DisableFlagParsing:    true,
			DisableFlagsInUseLine: true,
			RunE: func(cmd *cobra.Command, args []string) error {
				v.OldVersion = toVersionValue(v.ExecVersion)
				v.State = v.CmdVersionInfo(CmdVersionLatest)
				return v.State.GetState()
			},
		}
		v.SelfCmd.AddCommand(versionLatestCmd)
		versionLatestCmd.Example = cmdHelp.PrintExamples(v.SelfCmd, "")
		versionLatestCmd.Annotations = map[string]string{"group": Group}

		var versionUpdateCmd = &cobra.Command{
			Use:                   CmdVersionUpdate,
			Short:                 fmt.Sprintf("Update version of this executable."),
			Long:                  fmt.Sprintf("Check and update the latest version of this executable."),
			DisableFlagParsing:    true,
			DisableFlagsInUseLine: true,
			RunE: func(cmd *cobra.Command, args []string) error {
				v.OldVersion = toVersionValue(v.ExecVersion)
				v.State = v.CmdVersionUpdate()
				return v.State.GetState()
			},
		}
		v.SelfCmd.AddCommand(versionUpdateCmd)
		versionUpdateCmd.Example = cmdHelp.PrintExamples(v.SelfCmd, "")
		versionUpdateCmd.Annotations = map[string]string{"group": Group}

		if !disableVflag {
			v.cmd.Flags().BoolP(FlagVersion, "v", false, fmt.Sprintf("Display version of %s", v.ExecName))
		}

		v.cmd.SetVersionTemplate(DefaultVersionTemplate)
	}

	return v.State
}

func (v *Version) CmdVersion(cmd *cobra.Command, args ...string) State {
	for range Only.Once {
		v.CmdVersionShow()
		// su.SetHelp(cmd)
		_ = cmd.Help()
		v.State.SetOk("")
	}
	return v.State
}

func (v *Version) CmdVersionShow() State {
	v.PrintNameVersion()
	v.State.SetOk("")
	return v.State
}

func (v *Version) CmdVersionInfo(args ...string) State {
	for range Only.Once {
		if len(args) == 0 {
			fmt.Println("Showing details on current version.")
			args = []string{v.ExecVersion}
		}

		for _, vs := range args {
			fv := GetSemVer(vs)
			v.State = v.PrintVersion(fv)
			if v.State.IsNotOk() {
				break
			}
		}
	}

	return v.State
}

func (v *Version) CmdVersionList(args ...string) State {
	for range Only.Once {
		token := os.Getenv("GITHUB_TOKEN")
		if token == "" {
			token, _ = gitconfig.GithubToken()
		}

		var err error
		gh := github.NewClient(newHTTPClient(context.Background(), token))
		var rels []*github.RepositoryRelease
		rels, _, err = gh.Repositories.ListReleases(context.Background(), v.useRepo.Owner, v.useRepo.Name, nil)
		if err != nil {
			v.State.SetError(err.Error())
			break
		}

		for _, rel := range rels {
			v.State = v.PrintVersionSummary(*rel.TagName)
			if v.State.IsOk() {
				continue
			}
			v.State = v.PrintVersionSummary(*rel.Name)
			if v.State.IsOk() {
				continue
			}
			// WORKAROUND: (selfupdate) - If selfupdate.Release returns nil, then print direct.
			v.State = v.PrintSummary(rel)
		}
	}

	return v.State
}

func (v *Version) CmdVersionCheck() State {
	for range Only.Once {
		v.State = v.IsUpdated(true)
		if v.State.IsError() {
			break
		}
	}

	return v.State
}

func (v *Version) CmdVersionUpdate() State {
	for range Only.Once {
		v.State = v.IsUpdated(false)
		if v.State.IsError() {
			break
		}

		v.State = v.CreateDummyBinary()
		if v.State.IsError() {
			// Only break on error, NOT warning.
			break
		}

		v.State = v.UpdateTo(v.GetVersionValue())
		if v.State.IsNotOk() {
			break
		}

		if !v.AutoExec {
			break
		}

		// AutoExec will execute the new binary with the same args as given.
		v.State = v.AutoRun()
		if v.State.IsNotOk() {
			break
		}
	}

	return v.State
}

func (v *Version) FlagCheckVersion(cmd *cobra.Command) bool {
	var ok bool
	for range Only.Once {
		var fl *pflag.FlagSet
		if cmd == nil {
			fl = v.cmd.Flags()
		} else {
			fl = cmd.Flags()
		}

		// Show version.
		ok, _ = fl.GetBool(FlagVersion)
		if ok {
			v.CmdVersionShow()
			break
		}
	}
	return ok
}

func (v *Version) CmdHelp() State {
	// if state := su.IsNil(); state.IsError() {
	// 	return state
	// }

	err := v.SelfCmd.Help()
	if err != nil {
		v.State.SetError(err.Error())
	}
	return v.State
}

func _GetUsage(c *cobra.Command) string {
	var str string

	if c.Parent() == nil {
		str += fmt.Sprintf("%s ", c.Name())
	} else {
		str += fmt.Sprintf("%s ", c.Parent().Name())
		str += fmt.Sprintf("%s ", c.Use)
	}

	if c.HasAvailableSubCommands() {
		str += fmt.Sprintf("[command] ")
		str += fmt.Sprintf("<args> ")
	}

	return str
}

func _GetVersion(c *cobra.Command) string {
	var str string

	if c.Parent() == nil {
		//str = ux.SprintfBlue("%s ", su.Runtime.CmdName)
		//str += ux.SprintfCyan("v%s", su.Runtime.CmdVersion)
	}

	return str
}
