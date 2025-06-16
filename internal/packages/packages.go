package packages

import (
	"fmt"
	"slices"
)

type Packages []string

func (pkgs *Packages) Add(packages ...string) *Packages {
	for _, pkg := range packages {
		if !slices.Contains(*pkgs, pkg) {
			*pkgs = append(*pkgs, pkg)
		}
	}
	return pkgs
}

func (pkgs *Packages) Remove(pkg string) bool {
	index := slices.Index(*pkgs, pkg)
	if index == -1 {
		return false
	}
	*pkgs = slices.Delete(*pkgs, index, index+1)
	return true
}

func (pkgs *Packages) ToIndex(pkg string, index int) *Packages {
	if pkgs.Remove(pkg) {
		*pkgs = slices.Insert(*pkgs, index, pkg)
	}
	return pkgs
}

func (pkgs *Packages) String() string {
	return fmt.Sprintf("%v", *pkgs)
}
