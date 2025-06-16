package metafile

type PacmanPackages struct {
	fileHandler *TomlFileHandler[PacmanPackagesContent]
}

func (ppm *PacmanPackages) init() error {
	return ppm.fileHandler.Read()
}

func (ppm *PacmanPackages) ToSaved(pkg string) {
	ppm.Content().Saved.Add(pkg)
	ppm.Content().Ignored.Remove(pkg)
}

func (ppm *PacmanPackages) ToSavedIndex(pkg string, index int) {
	ppm.Content().Saved.ToIndex(pkg, index)
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
		fileHandler: NewTomlFileHandler(path, NewPacmanPackagesContent()),
	}
	return ppm, ppm.init()
}
