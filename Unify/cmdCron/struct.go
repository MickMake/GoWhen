package cmdCron

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdLog"
	"bytes"
	"errors"
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"sort"
	"strings"
	"time"
)

type Cron struct {
	Scheduler *gocron.Scheduler
	Job       *gocron.Job
	Error     error

	cmd     *cobra.Command
	SelfCmd *cobra.Command
}

func New() *Cron {
	var ret *Cron

	for range Only.Once {
		ret = &Cron{
			Scheduler: gocron.NewScheduler(time.Local),
			Error:     nil,

			cmd:     nil,
			SelfCmd: nil,
		}
	}

	return ret
}

func (c *Cron) GetCmd() *cobra.Command {
	return c.SelfCmd
}

func (c *Cron) PrintJobs() {
	for range Only.Once {
		cmdLog.Printf("PrintJobs: %s\n", time.Now().Format("2006/01/02 15:04:05"))

		crontab := make(map[string]*gocron.Job)
		var jobs []string
		for _, key := range c.Scheduler.Jobs() {
			name := strings.Join(key.Tags(), " ")
			crontab[name] = key
			jobs = append(jobs, name)
		}
		sort.Strings(jobs)

		buf := new(bytes.Buffer)
		table := tablewriter.NewWriter(buf)
		table.SetHeader([]string{"Job", "Last Run", "Next Run", "Run Count", "Running", "Error"})
		for _, key := range jobs {
			job := crontab[key]
			table.Append([]string{
				strings.Join(job.Tags(), " "),
				job.LastRun().Format("2006/01/02 15:04:05"),
				job.NextRun().Format("2006/01/02 15:04:05"),
				fmt.Sprintf("%d", job.RunCount()),
				fmt.Sprintf("%v", job.IsRunning()),
				fmt.Sprintf("%v", job.Error()),
				// job.ScheduledAtTime(),
			})
		}
		table.Render()
		cmdLog.Printf("\n%s", buf.String())
	}
}

func (c *Cron) AddJob(cron string, tag string, jobFunc interface{}, args ...interface{}) (*gocron.Job, error) {
	var job *gocron.Job

	for range Only.Once {
		job, c.Error = c.Scheduler.CronWithSeconds(cron).StartImmediately().Tag(tag).Do(jobFunc, args...)
	}

	return job, c.Error
}

func (c *Cron) StartBlocking() error {
	for range Only.Once {
		// fmt.Println(Cron.Scheduler.Location())
		// fmt.Println(Cron.Scheduler.Jobs())
		// fmt.Println(Cron.Scheduler.NextRun())

		c.Scheduler.RunAll()
		// PrintJobs()
		c.Scheduler.StartBlocking()

		if !c.Scheduler.IsRunning() {
			c.Error = errors.New("cron scheduler has not started")
			break
		}
	}

	return c.Error
}

func (c *Cron) Jobs() []*gocron.Job {
	var ret []*gocron.Job
	for range Only.Once {
		ret = c.Scheduler.Jobs()
	}
	return ret
}
