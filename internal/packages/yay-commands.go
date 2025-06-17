package packages

import (
	"dotman/internal/bashcmd"
	"fmt"
	"strings"
)

type YayCommands struct {
	bashCmd *bashcmd.BashCmd
}

func (c *YayCommands) Installed() (*Packages, error) {
	rawResult, err := c.bashCmd.ExecuteOutout("yay", "-Qqen")
	if err != nil {
		return nil, fmt.Errorf("[YayCommands] failed to get installed packages:\n%w", err)
	}
	splitted := strings.Split(rawResult, "\n")
	if splitted[len(splitted)-1] == "" {
		splitted = splitted[:len(splitted)-1]
	}
	return (&Packages{}).Add(splitted...), nil
}

func (c *YayCommands) FindPackage(pkg string) (bool, error) {
	result, err := c.bashCmd.ExecuteOutout("yay", "-Ss", fmt.Sprintf("^%s$", pkg))
	if err != nil {
		return false, fmt.Errorf("[YayCommands] failed to check if package is installed:\n%w", err)
	}
	return strings.Contains(result, fmt.Sprintf("extra/%s ", pkg)), nil
}

func (c *YayCommands) Install(pkg string, noConfirm bool) error {
	var flags = []string{"yay", "-S"}
	if noConfirm {
		flags = append(flags, "--noconfirm")
	}
	flags = append(flags, pkg)
	if err := c.bashCmd.Execute("sudo", flags...); err != nil {
		return fmt.Errorf("[YayCommands] failed to install package:\n%w", err)
	}
	if err := c.bashCmd.Execute("sudo", "pacman", "-D", "--asexplicit", pkg); err != nil {
		return fmt.Errorf("[YayCommands] failed to add package to explicit dependencies:\n%w", err)
	}
	return nil
}

func (c *YayCommands) Uninstall(pkg string) error {
	return c.bashCmd.Execute("sudo", "yay", "-Rs", pkg)
}

func NewYayCommands(bashCmd *bashcmd.BashCmd) *YayCommands {
	return &YayCommands{
		bashCmd: bashCmd,
	}
}
