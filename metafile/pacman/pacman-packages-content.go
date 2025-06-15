package pacman

import "slices"

type Packages []string

func (pkgs *Packages) Add(packages ...string) {
	for _, pkg := range packages {
		if !slices.Contains(*pkgs, pkg) {
			*pkgs = append(*pkgs, pkg)
		}
	}
}

func (pkgs *Packages) Remove(pkg string) bool {
	index := slices.Index(*pkgs, pkg)
	if index == -1 {
		return false
	}
	*pkgs = slices.Delete(*pkgs, index, index+1)
	return true
}

func (pkgs *Packages) ToIndex(pkg string, index int) {
	if pkgs.Remove(pkg) {
		*pkgs = slices.Insert(*pkgs, index, pkg)
	}
}

type PacmanPackagesContent struct {
	Packages Packages
}
