package pacman

import (
	"dotman/internal/bashcmd"
	meta "dotman/internal/metafile/pacman"
	"dotman/internal/pacman"
	"fmt"
	"slices"
	"strings"
)

type Packages struct {
	metafile *meta.PacmanPackages
	bashCmd  *bashcmd.BashCmd
}

func (pks *Packages) Saved() *pacman.Packages {
	return pks.metafile.Content().Saved
}

func (pks *Packages) Ignored() *pacman.Packages {
	return pks.metafile.Content().Ignored
}

func (pks *Packages) Installed(filterIgnored bool) (*pacman.Packages, error) {
	rawResult, err := pks.bashCmd.ExecuteOutout("pacman", "-Qqen")
	if err != nil {
		return nil, fmt.Errorf("failed to get installed packages: %w", err)
	}
	splitted := strings.Split(rawResult, "\n")
	installed := (&pacman.Packages{}).Add(splitted...)
	if filterIgnored {
		ignored := pks.Ignored()
		for _, pkg := range *ignored {
			installed.Remove(pkg)
		}
	}
	return installed, nil
}

func (pks *Packages) Surplus(filterIgnored bool) *pacman.Packages {
	installed, err := pks.Installed(filterIgnored)
	if err != nil {
		return nil
	}
	saved := pks.Saved()
	for _, pkg := range *saved {
		installed.Remove(pkg)

	}
	return (&pacman.Packages{}).Add(*installed...)
}

func (pks *Packages) Uninstalled() *pacman.Packages {
	installed, err := pks.Installed(true)
	if err != nil {
		return nil
	}
	saved := pks.Saved()
	uninstalled := &pacman.Packages{}
	for _, pkg := range *saved {
		if !slices.Contains(*installed, pkg) {
			uninstalled.Add(pkg)
		}
	}
	return uninstalled
}

func (pks *Packages) ToSaved(pkg string) error {
	isPackage, err := pks.IsPackage(pkg)
	if err != nil {
		return fmt.Errorf("failed to check if package is installed: %w", err)
	}
	if !isPackage {
		return fmt.Errorf("'%s' is not a valid package", pkg)
	}
	pks.metafile.ToSaved(pkg)
	return nil
}

func (pks *Packages) ToIgnored(pkg string) error {
	isPackage, err := pks.IsPackage(pkg)
	if err != nil {
		return fmt.Errorf("failed to check if package is installed: %w", err)
	}
	if !isPackage {
		return fmt.Errorf("'%s' is not a valid package", pkg)
	}
	pks.metafile.ToIgnored(pkg)
	return nil
}

func (pks *Packages) RemoveFromMetafile(pkg string) bool {
	content := pks.metafile.Content()
	res1 := content.Ignored.Remove(pkg)
	res2 := content.Saved.Remove(pkg)
	return res1 || res2
}

func (pks *Packages) IsPackage(pkg string) (bool, error) {
	result, err := pks.bashCmd.ExecuteOutout("pacman", "-Ss", fmt.Sprintf("^%s$", pkg))
	if err != nil {
		return false, fmt.Errorf("failed to check if package is installed: %w", err)
	}
	return strings.Contains(result, fmt.Sprintf("extra/%s ", pkg)), nil
}

func (pks *Packages) InstallMissing() (installedPackages *pacman.Packages, error error) {
	var result = &pacman.Packages{}
	for _, pkg := range *pks.Uninstalled() {
		err := pks.bashCmd.Execute("sudo", "pacman", "-S", pkg)
		if err != nil {
			return result, fmt.Errorf("failed to install package: %w", err)
		}
		result.Add(pkg)
	}
	return result, nil
}

func (pks *Packages) UninstallSurplus() (removePackages *pacman.Packages, error error) {
	var result = &pacman.Packages{}
	for _, pkg := range *pks.Surplus(true) {
		err := pks.bashCmd.Execute("sudo", "pacman", "-Rs", pkg)
		if err != nil {
			return result, fmt.Errorf("failed to uninstall package: %w", err)
		}
		result.Add(pkg)
	}
	return result, nil
}

func (pks *Packages) SaveMetafile() error {
	return pks.metafile.Save()
}

func NewPackages(metafile *meta.PacmanPackages, bashCmd *bashcmd.BashCmd) (*Packages, error) {
	packages := &Packages{
		metafile: metafile,
		bashCmd:  bashCmd,
	}
	return packages, nil
}
