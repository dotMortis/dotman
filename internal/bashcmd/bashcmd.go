package bashcmd

import (
	"fmt"
	"os"
	"os/exec"
)

type BashCmd struct {
	writer *IOWriter
}

func (b *BashCmd) Execute(command string, args ...string) error {
	cmd := exec.Command(command, args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = b.writer
	cmd.Stderr = b.writer

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("[BashCmd] failed to start command:\n%w", err)
	}
	return nil
}

func (b *BashCmd) ExecuteOutout(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("[BashCmd] failed to execute command:\n%w", err)
	}
	return string(output), nil
}

func NewBashCmd(writer *IOWriter) *BashCmd {
	return &BashCmd{writer: writer}
}
