package runner

import (
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

	p.Cmd.Process.Kill()
}