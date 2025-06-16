package metafile

import "dotman/internal/packages"

type PackagesContent struct {
	Saved   *packages.Packages
	Ignored *packages.Packages
}

func NewPackagesContent() *PackagesContent {
	return &PackagesContent{
		Saved:   &packages.Packages{},
		Ignored: &packages.Packages{},
	}
}
