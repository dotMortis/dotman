package pacman

import (
	"dotman/internal/bashcmd"
	meta "dotman/internal/metafile/pacman"
	"fmt"
)

type PacmanManager struct {
	packagesMetafile *meta.PacmanPackages
	bashCmd          *bashcmd.BashCmd
	Packages         *Packages
}

func NewPacmanManager(path string, bashCmd *bashcmd.BashCmd) (*PacmanManager, error) {
	packagesMetafile, err := meta.NewPacmanPackages(path)
	if err != nil {
		return nil, fmt.Errorf("failed to create pacman packages metafile: %w", err)
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
