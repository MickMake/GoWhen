package cmdSystray

import (
	"GoWhen/Unify/Only"
	"github.com/getlantern/systray"
)


type Item struct {
	MenuItem    *systray.MenuItem
	Modes      Modes
	selectFunc SelectFunc
}
type SelectFunc  func(config *Config, item *Item) error

func DummySelect(config *Config, item *Item) error {
	// fmt.Println("DummySelect()")
	// fmt.Printf("Ref: %v\n", ref)
	return nil
}

func (i *Item) AddMode(mode string, ref *Mode) *Mode {

	for range Only.Once {
		if i == nil {
			break
		}

		if i.Modes == nil {
			i.Modes = make(Modes)
		}

		if mode == "" {
			break
		}

		if r, ok := i.Modes[mode]; ok {
			if r.title != "" {
				// Already exists and is valid - return
				break
			}
		}

		if ref == nil {
			if mode == ModeDefault {
				// fmt.Printf("%s: NewMode()\n", mode)
				ref = NewMode()
			} else {
				// fmt.Printf("%s: ModeDefault\n", mode)
				ref = i.Modes[ModeDefault]
				if ref == nil {
					ref = NewMode()
				}
			}
		} else {
			// fmt.Printf("%s: ref != nil\n", mode)
		}

		// Copy struct.
		ref = &Mode{
			title:   ref.title,
			tooltip: ref.tooltip,
			icon:    ref.icon,
			shown:   ref.shown,
			enabled: ref.enabled,
			checked: ref.checked,
		}

		i.Modes[mode] = ref
	}

	return ref
}

func (i *Item) AddSelectFunc(fn SelectFunc) {

	for range Only.Once {
		if i == nil {
			break
		}

		if fn == nil {
			fn = DummySelect
		}

		i.selectFunc = fn
	}
}

func (i *Item) Update(mode string) {

	for range Only.Once {
		if i == nil {
			break
		}
		if i.MenuItem == nil {
			break
		}
		if mode == "" {
			break
		}

		var ref *Mode
		var ok bool
		if ref, ok = i.Modes[mode]; !ok {
			break
		}

		if ref.title != "" {
			i.MenuItem.SetTitle(ref.title)
		}
		if ref.tooltip != "" {
			i.MenuItem.SetTooltip(ref.tooltip)
		}
		if ref.icon != nil {
			// Sets the icon of a menu item. Only available on Mac.
			i.MenuItem.SetIcon(ref.icon)
		}

		// if ref.shown == nil {
		// 	// Don't do anything.
		// } else
		if ref.shown {
			i.MenuItem.Show()
		} else {
			i.MenuItem.Hide()
		}

		// if ref.enabled == nil {
		// 	// Don't do anything.
		// } else
		if ref.enabled {
			i.MenuItem.Enable()
		} else {
			i.MenuItem.Disable()
		}

		if ref.checked == nil {
			// fmt.Printf("NIL:%v\n", ref.title)
			// Don't do anything.
		} else if *ref.checked == true {
			// fmt.Printf("CHECKED:%v\n", ref.title)
			i.MenuItem.Check()
		} else {
			// fmt.Printf("UNCHECKED:%v\n", ref.title)
			i.MenuItem.Uncheck()
		}
	}
}

func (i *Item) GetMode(mode string) *Mode {
	var ref *Mode

	for range Only.Once {
		if i == nil {
			break
		}

		if mode == "" {
			break
		}

		var ok bool
		if ref, ok = i.Modes[mode]; !ok {
			ref = nil
			break
		}

		if mode == ModeDefault {
			break
		}

		if ref.title == "" {
			ref = i.AddMode(mode, nil)
			// ref = i.Modes[ModeDefault]
		}
	}

	return ref
}

func (i *Item) SetEnable() {
	i.MenuItem.Enable()
}

func (i *Item) SetDisable() {
	i.MenuItem.Disable()
}

func (i *Item) SetTooltip(t string) {
	if t != "" {
		i.MenuItem.SetTooltip(t)
	}
}

func (i *Item) SetTitle(t string) {
	if t != "" {
		i.MenuItem.SetTitle(t)
	}
}

func (i *Item) SetChecked() {
	i.MenuItem.Check()
}

func (i *Item) SetUnchecked() {
	i.MenuItem.Uncheck()
}

func (i *Item) SetShow() {
	i.MenuItem.Show()
}

func (i *Item) SetHide() {
	i.MenuItem.Hide()
}

func (i *Item) SetIcon(icon []byte) {
	if len(icon) > 0 {
		i.MenuItem.SetIcon(icon)
	}
}

func (i *Item) IsChecked() bool {
	return i.MenuItem.Checked()
}

func (i *Item) IsUnchecked() bool {
	return !i.MenuItem.Checked()
}

func (i *Item) IsEnabled() bool {
	return !i.MenuItem.Disabled()
}

func (i *Item) IsDisabled() bool {
	return i.MenuItem.Disabled()
}

// func (me *Item) Enable(mode string) *Mode {
// 	var ref *Mode
//
// 	for range Only.Once {
// 		if me == nil {
// 			break
// 		}
//
// 		if mode == "" {
// 			break
// 		}
//
// 		me.Modes[mode].SetEnable()
// 	}
//
// 	return ref
// }
