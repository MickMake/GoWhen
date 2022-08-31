package cal

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdVersion"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"regexp"
	"sort"
	"time"
)


type Convert struct {
	Filename string
	Data []struct {
		Go    string `json:"go"`
		Java  string `json:"java"`
		C     string `json:"c"`
		Notes string `json:"notes"`
	}
}

func (c *Convert) String() string {
	var ret string
	for range Only.Once {
		if c == nil {
			break
		}

		buf := new(bytes.Buffer)
		table := tablewriter.NewWriter(buf)
		table.SetHeader([]string{"GoLang layout", "Java notation", "C/CPP notation", "Notes"})
		table.SetBorder(true)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		for _, e := range c.Data {
			table.Append([]string{e.Go, e.Java, e.C, e.Notes})
		}
		table.Render()

		ret += "Conversion table:\n"
		ret += buf.String()
	}
	return ret
}

func (c *Convert) FromJava(format string) string {
	for range Only.Once {
		if format == "" {
			break
		}

		for _, e := range c.Data {
			if e.Go == "" {
				continue
			}

			// n := e.Java
			// fmt.Println(n)
			re := regexp.MustCompile(fmt.Sprintf("\\b%s\\b", e.Java))

			format = re.ReplaceAllString(format, e.Go)
			// fmt.Sprintf("")
			// format = strings.ReplaceAll(format, e.Java, e.Go)
		}
	}
	return format
}

func (c *Convert) FromCpp(format string) string {
	for range Only.Once {
		if format == "" {
			break
		}

		for _, e := range c.Data {
			if e.Go == "" {
				continue
			}

			// n := strings.ReplaceAll(e.C, "%", "%%")
			// n := e.C
			// fmt.Println(n)
			re := regexp.MustCompile(e.C)

			format = re.ReplaceAllString(format, e.Go)
			// fmt.Sprintf("")
			// format = strings.ReplaceAll(format, e.Java, e.Go)
		}
	}
	return format
}

func (c *Convert) ReadFile() error {
	var err error
	for range Only.Once {
		j := []byte(DefaultConvertConfig)

		if cmdVersion.NewPath(c.Filename).FileExists() {
			j, err = os.ReadFile(c.Filename)
			if err != nil {
				break
			}
		} else {
			err = c.WriteFile(j)
			if err != nil {
				break
			}
		}

		err = json.Unmarshal(j, &c.Data)
		if err != nil {
			break
		}
	}
	return err
}

func (c *Convert) WriteFile(data []byte) error {
	var err error
	for range Only.Once {
		// Open a file for writing
		var file *os.File
		file, err = os.Create(c.Filename)
		if err != nil {
			break
		}
		//goland:noinspection GoDeferInLoop,GoUnhandledErrorResult
		defer file.Close()

		_, err = file.Write(data)
		if err != nil {
			break
		}
	}

	return err
}

func ReadConvert(fn string) (*Convert, error) {
	var ret Convert
	var err error
	for range Only.Once {
		ret.Filename = fn
		err = ret.ReadFile()
	}
	return &ret, err
}

func PrintLayoutOptions(t *time.Time) string {
	var ret string
	for range Only.Once {
		now := time.Now()
		if t != nil {
			now = *t
			ret += fmt.Sprintf("All GoLang layout options for the date \"%s\":\n", now.Format(time.RFC1123))
		} else {
			ret += "All GoLang layout options:\n"
		}

		var sorted []string
		for k := range LayoutOptions {
			sorted = append(sorted, k)
		}
		sort.Strings(sorted)
		buf := new(bytes.Buffer)
		table := tablewriter.NewWriter(buf)
		if t != nil {
			table.SetHeader([]string{"GoLang layout", "Value", "Label"})
		} else {
			table.SetHeader([]string{"GoLang layout", "Label"})
		}
		table.SetBorder(true)
		table.SetAlignment(tablewriter.ALIGN_LEFT)

		for _, k := range sorted {
			if t != nil {
				table.Append([]string{
					LayoutOptions[k],
					now.Format(LayoutOptions[k]),
					k,
				})
			} else {
				table.Append([]string{
					LayoutOptions[k],
					k,
				})
			}
		}

		table.Render()
		ret += buf.String()
	}
	return ret
}

func PrintLayouts() string {
	var ret string
	for range Only.Once {
		ret += "All GoLang layouts:\n"

		var sorted []string
		for k := range Layouts {
			sorted = append(sorted, k)
		}
		sort.Strings(sorted)
		buf := new(bytes.Buffer)
		table := tablewriter.NewWriter(buf)
		table.SetHeader([]string{"Label", "GoLang layout"})
		table.SetBorder(true)
		table.SetAlignment(tablewriter.ALIGN_LEFT)

		for _, k := range sorted {
			table.Append([]string{
				k,
				Layouts[k],
			})
		}

		table.Render()
		ret += buf.String()
	}
	return ret
}
