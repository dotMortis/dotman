package packages

type Commands interface {
	Installed() (*Packages, error)
	FindPackage(pkg string) (bool, error)
	Install(pkg string, noConfirm bool) error
	Uninstall(pkg string) error
}
