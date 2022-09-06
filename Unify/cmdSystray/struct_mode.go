package cmdSystray

import (
	"GoWhen/Unify/Only"
)


type Modes map[string]*Mode

type Mode struct {
	title string
	tooltip string
	icon []byte

	shown bool
	checked *bool
	enabled bool
	// shown TriState
	// checked TriState
	// enabled TriState
}

func NewMode() *Mode {
	ret := Mode{
		title:   "",
		tooltip: "",
		icon:    nil,
		shown:   true,	// Default to shown.
		enabled: true,	// Default to enabled.
		checked: nil,
	}
	return &ret
}

func (m *Mode) SetEnable() {
	for range Only.Once {
		if m == nil {
			break
		}
		m.enabled = true
	}
}
func (m *Mode) SetDisable() {
	for range Only.Once {
		if m == nil {
			break
		}
		m.enabled = false
	}
}

func (m *Mode) SetTooltip(t string) {
	for range Only.Once {
		if m == nil {
			break
		}
		m.tooltip = t
	}
}

func (m *Mode) SetTitle(t string) {
	for range Only.Once {
		if m == nil {
			break
		}
		m.title = t
	}
}

func (m *Mode) SetChecked() {
	for range Only.Once {
		if m == nil {
			break
		}

		checked := true
		m.checked = &checked
	}
}
func (m *Mode) SetUnchecked() {
	for range Only.Once {
		if m == nil {
			break
		}

		checked := false
		m.checked = &checked
	}
}

func (m *Mode) SetShow() {
	for range Only.Once {
		if m == nil {
			break
		}
		m.shown = true
	}
}
func (m *Mode) SetHide() {
	for range Only.Once {
		if m == nil {
			break
		}
		m.shown = false
	}
}

func (m *Mode) SetIcon(icon []byte) {
	for range Only.Once {
		if m == nil {
			break
		}
		m.icon = icon
	}
}

// type TriState bool
// func (me *TriState) IsTrue() bool {
// 	if *me == true {
// 		return true
// 	}
// 	return false
// }
// func (me *TriState) IsFalse() bool {
// 	if *me == false {
// 		return true
// 	}
// 	return false
// }
// func (me *TriState) IsIgnore() bool {
// 	if me == nil {
// 		return true
// 	}
// 	return false
// }
