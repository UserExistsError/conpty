# conpty

Windows Pseudo Console (ConPTY) for Golang

See:

https://devblogs.microsoft.com/commandline/windows-command-line-introducing-the-windows-pseudo-console-conpty/

## Usage

```go
package main

import (
	"os"
	"io"
	"log"
	"github.com/UserExistsError/conpty"
)

func main() {
	commandLine := `c:\windows\system32\cmd.exe`
	cpty, err := conpty.Start(commandLine)
	if err != nil {
		log.Fatalf("Failed to spawn a pty:  %v", err)
	}

	go io.Copy(os.Stdout, cpty)
	io.Copy(cpty, os.Stdin)
	log.Printf("ExitCode %d", cpty.Close())
}
```