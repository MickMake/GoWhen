package cmdDaemon

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdHelp"
	"GoWhen/Unify/cmdLog"
	"GoWhen/Unify/cmdVersion"
	"fmt"
	"github.com/sevlyar/go-daemon"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"syscall"
)

const Group = "Daemon"

func (d *Daemon) AttachCommands(cmd *cobra.Command) *cobra.Command {
	for range Only.Once {
		if cmd == nil {
			break
		}
		d.cmd = cmd

		// ******************************************************************************** //
		d.SelfCmd = &cobra.Command{
			Use:                   CmdDaemon,
			Aliases:               []string{""},
			Short:                 fmt.Sprintf("Daemonize commands."),
			Long:                  fmt.Sprintf("Daemonize commands."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               d.InitArgs,
			RunE:                  d.CmdDaemon,
			Args:                  cobra.MinimumNArgs(1),
		}
		cmd.AddCommand(d.SelfCmd)
		d.SelfCmd.Example = cmdHelp.PrintExamples(d.SelfCmd, "exec web run", "kill")
		d.SelfCmd.Annotations = map[string]string{"group": Group}

		// ******************************************************************************** //
		var cmdDaemonExec = &cobra.Command{
			Use:                   CmdDaemonExec,
			Aliases:               AliasesDaemonExec,
			Short:                 fmt.Sprintf("Execute commands as a daemon."),
			Long:                  fmt.Sprintf("Execute commands as a daemon."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               d.InitArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				return d.CmdDaemonExec(DummyFunc, nil, args)
			},
			Args: cobra.MinimumNArgs(1),
		}
		d.SelfCmd.AddCommand(cmdDaemonExec)
		cmdDaemonExec.Example = cmdHelp.PrintExamples(cmdDaemonExec, "")
		cmdDaemonExec.Annotations = map[string]string{"group": Group}

		// ******************************************************************************** //
		var cmdDaemonKill = &cobra.Command{
			Use:                   CmdDaemonStop,
			Aliases:               AliasesDaemonStop,
			Short:                 fmt.Sprintf("Terminate daemon."),
			Long:                  fmt.Sprintf("Terminate daemon."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               d.InitArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				d.Error = cmdLog.LogFileSet("")
				return d.CmdDaemonKill()
			},
			// Args:                  cobra.MinimumNArgs(1),
		}
		d.SelfCmd.AddCommand(cmdDaemonKill)
		cmdDaemonKill.Example = cmdHelp.PrintExamples(cmdDaemonKill, "")
		cmdDaemonKill.Annotations = map[string]string{"group": Group}

		// ******************************************************************************** //
		var cmdDaemonReload = &cobra.Command{
			Use:                   CmdDaemonReload,
			Aliases:               AliasesDaemonReload,
			Short:                 fmt.Sprintf("Reload daemon config."),
			Long:                  fmt.Sprintf("Reload daemon config."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               d.InitArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				d.Error = cmdLog.LogFileSet("")
				return d.CmdDaemonReload()
			},
			// Args:                  cobra.MinimumNArgs(1),
		}
		d.SelfCmd.AddCommand(cmdDaemonReload)
		cmdDaemonReload.Example = cmdHelp.PrintExamples(cmdDaemonReload, "")
		cmdDaemonReload.Annotations = map[string]string{"group": Group}

		// ******************************************************************************** //
		var cmdDaemonList = &cobra.Command{
			Use:                   CmdDaemonList,
			Aliases:               AliasesDaemonList,
			Short:                 fmt.Sprintf("List running daemon."),
			Long:                  fmt.Sprintf("List running daemon."),
			DisableFlagParsing:    false,
			DisableFlagsInUseLine: false,
			PreRunE:               d.InitArgs,
			RunE: func(cmd *cobra.Command, args []string) error {
				d.Error = cmdLog.LogFileSet("")
				return d.CmdDaemonList()
			},
			// Args:                  cobra.MinimumNArgs(1),
		}
		d.SelfCmd.AddCommand(cmdDaemonList)
		cmdDaemonList.Example = cmdHelp.PrintExamples(cmdDaemonList, "")
		cmdDaemonList.Annotations = map[string]string{"group": Group}
	}

	return d.SelfCmd
}

func (d *Daemon) InitArgs(_ *cobra.Command, _ []string) error {
	var err error
	for range Only.Once {
		// @TODO - Sort out this Daemon mess.
		d.cntxt = &daemon.Context{
			PidFileName: d.cmd.Name() + ".pid",
			PidFilePerm: 0644,
			LogFileName: d.cmd.Name() + ".log",
			LogFilePerm: 0640,
			WorkDir:     "./",
			Umask:       027,
			Args:        []string{fmt.Sprintf("[%s]", d.cmd.Name())},
			// 	Chroot:      "",
			// 	Env:         nil,
			// 	Credential:  nil,
		}
	}
	return err
}

func (d *Daemon) CmdDaemon(cmd *cobra.Command, _ []string) error {
	for range Only.Once {
		d.Error = cmd.Help()
	}

	return d.Error
}

func (d *Daemon) CmdDaemonExec(fn DaemonFunc, _ *cobra.Command, args []string) error {
	for range Only.Once {
		d.Error = cmdLog.LogFileSet("")

		var child *os.Process
		child, _ = d.cntxt.Search()
		if child != nil {
			fmt.Println("Daemon already running")
			break
		}

		d.cntxt.Args = []string{fmt.Sprintf("[%s]", d.cmd.Name())}
		d.cntxt.Args = append(d.cntxt.Args, args...)

		daemon.SetSigHandler(termHandler, syscall.SIGQUIT)
		daemon.SetSigHandler(termHandler, syscall.SIGTERM)
		daemon.SetSigHandler(reloadHandler, syscall.SIGHUP)

		// go worker()

		fmt.Printf("Starting daemon: %s\n", strings.Join(d.cntxt.Args, " "))
		child, d.Error = d.cntxt.Reborn()
		if d.Error != nil {
			// log.Printf("Error: %s\n", err)
			fmt.Println("Daemon already running.")

			pid := d.ReadPid()
			if pid != -1 {
				fmt.Printf("PID: %d\n", pid)
			}
			break
		}

		if child != nil {
			fmt.Printf("Daemon started. PID: %d\n", child.Pid)
			d.Error = d.WritePid(child.Pid)
			if d.Error != nil {
				break
			}
			break
		}
		//goland:noinspection GoDeferInLoop,GoUnhandledErrorResult
		defer d.cntxt.Release()

		// @TODO - Never seems to get to here!
		fmt.Println("Daemon started.")
		// Cmd.Error = fn(cmd, args)
	}

	return d.Error
}

func (d *Daemon) CmdDaemonKill() error {
	for range Only.Once {
		// if d.cntxt == nil {
		// 	d.Error = errors.New("daemon PID empty")
		// 	break
		// }
		//
		// if d.cntxt.PidFileName == "" {
		// 	d.Error = errors.New("daemon PID filename empty")
		// 	break
		// }
		//
		// pid := d.ReadPid()
		// if pid == -1 {
		// 	d.Error = errors.New("PID file empty or no PID file")
		// 	break
		// }
		//
		// fmt.Printf("Killing daemon. PID: %d\n", pid)
		// // Cmd.Error = syscall.Kill(pid, syscall.SIGTERM)
		// var child *os.Process
		// child, d.Error = os.FindProcess(pid)
		// if d.Error != nil {
		// 	break
		// }

		var child *os.Process
		child, d.Error = d.cntxt.Search()
		if d.Error != nil {
			break
		}
		if child == nil {
			fmt.Println("Daemon not running")
			break
		}
		fmt.Printf("Killing daemon. PID: %d\n", child.Pid)

		d.Error = child.Signal(syscall.SIGTERM)
		if d.Error != nil {
			break
		}

		d.Error = d.cntxt.Release()

		if cmdVersion.NewPath(d.cntxt.PidFileName).FileExists() {
			// @TODO - Workaround for Mac OSX.
			d.Error = os.Remove(d.cntxt.PidFileName)
			if d.Error != nil {
				break
			}
		}
	}

	return d.Error
}

func (d *Daemon) CmdDaemonReload() error {
	for range Only.Once {
		// pid := d.ReadPid()
		// if pid == -1 {
		// 	d.Error = errors.New("PID file empty or no PID file")
		// 	break
		// }
		//
		// fmt.Printf("Reloading daemon. PID: %d\n", pid)
		// // Cmd.Error = syscall.Kill(pid, syscall.SIGHUP)
		// var child *os.Process
		// child, d.Error = os.FindProcess(pid)
		// if d.Error != nil {
		// 	break
		// }

		var child *os.Process
		child, d.Error = d.cntxt.Search()
		if d.Error != nil {
			break
		}
		if child == nil {
			fmt.Println("Daemon not running")
			break
		}
		fmt.Printf("Reloading daemon. PID: %d\n", child.Pid)

		d.Error = child.Signal(syscall.SIGHUP)
		if d.Error != nil {
			break
		}
	}

	return d.Error
}

func (d *Daemon) CmdDaemonList() error {
	for range Only.Once {

		var child *os.Process
		child, d.Error = d.cntxt.Search()
		if d.Error != nil {
			break
		}

		pid := d.ReadPid()
		switch {
		// If no discovered PID and no PID file.
		case (child == nil) && (pid == -1):
			fmt.Println("No daemon running.")

		// If no discovered PID and a PID file.
		case (child == nil) && (pid != -1):
			fmt.Println("Removing stale PID file.")
			d.Error = os.Remove(d.cntxt.PidFileName)
			if d.Error != nil {
				break
			}

		// If discovered PID and no PID file.
		case (child != nil) && (pid == -1):
			fmt.Printf("Daemon running. PID: %d\n", child.Pid)
			fmt.Println("Creating PID file.")
			d.Error = d.WritePid(child.Pid)
			if d.Error != nil {
				break
			}

		// If discovered PID and a PID file.
		case (child != nil) && (pid != -1):
			fmt.Printf("Daemon running. PID: %d\n", child.Pid)
			if child.Pid == pid {
				break
			}
			fmt.Printf("Creating PID file. (Mismatch: %d != %d)\n", child.Pid, pid)
			d.Error = d.WritePid(child.Pid)
			if d.Error != nil {
				break
			}
		}
	}

	return d.Error
}

type DaemonFunc func(cmd *cobra.Command, args []string) error

func DummyFunc(_ *cobra.Command, _ []string) error {
	var err error
	for range Only.Once {
		//
	}
	return err
}
