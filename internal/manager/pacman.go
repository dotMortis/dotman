package manager

import (
	"dotman/internal/bashcmd"
	"dotman/internal/metafile"
	"dotman/internal/packages"
	"fmt"
)

type PacmanManager struct {
	packagesMetafile *metafile.Packages
	bashCmd          *bashcmd.BashCmd
	packages         *Packages
}

func (pm *PacmanManager) Packages() *Packages {
	return pm.packages
}

func NewPacmanManager(metafilePath string, bashCmd *bashcmd.BashCmd) (*PacmanManager, error) {
	packagesMetafile, err := metafile.NewPackages(metafilePath)
	if err != nil {
		return nil, fmt.Errorf("[PacmanManager] failed to create pacman packages metafile:\n%w", err)
	}
	packages, err := NewPackages(packagesMetafile, packages.NewPacmanCommands(bashCmd))
	if err != nil {
		return nil, fmt.Errorf("[PacmanManager] failed to load packages:\n%w", err)
	}
	return &PacmanManager{
		packages:         packages,
		packagesMetafile: packagesMetafile,
		bashCmd:          bashCmd,
	}, nil
}
