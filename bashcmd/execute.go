package bashcmd

import (
	"bufio"
	"dotman/bashcmd/writer"
	"fmt"
	"os"
	"os/exec"
)

type BashCmd struct {
	writer writer.BashCmdWriter
}

func (b *BashCmd) Execute(command string, args ...string) error {
	cmd := exec.Command(command, args...)

	cmd.Stdin = os.Stdin

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to create stdout pipe: %v", err)
	}
	defer stdout.Close()

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("failed to create stderr pipe: %v", err)
	}
	defer stderr.Close()

	done := make(chan error)
	defer close(done)

	go func() {
		scanner := bufio.NewScanner(stderr)
		scanner.Split(bufio.ScanBytes)
		for scanner.Scan() {
			bytes := scanner.Bytes()
			b.writer.WriteErr(&bytes)
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stdout)
		scanner.Split(bufio.ScanBytes)
		for scanner.Scan() {
			bytes := scanner.Bytes()
			b.writer.WriteStd(&bytes)
		}
	}()

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start command: %v", err)
	}

	go func() {
		done <- cmd.Wait()
	}()

	return <-done
}

func (b *BashCmd) Close() error {
	return b.writer.Close()
}

func NewBashCmd(writer writer.BashCmdWriter) *BashCmd {
	return &BashCmd{
		writer: writer,
	}
}
