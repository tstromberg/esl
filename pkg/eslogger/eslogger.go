// eslogger allows you to read EndpointSecurity events via eslogger
//
// IMPORTANT NOTE FROM eslogger(1):
// eslogger is NOT API in any sense.  Do NOT rely on the structure or information emitted for
// ANY reason.  It may change from release to release without warning.
package eslogger

import (
	"bufio"
	"encoding/json"
	"os/exec"
)

type HandlerFunc func(r *Row) bool
type Options struct {
	Kind string
}

// Run executes eslogger and streams events to a handler function
func Run(o Options, h HandlerFunc) error {
	cmd := exec.Command("sudo", "eslogger", o.Kind)
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		t := scanner.Text()
		r := &Row{}
		// log.Printf("t: %s", t)
		if err := json.Unmarshal([]byte(t), r); err != nil {
			return err
		}

		if !h(r) {
			return nil
		}
	}
	return cmd.Wait()
}
