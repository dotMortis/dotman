package pacman

import (
	"dotman/bashcmd"
	"dotman/metafile/pacman"
	"fmt"
)

type PacmanManager struct {
	packagesMetafile *pacman.PacmanPackages
	bashCmd          *bashcmd.BashCmd
	Packages         *Packages
}

func (pm *PacmanManager) Init() error {
	return nil
}

func NewPacmanManager(path string, bashCmd *bashcmd.BashCmd) (*PacmanManager, error) {
	packagesMetafile, err := pacman.NewPacmanPackages(path)
	if err != nil {
		return nil, err
	}
	packages, err := NewPackages(packagesMetafile, bashCmd)
	if err != nil {
		return nil, fmt.Errorf("failed to load packages: %w", err)
	}
	return &PacmanManager{
		packagesMetafile: packagesMetafile,
		bashCmd:          bashCmd,
		Packages:         packages,
	}, nil
}
