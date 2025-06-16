package manager

import (
	"dotman/internal/bashcmd"
	"dotman/internal/metafile"
	"dotman/internal/packages"
	"fmt"
)

type PacmanManager struct {
	packagesMetafile *metafile.PacmanPackages
	bashCmd          *bashcmd.BashCmd
	Packages         *Packages
}

func NewPacmanManager(path string, bashCmd *bashcmd.BashCmd) (*PacmanManager, error) {
	packagesMetafile, err := metafile.NewPacmanPackages(path)
	if err != nil {
		return nil, fmt.Errorf("failed to create pacman packages metafile: %w", err)
	}
	packages, err := NewPackages(packagesMetafile, packages.NewPacmanCommands(bashCmd))
	if err != nil {
		return nil, fmt.Errorf("failed to load packages: %w", err)
	}
	return &PacmanManager{
		packagesMetafile: packagesMetafile,
		bashCmd:          bashCmd,
		Packages:         packages,
	}, nil
}
