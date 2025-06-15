package pacman

import "dotman/metafile/pacman"

type PacmanManager struct {
	packagesMetafile *pacman.PacmanPackages
}

func NewPacmanManager(path string) (*PacmanManager, error) {
	packagesMetafile, err := pacman.NewPacmanPackages(path)
	if err != nil {
		return nil, err
	}
	return &PacmanManager{
		packagesMetafile: packagesMetafile,
	}, nil
}
