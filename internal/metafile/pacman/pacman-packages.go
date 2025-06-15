package pacman

import "dotman/internal/metafile"

type PacmanPackages struct {
	fileHandler *metafile.TomlFileHandler[PacmanPackagesContent]
}

func (ppm *PacmanPackages) init() error {
	return ppm.fileHandler.Read()
}

func (ppm *PacmanPackages) ToSaved(pkg string) {
	ppm.Content().Saved.Add(pkg)
	ppm.Content().Ignored.Remove(pkg)
}

func (ppm *PacmanPackages) ToIgnored(pkg string) {
	ppm.Content().Ignored.Add(pkg)
	ppm.Content().Saved.Remove(pkg)
}

func (ppm *PacmanPackages) Content() *PacmanPackagesContent {
	return ppm.fileHandler.Content
}

func (ppm *PacmanPackages) Save() error {
	return ppm.fileHandler.Write()
}

func NewPacmanPackages(path string) (*PacmanPackages, error) {
	ppm := &PacmanPackages{
		fileHandler: metafile.NewTomlFileHandler(path, NewPacmanPackagesContent()),
	}
	return ppm, ppm.init()
}
