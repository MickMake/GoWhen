package cmdDaemon

const (
	// pidFile = defaults.BinaryName + ".pid"

	CmdDaemon       = "daemon"
	CmdDaemonExec   = "exec"
	CmdDaemonStop   = "kill"
	CmdDaemonReload = "reload"
	CmdDaemonList   = "list"
)

var (
	AliasesDaemonExec   = []string{"run"}
	AliasesDaemonStop   = []string{"stop"}
	AliasesDaemonReload = []string{"hup"}
	AliasesDaemonList   = []string{"ls"}
)
