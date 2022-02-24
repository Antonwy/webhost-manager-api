package cli

import (
	"bufio"
	"errors"
	"log"
	"os/exec"
	"strings"
)

const (
	PortAlreadyAllocated = "Port Is Already Allocated"
	DockerComposePath    = "/usr/local/bin"
	ErrorStart           = "Error response from daemon: "
	ExitStatus           = "exit status "
)

func Run(cmd *exec.Cmd) error {
	r, _ := cmd.StdoutPipe()

	cmd.Stderr = cmd.Stdout

	scannerError := make(chan error)

	scanner := bufio.NewScanner(r)

	go func() {
		for scanner.Scan() {
			line := scanner.Text()

			if strings.Contains(line, ErrorStart) {
				scannerError <- errors.New(strings.Title(strings.TrimPrefix(line, ErrorStart)))
			}

			log.Println(line)
		}

		scannerError <- nil
	}()

	err := cmd.Start()

	execError := <-scannerError

	log.Println(execError)
	if err != nil {
		log.Println(err)
		if execError != nil {
			return execError
		}
		return err
	}

	err = cmd.Wait()
	if err != nil {
		log.Println(err)
		if execError != nil {
			return execError
		}
		return err
	}

	return execError
}
