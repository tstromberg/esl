// The "opens" tool shows file opens
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tstromberg/esl/pkg/eslogger"
)

func main() {
	o := eslogger.Options{
		Kind: os.Args[1],
	}

	fmt.Printf("listening for %s events ...\n", o.Kind)

	handler := func(r *eslogger.Row) bool {
		switch o.Kind {
		case "open":
			fmt.Printf("%s [%d] opened %s\n", r.Process.Executable.Path, r.Process.AuditToken.Pid, r.Event.Open.File.Path)
		case "exec":
			fmt.Printf("%s [%d] launched %s [%d]\n", r.Process.Executable.Path, r.Process.Ppid, r.Event.Exec.Target.Executable.Path, r.Process.AuditToken.Pid)
		default:
			fmt.Printf("%s: %+v", o.Kind, r)
		}
		return true
	}

	if err := eslogger.Run(o, handler); err != nil {
		log.Fatalf("process failed: %v", err)
	}
}
