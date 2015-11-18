package shcmd

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
	"time"
)

// RunWithin executes a shell cmd within the time in Duration or return when
// command is finished. Note the function is blocking
func RunWithin(cmdstr string, d time.Duration) (*bytes.Buffer, error) {
	parts := strings.Fields(cmdstr)
	cmd := exec.Command(parts[0], parts[1:]...)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Start(); err != nil {
		return &out, err
	}
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()
	select {
	case <-time.After(d):
		err := cmd.Process.Kill()
		log.Fatal("task has been killed: ", err)
		<-done
		return &out, err
	case err := <-done:
		return &out, err
	}
}
