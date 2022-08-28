package cmdDaemon

import (
	"GoWhen/Unify/Only"
	"GoWhen/Unify/cmdVersion"
	"fmt"
	"github.com/sevlyar/go-daemon"
	"log"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func (d *Daemon) ReadPid() int {
	ret := -1

	for range Only.Once {
		if d.cntxt == nil {
			break
		}

		if d.cntxt.PidFileName == "" {
			break
		}

		if !cmdVersion.NewPath(d.cntxt.PidFileName).FileExists() {
			// if !mmWebcam.FileExists(pidFile) {
			ret = -1
			break
		}

		// Open PID file
		var pid []byte
		pid, d.Error = os.ReadFile(d.cntxt.PidFileName)
		if d.Error != nil {
			ret = -1
			break
		}

		ps := strings.TrimSpace(string(pid))
		ret, d.Error = strconv.Atoi(ps)
		if d.Error != nil {
			ret = -1
			break
		}
	}

	return ret
}

func (d *Daemon) WritePid(pid int) error {
	for range Only.Once {
		// Open a file for writing
		var file *os.File
		file, d.Error = os.Create(d.cntxt.PidFileName)
		if d.Error != nil {
			break
		}
		//goland:noinspection GoDeferInLoop,GoUnhandledErrorResult
		defer file.Close()

		_, d.Error = file.Write([]byte(fmt.Sprintf("%d", pid)))
		if d.Error != nil {
			break
		}
	}

	return d.Error
}

//goland:noinspection GoUnusedExportedFunction
func DaemonizeClose(cntxt *daemon.Context) error {
	return cntxt.Release()
}

func worker() {
	fmt.Println("DEBUG")
LOOP:
	for {
		time.Sleep(time.Second) // this is work to be done by worker.
		select {
		case <-stop:
			break LOOP
		default:
		}
	}
	done <- struct{}{}
}

var (
	stop = make(chan struct{})
	done = make(chan struct{})
)

func termHandler(sig os.Signal) error {
	log.Println("Daemon terminating...")
	stop <- struct{}{}
	if sig == syscall.SIGQUIT {
		<-done
	}
	return daemon.ErrStop
}

func reloadHandler(_ os.Signal) error {
	log.Println("Daemon configuration reloaded")
	return nil
}
