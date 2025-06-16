package manager

import (
	"dotman/internal/bashcmd"
	"dotman/internal/metafile"
	"dotman/internal/packages"
	"fmt"
)

type YayManager struct {
	packagesMetafile *metafile.Packages
	bashCmd          *bashcmd.BashCmd
	Packages         *Packages
}

func NewYayManager(metafilePath string, bashCmd *bashcmd.BashCmd) (*YayManager, error) {
	packagesMetafile, err := metafile.NewPackages(metafilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create pacman packages metafile: %w", err)
	}
	packages, err := NewPackages(packagesMetafile, packages.NewYayCommands(bashCmd))
	if err != nil {
		return nil, fmt.Errorf("failed to load packages: %w", err)
	}
	return &YayManager{
		packagesMetafile: packagesMetafile,
		bashCmd:          bashCmd,
		Packages:         packages,
	}, nil
}
