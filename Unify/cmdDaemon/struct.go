package cmdDaemon

import (
	"GoWhen/Unify/Only"
	"github.com/sevlyar/go-daemon"
	"github.com/spf13/cobra"
)

type Daemon struct {
	Error error
	cntxt *daemon.Context

	cmd     *cobra.Command
	SelfCmd *cobra.Command
}

func New() *Daemon {
	var ret *Daemon

	for range Only.Once {
		ret = &Daemon{
			Error: nil,
			cntxt: &daemon.Context{},

			cmd:     nil,
			SelfCmd: nil,
		}
	}

	return ret
}

func (d *Daemon) GetCmd() *cobra.Command {
	return d.SelfCmd
}
