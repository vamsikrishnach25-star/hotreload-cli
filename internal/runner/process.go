package runner
import (
	"fmt"
	"log"
	"os"
	"os/exec"
)
type Process struct {
	Cmd *exec.Cmd
}
func (p *Process) Start(command string) error {

	log.Println("Starting server:", command)

	cmd := exec.Command("cmd", "/C", command)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return err
	}

	p.Cmd = cmd

	return nil
}
func (p *Process) Stop() {

	if p.Cmd == nil || p.Cmd.Process == nil {
		return
	}

	log.Println("Stopping previous server")

	pid := p.Cmd.Process.Pid

	cmd := exec.Command("taskkill", "/PID", fmt.Sprintf("%d", pid), "/T", "/F")

	err := cmd.Run()
	if err != nil {
		log.Println("Process already stopped or not found")
	}
}