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

	handler := func(r *eslogger.Row) bool {
		if o.Kind == "open" {
			fmt.Printf("%s [%d] opened %s\n", r.Process.Executable.Path, r.Process.ParentAuditToken.Pid, r.Event.Open.File.Path)
			return true
		}
		fmt.Printf("%s: %+v", o.Kind, r)
		return true
	}

	if err := eslogger.Run(o, handler); err != nil {
		log.Fatalf("process failed: %v", err)
	}
}
