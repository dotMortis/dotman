package packages

import (
	"dotman/internal/bashcmd"
	"fmt"
	"strings"
)

type PacmanCommands struct {
	bashCmd *bashcmd.BashCmd
}

func (c *PacmanCommands) Installed() (*Packages, error) {
	rawResult, err := c.bashCmd.ExecuteOutout("pacman", "-Qqen")
	if err != nil {
		return nil, fmt.Errorf("[PacmanCommands] failed to get installed packages:\n%w", err)
	}
	splitted := strings.Split(rawResult, "\n")
	if splitted[len(splitted)-1] == "" {
		splitted = splitted[:len(splitted)-1]
	}
	return (&Packages{}).Add(splitted...), nil
}

func (c *PacmanCommands) FindPackage(pkg string) (bool, error) {
	result, err := c.bashCmd.ExecuteOutout("pacman", "-Ss", fmt.Sprintf("^%s$", pkg))
	if err != nil {
		return false, fmt.Errorf("[PacmanCommands] failed to check if package is installed:\n%w", err)
	}
	return strings.Contains(result, fmt.Sprintf("extra/%s ", pkg)), nil
}

func (c *PacmanCommands) Install(pkg string, noConfirm bool) error {
	var flags = []string{"pacman", "-S"}
	if noConfirm {
		flags = append(flags, "--noconfirm")
	}
	flags = append(flags, pkg)
	if err := c.bashCmd.Execute("sudo", flags...); err != nil {
		return fmt.Errorf("[PacmanCommands] failed to install package:\n%w", err)
	}
	if err := c.bashCmd.Execute("sudo", "pacman", "-D", "--asexplicit", pkg); err != nil {
		return fmt.Errorf("[PacmanCommands] failed to add package to explicit dependencies:\n%w", err)
	}
	return nil
}

func (c *PacmanCommands) Uninstall(pkg string) error {
	return c.bashCmd.Execute("sudo", "pacman", "-Rs", pkg)
}

func NewPacmanCommands(bashCmd *bashcmd.BashCmd) *PacmanCommands {
	return &PacmanCommands{
		bashCmd: bashCmd,
	}
}
