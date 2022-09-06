package cmdSystray

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdConfig"
	"GoWhen/Unify/cmdVersion"
	"fmt"
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"github.com/spf13/cobra"
	"github.com/zserge/lorca"
	"log"
	"net/url"
	"reflect"
	"strings"
	"time"
)


type Config struct {
	items            map[string]*Item
	itemOrder        []string
	defaultItemOrder []string
	mode             string
	modes            map[string]bool
	Error            error
	PollFunc         PollFunc
	PollDelay        time.Duration
	// webView	         webview.WebView

	cmdConfig  *cmdConfig.Config
	cmdVersion *cmdVersion.Version
	cmd        *cobra.Command
	SelfCmd    *cobra.Command
}
type PollFunc func() error

func New(config *cmdConfig.Config, version *cmdVersion.Version) *Config {
	var c Config

	for range Only.Once {
		c = Config{
			items:      make(map[string]*Item),
			mode:       ModeDefault,
			cmdConfig:  config,
			cmdVersion: version,
			modes:      make(map[string]bool),
			// webView:    nil,
			Error:      nil,
			cmd:        nil,
			SelfCmd:    nil,
		}

		_ = c.initMenus()
	}

	return &c
}

func (c *Config) Run() error {
	var err error

	for range Only.Once {
		fmt.Println("Starting app...")
		systray.Run(c.onReady, c.onExit)
		// systray.RunWithAppWindow("mmAutoRecord", 950, 700, onReady, onExit)

		// https://stackoverflow.com/questions/12145301/nswindow-doesnt-receive-keyboard-events
		// Make sure to add the following to github.com/getlantern/systray/systray_browser_darwin.m
		//
		// ProcessSerialNumber psn = {0, kCurrentProcess};
		// OSStatus status = TransformProcessType(&psn, kProcessTransformToForegroundApplication);
	}

	return err
}

func (c *Config) initMenus() error {
	for range Only.Once {
		// Create menu items in order.
		var menuItem *Item
		var modeItem *Mode

		c.AddMode(ModeDefault)

		c.AddSeparator()

		menuItem = c.AddMenuItem(MenuConfig)
		menuItem.AddSelectFunc(c.ShowConfig)
		modeItem = menuItem.GetMode(ModeDefault)
		modeItem.SetTitle("Config")
		modeItem.SetTooltip(fmt.Sprintf("Configure %s.", c.cmdVersion.ExecName))
		// modeItem.SetHide()

		menuItem = c.AddMenuItem(MenuInfo)
		menuItem.AddSelectFunc(DummySelect)
		modeItem = menuItem.GetMode(ModeDefault)
		modeItem.SetTitle(fmt.Sprintf("%s v%s", c.cmdVersion.ExecName, c.cmdVersion.ExecVersion))
		modeItem.SetTooltip("Click here for more info.")
		// modeItem.SetEnable()

		menuItem = c.AddMenuItem(MenuQuit)
		menuItem.AddSelectFunc(c.Quit)
		modeItem = menuItem.GetMode(ModeDefault)
		modeItem.SetTitle("Quit")
		modeItem.SetTooltip(fmt.Sprintf("Exit %s.", c.cmdVersion.ExecName))
		modeItem.SetIcon(icon.Data)
		// modeItem.SetEnable()

		// Make these the default.
		c.defaultItemOrder = c.itemOrder
		c.itemOrder = []string{}
	}

	return nil
}

func (c *Config) ShowConfig(_ *Config, _ *Item) error {
	var err error
	for range Only.Once {
		txt := fmt.Sprintf(`
	<html>
		<head><title>Simple In/Out Desktop</title></head>
		<body><h1>Simple In/Out Desktop Client Config</h1></body>
		<pre>%s</pre>
	</html>
	`, c.cmdConfig.PrintConfig())
		var ui lorca.UI
		ui, err = lorca.New("data:text/html," + url.PathEscape(txt), "", 1000, 500)
		if err != nil {
			log.Fatal(err)
		}
		defer ui.Close()
		// Wait until UI window is closed
		<-ui.Done()
		fmt.Println("WebView exiting.")

		// Doesn't work - throws a SEGV.
		// runtime.LockOSThread()
		// c.webView = webview.New(true)
		// defer c.webView.Destroy()
		// c.webView.SetTitle("Simple In/Out Desktop Authorize")
		// c.webView.SetSize(800, 600, webview.HintNone)
		// // w.Navigate(u)
		// // w.Bind()
		// // w.Dispatch()
		// c.webView.SetHtml(fmt.Sprintf(`
		// %s
		// `,
		// c.cmdConfig.PrintConfig(),
		// ))
		// c.webView.Run()
		// fmt.Println("WebView exiting.")
	}
	return err
}

func (c *Config) Quit(_ *Config, _ *Item) error {
	systray.Quit()
	return nil
}

func (c *Config) SetPollFunc(fn PollFunc) {
	c.PollFunc = fn
}

func (c *Config) onReady() {
	for range Only.Once {
		var exit bool

		// Start web services
		go func() {
			if c.PollFunc != nil {
				if c.PollDelay == 0 {
					c.PollDelay = time.Second * 60
				}
				for !exit {
					err := c.PollFunc()
					if err != nil {
						exit = true
					}
					time.Sleep(c.PollDelay)
				}
			}
		}()

		systray.SetIcon(icon.Data)
		systray.SetTitle(c.cmdVersion.ExecName)
		systray.SetTooltip("Written by Mick Hellstrom (MickMake)")

		// Append default items to end.
		c.itemOrder = append(c.itemOrder, c.defaultItemOrder...)

		// Create systray menu items.
		for _, name := range c.itemOrder {
			if strings.HasPrefix(name, SeparatorPrefix) {
				// c.items[name].MenuItem =
				systray.AddSeparator()
				continue
			}

			t := c.items[name].Modes[ModeDefault].title
			tt := c.items[name].Modes[ModeDefault].tooltip
			c.items[name].MenuItem = systray.AddMenuItem(t, tt)
		}

		c.SetMode(ModeDefault)

		// Run through all channels waiting for select.
		for !exit {
			// var chans = make(chan struct{}, len(c.itemOrder))
			var chans []chan struct{}
			var chNames []string
			for i, name := range c.itemOrder {
				if strings.HasPrefix(name, SeparatorPrefix) {
					continue
				}

				chans = append(chans, c.items[name].MenuItem.ClickedCh)
				chNames = append(chNames, name)
				go c.produce(c.items[name].MenuItem.ClickedCh, i+1, name)
			}

			cases := make([]reflect.SelectCase, len(chans))
			for i, ch := range chans {
				cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
			}

			remaining := len(cases)
			for remaining > 0 {
				// chosen, value, ok := reflect.Select(cases)
				chosen, _, ok := reflect.Select(cases)
				if !ok {
					// The chosen channel has been closed, so zero out the channel to disable the case
					cases[chosen].Chan = reflect.ValueOf(nil)
					remaining -= 1
					continue
				}

				name := chNames[chosen]
				// fmt.Printf("Channel[%d]/[%s] %v and received %s\n", chosen, name, chans[chosen], value.String())
				fmt.Printf("Selected[%d]/[%s]:\n", chosen, name)
				if c.items[name].selectFunc != nil {
					err := c.items[name].selectFunc(c, c.items[name])
					if err != nil {
						continue
					}
				}
			}
		}
	}
}

func (c *Config) produce(ch chan<- struct{}, i int, name string) {
	for j := 0; j < len(c.itemOrder); j++ {
		// ch <- fmt.Sprintf("%d", i*10 + j)
		// fmt.Printf("Channel[%d]/[%s]: %v\n", i, name, ch)
		// fmt.Printf("Channel[%d]/[%s]\n", i, name)
	}
	// close(ch)
}

func (c *Config) onExit() {
	for range Only.Once {
		// err := c.cmdConfig.Write()
		// if err != nil {
		// 	fmt.Printf("ERROR: %s\n", err)
		// 	break
		// }
		fmt.Println("Finished")
	}
}

func (c *Config) AddMode(name string) {
	for range Only.Once {
		for _, m := range c.itemOrder {
			 // ref := c.items[m].GetMode(ModeDefault)
			 // if ref == nil {
				//  ref = NewMode()
			 // }
			// c.items[m].AddMode(name, ref)

			c.items[m].AddMode(name, nil)
		}
		c.modes[name] = true
	}
}

func (c *Config) AddMenuItem(name string) *Item {
	var menu *Item

	for range Only.Once {
		if name == "" {
			break
		}

		if _, ok := c.items[name]; !ok {
			menu = &Item{}
			c.items[name] = menu
		} else {
			menu = c.items[name]
		}

		for m := range c.modes {
			c.items[name].AddMode(m, nil)
		}
		c.itemOrder = append(c.itemOrder, name)
	}

	return menu
}

const SeparatorPrefix = "Separator"
func (c *Config) AddSeparator() *Item {
	var menu *Item

	for range Only.Once {
		name := fmt.Sprintf("%s%d", SeparatorPrefix, len(c.items))

		if _, ok := c.items[name]; !ok {
			menu = &Item{}
			c.items[name] = menu
		} else {
			menu = c.items[name]
		}

		ref := NewMode()
		ref.SetDisable()
		for m := range c.modes {
			c.items[name].AddMode(m, ref)
		}
		c.itemOrder = append(c.itemOrder, name)
	}

	return menu
}

func (c *Config) GetMenuItem(name string) *Item {
	var menu *Item

	for range Only.Once {
		if name == "" {
			break
		}

		if _, ok := c.items[name]; !ok {
			break
		}

		menu = c.items[name]
	}

	return menu
}

func (c *Config) SetMode(mode string) {

	for range Only.Once {
		if mode == "" {
			break
		}

		c.mode = mode
		for _, v := range c.itemOrder {
			c.items[v].Update(c.mode)
		}
	}
}

func (c *Config) GetMode() string {
	return c.mode
}
