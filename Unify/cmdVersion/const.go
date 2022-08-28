package cmdVersion

const (
	errorNoRepo    = "repo is not defined - selfupdate disabled"
	errorNoVersion = "no versions in repo - selfupdate disabled"
	LatestVersion  = "latest"
	LatestSemVer   = "0.0.0"
	CurrentVersion = "current"
	EarliestSemVer = "0.0.1" // Using semver, "0.0.0" is defined as "latest".

	CmdSelfUpdate = "selfupdate"

	CmdVersion       = "version"
	CmdVersionInfo   = "info"
	CmdVersionList   = "list"
	CmdVersionLatest = "latest"
	CmdVersionCheck  = "check"
	CmdVersionUpdate = "update"

	FlagVersion = "version"

	DefaultRepoServer = "github.com"

	BootstrapBinaryName = "bootstrap"
	DefaultVersion      = "0.4.2"
)

var defaultFalse = FlagValue(false)

const DefaultVersionTemplate = `
{{with .Name}}{{printf "%s " .}}{{end}} {{printf "version %s" .Version}}
`
