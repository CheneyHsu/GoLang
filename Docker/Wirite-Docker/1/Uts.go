package main 

import (
	"os/exec"
	"syscall"
	"os"
	"log"
)

func main {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNTS,
	}
	cmd.stdin = os.Stdin
	cmd.stdout = os.Stdout
	cmd.stderr = os.Stderr

	if err := cmd.Run();err != nil{
		log.Flags(err)

	}
}
