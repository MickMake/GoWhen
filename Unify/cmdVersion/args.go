package cmdVersion

import (
	"GoWhen/Unify/Only"
	"strings"
)

func (v *Version) SetArgs(a ...string) {
	v.Args.Set(a...)
	//var err error
	//
	//for range onlyOnce {
	//	r.Args = a
	//}
	//
	//return err
}

func (v *Version) AddArgs(a ...string) {
	v.Args.Append(a...)
	//var err error
	//
	//for range onlyOnce {
	//	r.Args = append(r.Args, a...)
	//}
	//
	//return err
}

func (v *Version) GetArgs() []string {
	return v.Args.GetAll()
}

func (v *Version) GetArg(index int) string {
	return v.Args.Get(index)
}

func (v *Version) GetArgRange(lower int, upper int) []string {
	return v.Args.Range(lower, upper)
}

func (v *Version) SprintfArgRange(lower int, upper int) string {
	return v.Args.SprintfRange(lower, upper)
}

func (v *Version) SprintfArgsFrom(lower int) string {
	return v.Args.SprintfFrom(lower)
}

//goland:noinspection SpellCheckingInspection
func (v *Version) GetNargs(begin int, size int) []string {
	return v.Args.GetFromSize(begin, size)
}

//goland:noinspection SpellCheckingInspection
func (v *Version) SprintfNargs(lower int, upper int) string {
	return v.Args.SprintfFromSize(lower, upper)
}

func (v *Version) SetFullArgs(a ...string) {
	v.FullArgs.Set(a...)
	//var err error
	//
	//for range onlyOnce {
	//	r.FullArgs = a
	//}
	//
	//return err
}

func (v *Version) AddFullArgs(a ...string) {
	v.FullArgs.Append(a...)
	//var err error
	//
	//for range onlyOnce {
	//	r.FullArgs = append(r.FullArgs, a...)
	//}
	//
	//return err
}

func (v *Version) GetFullArgs() []string {
	return v.FullArgs.GetAll()
}

type ExecArgs []string

func (r *ExecArgs) ToString() string {
	return strings.Join(*r, " ")
}

func (r *ExecArgs) String() string {
	return strings.Join(*r, " ")
}

func (r *ExecArgs) Set(a ...string) *ExecArgs {
	for range Only.Once {
		*r = a
	}
	return r
}

func (r *ExecArgs) Add(a ...string) *ExecArgs {
	*r = append(*r, a...)
	return r
}

func (r *ExecArgs) Append(a ...string) *ExecArgs {
	*r = append(*r, a...)
	return r
}

func (r *ExecArgs) Get(index int) string {
	var ret string

	for range Only.Once {
		if index > len(*r)-1 {
			break
		}

		ret = (*r)[index]
	}

	return ret
}

func (r *ExecArgs) GetAll() []string {
	return *r
}

func (r *ExecArgs) Sprintf() string {
	return strings.Join(*r, " ")
}

func (r *ExecArgs) Range(lower int, upper int) []string {
	var ret []string

	for range Only.Once {
		if lower < 0 {
			break
		}

		if upper <= 0 {
			break
		}

		upper++

		if lower > upper {
			break
		}

		as := len(*r) - 1

		if lower > as {
			break
		}

		if upper > as {
			// @TODO - Should we pad out this array to the full 'size' if it's less?
			ret = (*r)[lower:]
			break
		}

		if lower == upper {
			ret = []string{(*r)[lower]}
			break
		}

		ret = (*r)[lower:upper]
	}

	return ret
}

func (r *ExecArgs) SprintfRange(lower int, upper int) string {
	return strings.Join(r.Range(lower, upper), " ")
}

func (r *ExecArgs) GetFromSize(begin int, size int) []string {
	return r.Range(begin, begin+size-1)
}

func (r *ExecArgs) SprintfFromSize(lower int, upper int) string {
	return strings.Join(r.GetFromSize(lower, upper), " ")
}

func (r *ExecArgs) GetFrom(lower int) []string {
	return r.Range(lower, len(*r))
}

func (r *ExecArgs) SprintfFrom(lower int) string {
	return strings.Join(r.GetFrom(lower), " ")
}
