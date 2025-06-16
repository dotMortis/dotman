package manager

import (
	"dotman/internal/metafile"
	"dotman/internal/packages"
	"fmt"
	"slices"
)

type Packages struct {
	metafile *metafile.Packages
	commands packages.Commands
}

func (pks *Packages) Saved() *packages.Packages {
	return pks.metafile.Content().Saved
}

func (pks *Packages) Ignored() *packages.Packages {
	return pks.metafile.Content().Ignored
}

func (pks *Packages) Installed(filterIgnored bool) (*packages.Packages, error) {
	installed, err := pks.commands.Installed()
	if err != nil {
		return nil, fmt.Errorf("failed to get installed packages: %w", err)
	}
	if filterIgnored {
		ignored := pks.Ignored()
		for _, pkg := range *ignored {
			installed.Remove(pkg)
		}
	}
	return installed, nil
}

func (pks *Packages) Surplus(filterIgnored bool) *packages.Packages {
	installed, err := pks.Installed(filterIgnored)
	if err != nil {
		return nil
	}
	saved := pks.Saved()
	for _, pkg := range *saved {
		installed.Remove(pkg)

	}
	return (&packages.Packages{}).Add(*installed...)
}

func (pks *Packages) Uninstalled() *packages.Packages {
	installed, err := pks.Installed(true)
	if err != nil {
		return nil
	}
	saved := pks.Saved()
	uninstalled := &packages.Packages{}
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

func (pks *Packages) ToSavedIndex(pkg string, index int) error {
	saved := pks.Saved()
	if !slices.Contains(*saved, pkg) {
		return fmt.Errorf("'%s' is not in the list of saved packages", pkg)
	}
	pks.metafile.ToSavedIndex(pkg, index)
	return nil
}

func (pks *Packages) ToIgnored(pkg string, force bool) error {
	if !force {
		isPackage, err := pks.IsPackage(pkg)
		if err != nil {
			return fmt.Errorf("failed to check if package is installed: %w", err)
		}
		if !isPackage {
			return fmt.Errorf("'%s' is not a valid package", pkg)
		}
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
	return pks.commands.FindPackage(pkg)
}

func (pks *Packages) InstallMissing(packagesToInstall *[]string, noConfirm bool) (installedPackages *packages.Packages, error error) {
	var result = &packages.Packages{}
	uninstalled := pks.Uninstalled()
	if len(*packagesToInstall) >= 0 {
		for _, pkg := range *packagesToInstall {
			if !slices.Contains(*uninstalled, pkg) {
				return result, fmt.Errorf("'%s' is not in the list of available packages", pkg)
			}
		}
		uninstalled = (&packages.Packages{}).Add(*packagesToInstall...)
	}

	for _, pkg := range *uninstalled {
		if err := pks.commands.Install(pkg, noConfirm); err != nil {
			return result, fmt.Errorf("failed to install package: %w", err)
		}
		result.Add(pkg)
	}
	return result, nil
}

func (pks *Packages) UninstallSurplus() (removedPackages *packages.Packages, error error) {
	var result = &packages.Packages{}
	for _, pkg := range *pks.Surplus(true) {
		if err := pks.commands.Uninstall(pkg); err != nil {
			return result, fmt.Errorf("failed to uninstall package: %w", err)
		}
		result.Add(pkg)
	}
	return result, nil
}

func (pks *Packages) SaveMetafile() error {
	return pks.metafile.Save()
}

func NewPackages(metafile *metafile.Packages, commands packages.Commands) (*Packages, error) {
	packages := &Packages{
		metafile: metafile,
		commands: commands,
	}
	return packages, nil
}
