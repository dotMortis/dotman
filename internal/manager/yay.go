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
	packages         *Packages
}

func (pm *YayManager) Packages() *Packages {
	return pm.packages
}

func NewYayManager(metafilePath string, bashCmd *bashcmd.BashCmd) (*YayManager, error) {
	packagesMetafile, err := metafile.NewPackages(metafilePath)
	if err != nil {
		return nil, fmt.Errorf("[YayManager] failed to create pacman packages metafile:\n%w", err)
	}
	packages, err := NewPackages(packagesMetafile, packages.NewYayCommands(bashCmd))
	if err != nil {
		return nil, fmt.Errorf("[YayManager] failed to load packages:\n%w", err)
	}
	return &YayManager{
		packagesMetafile: packagesMetafile,
		bashCmd:          bashCmd,
		packages:         packages,
	}, nil
}
