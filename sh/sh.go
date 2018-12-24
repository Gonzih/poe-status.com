package sh

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func Sh(args ...string) ([]byte, error) {
	log.Printf("$ %s", strings.Join(args, " "))
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	return cmd.Output()
}
