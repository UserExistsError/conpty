package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/UserExistsError/conpty"
)

func main() {
	commandLine := `c:\windows\system32\cmd.exe`
	cpty, err := conpty.Start(commandLine)
	if err != nil {
		log.Fatalf("Failed to spawn a pty:  %v", err)
	}
	defer cpty.Close()

	go io.Copy(os.Stdout, cpty)
	go io.Copy(cpty, os.Stdin)

	exitCode, err := cpty.Wait(context.Background())
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("ExitCode: %d", exitCode)
}
