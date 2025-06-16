package metafile

type Packages interface {
	ToSaved(pkg string)
	ToSavedIndex(pkg string, index int)
	ToIgnored(pkg string)
	Content() *PackagesContent
	Save() error
}
