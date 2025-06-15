package pacman

import (
	"dotman/bashcmd"
	meta "dotman/metafile/pacman"
	"dotman/pacman"
	"fmt"
	"slices"
	"strings"
)

type Packages struct {
	Uninstalled      *pacman.Packages
	Surplus          *pacman.Packages
	packagesMetafile *meta.PacmanPackages
	bashCmd          *bashcmd.BashCmd
}

func (pks *Packages) Saved() *pacman.Packages {
	return pks.packagesMetafile.Content().Packages
}

func (pks *Packages) Installed() (*pacman.Packages, error) {
	rawResult, err := pks.bashCmd.ExecuteOutout("pacman", "-Qqen")
	if err != nil {
		return nil, fmt.Errorf("failed to get installed packages: %w", err)
	}
	splitted := strings.Split(rawResult, "\n")
	installed := &pacman.Packages{}
	return installed.Add(splitted...), nil
}

func (pks *Packages) Reload() error {
	installed, err := pks.Installed()
	if err != nil {
		return fmt.Errorf("failed to get installed packages: %w", err)
	}
	saved := pks.Saved()
	pks.Surplus = getSurplus(installed, saved)
	pks.Uninstalled = getUninstalled(installed, saved)
	return nil
}

func NewPackages(metafile *meta.PacmanPackages, bashCmd *bashcmd.BashCmd) (*Packages, error) {
	packages := &Packages{}
	if err := packages.Reload(); err != nil {
		return nil, fmt.Errorf("failed to load packages: %w", err)
	}
	return packages, nil
}

func getSurplus(installed *pacman.Packages, saved *pacman.Packages) *pacman.Packages {
	surplus := &pacman.Packages{}
	for _, pkg := range *installed {
		if !slices.Contains(*saved, pkg) {
			surplus.Add(pkg)
		}
	}
	return surplus
}

func getUninstalled(installed *pacman.Packages, saved *pacman.Packages) *pacman.Packages {
	uninstalled := &pacman.Packages{}
	for _, pkg := range *saved {
		if !slices.Contains(*installed, pkg) {
			uninstalled.Add(pkg)
		}
	}
	return uninstalled
}
