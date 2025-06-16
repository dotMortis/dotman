package metafile

type Packages struct {
	fileHandler *TomlFileHandler[PackagesContent]
}

func (p *Packages) init() error {
	return p.fileHandler.Read()
}

func (p *Packages) ToSaved(pkg string) {
	p.Content().Saved.Add(pkg)
	p.Content().Ignored.Remove(pkg)
}

func (p *Packages) ToSavedIndex(pkg string, index int) {
	p.Content().Saved.ToIndex(pkg, index)
}

func (p *Packages) ToIgnored(pkg string) {
	p.Content().Ignored.Add(pkg)
	p.Content().Saved.Remove(pkg)
}

func (p *Packages) Content() *PackagesContent {
	return p.fileHandler.Content
}

func (p *Packages) Save() error {
	return p.fileHandler.Write()
}

func NewPackages(path string) (*Packages, error) {
	p := &Packages{
		fileHandler: NewTomlFileHandler(path, NewPackagesContent()),
	}
	return p, p.init()
}
