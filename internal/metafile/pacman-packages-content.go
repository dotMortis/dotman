package metafile

import "dotman/internal/packages"

type PacmanPackagesContent struct {
	Saved   *packages.Packages
	Ignored *packages.Packages
}

func NewPacmanPackagesContent() *PacmanPackagesContent {
	return &PacmanPackagesContent{
		Saved:   &packages.Packages{},
		Ignored: &packages.Packages{},
	}
}
