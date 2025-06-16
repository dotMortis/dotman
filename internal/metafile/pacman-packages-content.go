package metafile

import "dotman/internal/pacman"

type PacmanPackagesContent struct {
	Saved   *pacman.Packages
	Ignored *pacman.Packages
}

func NewPacmanPackagesContent() *PacmanPackagesContent {
	return &PacmanPackagesContent{
		Saved:   &pacman.Packages{},
		Ignored: &pacman.Packages{},
	}
}
