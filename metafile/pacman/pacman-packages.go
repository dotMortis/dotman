package pacman

import "dotman/metafile"

type PacmanPackages struct {
	fileHandler *metafile.TomlFileHandler[PacmanPackagesContent]
}

func (ppm *PacmanPackages) init() error {
	return ppm.fileHandler.Read()
}

func (ppm *PacmanPackages) Content() *PacmanPackagesContent {
	return ppm.fileHandler.Content
}

func (ppm *PacmanPackages) Save() error {
	return ppm.fileHandler.Write()
}

func NewPacmanPackages(path string) (*PacmanPackages, error) {
	ppm := &PacmanPackages{
		fileHandler: metafile.NewTomlFileHandler(path, new(PacmanPackagesContent)),
	}
	return ppm, ppm.init()
}
