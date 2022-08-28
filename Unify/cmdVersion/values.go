package cmdVersion

import (
	"GoWhen/Unify/Only"
	"errors"
	"fmt"
	"github.com/blang/semver"
	"net/url"
	"strings"
)

type StringValue string
type UrlValue struct {
	*url.URL
	Owner   string
	Name    string
	Version *VersionValue
}

// type VersionValue struct {
// 	semver.Version
// 	string
// }
type VersionValue struct {
	semver.Version
	string
}

type FlagValue bool

// func ReflectStringValue(ref interface{}) *StringValue {
// 	var ret *StringValue
// 	switch ref.(type) {
// 		case *[]byte:
// 			ret = ref.(*StringValue)
// 		case *string:
// 			ret = ref.(*StringValue)
// 		case []byte:
// 			ret = ref.(*StringValue)
// 		case string:
// 			ret = ref.(*StringValue)
// 	}
// 	return ret
// }
//
//
// func ReflectVersionValue(ref interface{}) *VersionValue {
// 	var ret VersionValue
// 	switch ref.(type) {
// 		case *[]byte:
// 			ret.string = fixVersion(*(ref.(*string)))
// 			ret.Version, _ = semver.Parse(ret.string)
// 		case *string:
// 			ret.string = fixVersion(*(ref.(*string)))
// 			ret.Version, _ = semver.Parse(ret.string)
// 		case []byte:
// 			ret.string = fixVersion(ref.(string))
// 			ret.Version, _ = semver.Parse(ret.string)
// 		case string:
// 			ret.string = fixVersion(ref.(string))
// 			ret.Version, _ = semver.Parse(ret.string)
// 	}
// 	return &ret
// }
//
//
// func ReflectFlagValue(ref interface{}) *FlagValue {
// 	var ret *FlagValue
// 	switch ref.(type) {
// 		case *bool:
// 			ret = ref.(*FlagValue)
// 	}
// 	return ret
// }

func GetSemVer(v string) *VersionValue {
	var sver VersionValue
	if v == LatestVersion {
		return &sver
	}
	v = dropVprefix(v)
	sv, _ := semver.Parse(v)
	sver.Version = sv
	sver.string = sv.String()
	return &sver
}

func (v *VersionValue) String() string {
	if v == nil {
		return ""
	}
	return dropVprefix(v.Version.String())
}

func toVersionValue(version string) *VersionValue {
	ret := &VersionValue{}
	for range Only.Once {
		if version == LatestVersion {
			ret.string = LatestVersion
			break
		}

		if version == "" {
			ret.string = LatestVersion
			break
		}

		version = fixVersion(version)
		var err error
		ret.Version, err = semver.Parse(version)
		if err != nil {
			ret = nil
			break
		}

		ret.string = ret.Version.String()
	}
	return ret
}

func toOwnerValue(s string) *StringValue {
	s = stripUrlPrefix(s)
	if strings.Contains(s, "/") {
		sa := strings.Split(s, "/")
		switch {
		case len(sa) == 0:
			// Nada
		default:
			s = sa[0]
		}

	}
	v := StringValue(s)
	return &v
}

func (v *VersionValue) ToSemVer() semver.Version {
	return v.Version
}

func (v *VersionValue) IsValid() bool {
	var ok bool
	for range Only.Once {
		if v == nil {
			break
		}

		err := v.Version.Validate()
		if err != nil {
			break
		}

		ok = true
	}
	return ok
}

func (v *VersionValue) IsNotValid() bool {
	return !v.IsValid()
}

func (v *VersionValue) IsLatest() bool {
	var ok bool
	for range Only.Once {
		cv := v.Version.String()
		if cv == LatestVersion {
			ok = true
			break
		}

		if cv == LatestSemVer {
			ok = true
			break
		}

		lv, _ := semver.Parse("")
		if cv == lv.String() {
			ok = true
			break
		}
	}
	return ok
}

func (v *StringValue) String() string {
	return string(*v)
}

func toStringValue(s string) *StringValue {
	v := StringValue(s)
	return &v
}

func toBoolValue(b bool) *FlagValue {
	v := FlagValue(b)
	return &v
}

func (v *StringValue) IsValid() bool {
	var ok bool
	for range Only.Once {
		if v == nil {
			break
		}
		if *v == "" {
			break
		}
		ok = true
	}
	return ok
}
func (v *StringValue) IsNotValid() bool {
	return !v.IsValid()
}

func (v *StringValue) IsNil() bool {
	if v == nil {
		return true
	}
	return false
}
func (v *StringValue) IsNotNil() bool {
	return !v.IsNil()
}

func (v *StringValue) IsEmpty() bool {
	if v == nil {
		return true
	}
	return false
}
func (v *StringValue) IsNotEmpty() bool {
	return !v.IsEmpty()
}

func (v *UrlValue) String() string {
	return v.URL.String()
}
func toUrlValue(s ...string) UrlValue {
	v := UrlValue{}
	_ = v.Set(s...)
	return v
}

func (v *UrlValue) IsValid() bool {
	var ok bool
	for range Only.Once {
		if v == nil {
			break
		}
		if v.String() == "" {
			break
		}
		ok = true
	}
	return ok
}
func (v *UrlValue) IsNotValid() bool {
	return !v.IsValid()
}

func (v *UrlValue) IsNil() bool {
	if v == nil {
		return true
	}
	return false
}
func (v *UrlValue) IsNotNil() bool {
	return !v.IsNil()
}

func (v *UrlValue) IsEmpty() bool {
	if v == nil {
		return true
	}
	return false
}
func (v *UrlValue) IsNotEmpty() bool {
	return !v.IsEmpty()
}

func (v *UrlValue) Set(args ...string) error {
	var err error

	for range Only.Once {
		if len(args) == 0 {
			break
		}

		repoString := addUrlPrefix(args...)

		v.URL, err = url.Parse(repoString)
		if err != nil {
			break
		}

		repoArgs := strings.Split(v.Path, "/")
		switch len(repoArgs) {
		case 0:
			err = errors.New(fmt.Sprintf("Url path empty"))
		case 1:
			err = errors.New(fmt.Sprintf("Url path invalid"))

		case 2:
			err = errors.New(fmt.Sprintf("Url missing repo name"))

		case 3:
			// Assume we have also been given a repo name.
			v.Owner = repoArgs[1]
			v.Name = repoArgs[2]
			if v.Version.String() == "" {
				v.Version = toVersionValue(LatestVersion)
			}

		default:
			// Assume we have also been given a repo version.
			v.Owner = repoArgs[1]
			v.Name = repoArgs[2]
			switch repoArgs[3] {
			case "":
				fallthrough
			case LatestVersion:
				//

			default:
				v.Version = toVersionValue(dropVprefix(repoArgs[3]))
			}
		}

	}

	return err
}

func (v *UrlValue) GetShortUrl() string {
	return fmt.Sprintf("%s/%s", v.Owner, v.Name)
}

func (v *UrlValue) GetUrl() string {
	return fmt.Sprintf("%s://%s/%s/%s", v.Scheme, v.Host, v.Owner, v.Name)
}
