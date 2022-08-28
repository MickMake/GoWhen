package cmdVersion

import (
	"GoWhen/Unify/Only"
	"fmt"
	"github.com/google/go-github/v30/github"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"path/filepath"
	"runtime"
	"strings"
)

func (v *Version) Update() State {
	for range Only.Once {
		if v.IsNotValid() {
			break
		}

		fmt.Printf("Checking '%s' for version greater than v%s\n", v.useRepo.GetUrl(), v.OldVersion.String())
		previous := v.OldVersion.ToSemVer()
		latest, err := v.ref.UpdateSelf(previous, v.useRepo.GetShortUrl())
		if err != nil {
			v.State.SetError(err.Error())
			break
		}

		if previous.Equals(latest.Version) {
			fmt.Printf("%s is up to date: v%s\n", v.useRepo.Name, latest.Version.String())
		} else {
			fmt.Printf("%s updated to v%s\n", v.useRepo.Name, latest.Version)
			if latest.ReleaseNotes != "" {
				fmt.Printf("%s %s Release Notes:%s\n", v.useRepo.Name, latest.Version, latest.ReleaseNotes)
			}
		}
	}

	return v.State
}

func (v *Version) PrintVersion(version *VersionValue) State {
	for range Only.Once {
		if v.IsNotValid() {
			break
		}

		release := v.FetchVersion(version)
		if v.State.IsNotOk() {
			break
		}
		if release == nil {
			break
		}

		fmt.Printf(printVersion(release))
	}

	return v.State
}

func (v *Version) IsUpdated(print bool) State {
	for range Only.Once {
		if v.IsNotValid() {
			break
		}

		latest := v.FetchVersion(v.GetVersionValue())
		if v.State.IsNotOk() {
			break
		}
		v.SetVersion(latest.Version.String())

		current := v.FetchVersion(v.OldVersion)
		if current == nil {
			if v.OldVersion.GT(latest.Version) {
				v.State.SetError("%s has version, (v%s), greater than repository, (v%s).",
					v.useRepo.Name,
					v.OldVersion.String(),
					latest.Version.String(),
				)
				if print {
					fmt.Printf("%s\n", v.State.GetError())
					fmt.Printf("Current version info unknown\n")
					fmt.Printf(printVersion(latest))
				}
				break
			}

			v.State.SetWarning("%s can be updated to v%s.",
				v.useRepo.Name,
				latest.Version.String())
			if print {
				fmt.Printf("%s\n", v.State.GetOk())
				fmt.Printf("Current version info unknown\n")
				fmt.Printf(printVersion(latest))
			}
			v.State.SetOk("")
			break
		}

		if current.Version.Equals(latest.Version) {
			v.State.SetOk("%s is up to date at v%s.",
				v.useRepo.Name,
				v.OldVersion.String())
			if print {
				fmt.Printf("%s\n", v.State.GetOk())
				fmt.Printf(printVersion(current))
			}
			break
		}

		if current.Version.LE(latest.Version) {
			v.State.SetWarning("%s can be updated to v%s.",
				v.useRepo.Name,
				latest.Version.String())
			if print {
				fmt.Printf("%s\n", v.State.GetWarning())
				fmt.Printf("Current version (v%s)\n", current.Version.String())
				fmt.Printf(printVersion(current))

				fmt.Printf("Updated version (v%s)\n", latest.Version.String())
				fmt.Printf(printVersion(latest))
			}
			break
		}

		if current.Version.GT(latest.Version) {
			v.State.SetWarning("%s is more recent at v%s, (latest is %s).",
				v.useRepo.Name,
				v.OldVersion.String(),
				latest.Version.String())
			if print {
				fmt.Printf("%s\n", v.State.GetWarning())
				fmt.Printf("Current version (v%s)\n", current.Version.String())
				fmt.Printf(printVersion(current))

				fmt.Printf("Updated version (v%s)\n", latest.Version.String())
				fmt.Printf(printVersion(latest))
			}
			break
		}
	}

	return v.State
}

func (v *Version) Set(s SelfUpdateArgs) State {
	if s.owner != nil {
		v.useRepo.Owner = *s.owner
	}

	if s.name != nil {
		v.useRepo.Name = *s.name
	}

	if s.version != nil {
		v.SetVersion(*s.version)
	}

	if s.binaryRepo != nil {
		_ = v.ExecBinaryRepo.Set(*s.binaryRepo)
	}

	if s.sourceRepo != nil {
		_ = v.ExecSourceRepo.Set(*s.sourceRepo)
	}

	if s.logging != nil {
		v.logging = (*FlagValue)(s.logging)
	} else {
		v.logging = &defaultFalse
	}

	if v.IsNotValid() {
		v.State.SetError("Invalid repo")
	}

	return v.State
}

func (v *Version) SetDebug(value bool) State {
	v.logging = (*FlagValue)(&value)
	if v.IsNotValid() {
		v.State.SetError("Invalid value")
	}
	return v.State
}

func (v *Version) SetName(value string) State {
	v.useRepo.Name = value
	if v.IsNotValid() {
		v.State.SetError("Invalid value")
	}
	return v.State
}

func (v *Version) GetName() string {
	return v.useRepo.Name
}

func (v *Version) SetOldVersion(value string) State {
	for range Only.Once {
		if v.OldVersion != nil {
			v.State.SetOk("")
			break
		}
		v.OldVersion = toVersionValue(value)
		if v.IsNotValid() {
			v.State.SetError("Invalid value")
		}
	}
	return v.State
}

func (v *Version) SetVersion(value string) State {
	if v.OldVersion == nil {
		v.OldVersion = v.useRepo.Version
	}

	v.useRepo.Version = toVersionValue(value)
	if v.IsNotValid() {
		v.State.SetError("Invalid value")
	}
	return v.State
}

func (v *Version) GetVersion() string {
	return v.useRepo.Version.String()
}

func (v *Version) FetchVersion(version *VersionValue) *selfupdate.Release {
	var release *selfupdate.Release

	for range Only.Once {
		if v.IsNotValid() {
			break
		}

		var ok bool
		var err error

		repo := v.useRepo.GetShortUrl()

		switch {
		case version.IsNotValid():
			fallthrough
		case version.IsLatest():
			release, ok, err = v.ref.DetectLatest(repo)

		default:
			vs := version.String()
			if !strings.HasPrefix(vs, "v") {
				vs = "v" + vs
			}
			release, ok, err = v.ref.DetectVersion(repo, vs)
		}

		if !ok {
			v.State.SetWarning("Version '%s' not found within '%s'", version.String(), v.useRepo.GetUrl())
			break
		}
		if err != nil {
			v.State.SetWarning("Version '%s' not found within '%s' - ERROR:%s", version.String(), v.useRepo.GetUrl(), err)
			break
		}

		// su.State.SetOutput(release)
	}

	return release
}

func (v *Version) SetRepo(value ...string) State {
	for range Only.Once {
		if v.OldVersion == nil {
			v.OldVersion = v.useRepo.Version
		}

		err := v.useRepo.Set(value...)
		if err != nil {
			v.State.SetError("Invalid value")
			break
		}

		v.config.Filters = addFilters(v.useRepo.Name, runtime.GOOS, runtime.GOARCH)
		v.ref, _ = selfupdate.NewUpdater(*v.config)

		v.TargetBinary = filepath.Join(v.CmdDir, v.useRepo.Name)

		v.State.SetOk("")
	}

	return v.State
}

func (v *Version) GetRepo() string {
	return v.useRepo.GetUrl()
}

func (v *Version) SetSourceRepo(value ...string) State {
	for range Only.Once {
		if v.OldVersion == nil {
			v.OldVersion = v.ExecSourceRepo.Version
		}

		_ = v.ExecSourceRepo.Set(value...)
		if v.IsNotValid() {
			v.State.SetError("Invalid value")
		}
	}
	return v.State
}

func (v *Version) GetSourceRepo() string {
	return v.ExecSourceRepo.GetUrl()
}

func (v *Version) SetBinaryRepo(value ...string) State {
	for range Only.Once {
		if v.OldVersion == nil {
			v.OldVersion = v.ExecBinaryRepo.Version
		}

		_ = v.ExecBinaryRepo.Set(value...)
		if v.IsNotValid() {
			v.State.SetError("Invalid value")
		}
	}
	return v.State
}

func (v *Version) GetBinaryRepo() string {
	return v.ExecBinaryRepo.GetUrl()
}

func (v *Version) UpdateTo(newVersion *VersionValue) State {
	for range Only.Once {
		if v.IsNotValid() {
			break
		}

		newRelease := v.FetchVersion(newVersion)
		if newRelease == nil {
			fmt.Printf("Version v%s doesn't exist for '%s'\n", v.useRepo.Version.String(), v.useRepo.Name)
			break
		}
		fmt.Printf("Updated version v%s => v%s\n", v.OldVersion.String(), newRelease.Version.String())

		err := v.ref.UpdateTo(newRelease, v.TargetBinary)
		if err != nil {
			v.State.SetError(err.Error())
			break
		}

		fmt.Printf("%s updated from v%s to v%s\n", v.useRepo.Name, v.OldVersion.String(), newRelease.Version.String())
	}

	return v.State
}

func (v *Version) GetVersionValue() *VersionValue {
	var ret *VersionValue

	for range Only.Once {
		ret = v.useRepo.Version
	}

	return ret
}

func (v *Version) PrintVersionSummary(version string) State {
	for range Only.Once {
		if v.IsNotValid() {
			break
		}

		var release *selfupdate.Release
		if version == LatestVersion {
			release = v.FetchVersion(nil)
		} else {
			release = v.FetchVersion(toVersionValue(version))
		}

		if release == nil {
			v.State.SetWarning("Version '%s' (%s) is empty.", version, toVersionValue(version))
			break
		}

		if version != release.Name {
			v.State.SetWarning("Version '%s' (%s) differs to repo '%s'.", version, toVersionValue(version), release.Name)
			// Don't show mismatched versions.
			break
		}

		fmt.Print(printVersionSummary(release))
	}

	return v.State
}

func (v *Version) PrintSummary(release *github.RepositoryRelease) State {
	for range Only.Once {
		if v.IsNotValid() {
			break
		}

		if release == nil {
			break
		}

		fmt.Printf("\nExecutable: ")
		fmt.Printf("%s ", v.useRepo.Name)
		fmt.Printf("%s\n", release.GetName())

		fmt.Printf("Url: ")
		fmt.Printf("%s\n", release.GetHTMLURL())

		fmt.Printf("Binary Size: ")
		fmt.Printf("unknown")

		fmt.Printf("Published Date: ")
		fmt.Printf("%s\n", release.GetPublishedAt().String())
	}

	return v.State
}
