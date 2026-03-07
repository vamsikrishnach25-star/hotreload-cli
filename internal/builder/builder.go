package builder
import (
	"log"
	"os/exec"
)
func RunBuild(cmdStr string) error {
	log.Println("Running build:", cmdStr)
	cmd := exec.Command("cmd", "/C", cmdStr)
	cmd.Stdout = nil
	cmd.Stderr = nil
	err := cmd.Run()
	if err != nil {
		log.Println("Build failed:", err)
		return err
	}
	log.Println("Build successful")
	return nil
}
