package workflow

import (
	"dotman/internal/manager"
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
)

func Missing(pm *manager.PacmanManager) {
	uninstalled := pm.Packages.Uninstalled()

	if len(*uninstalled) == 0 {
		fmt.Println("No missing packages found UwU")
		return
	}

	options := make([]huh.Option[string], len(*uninstalled))
	for i, pkg := range *uninstalled {
		options[i] = huh.NewOption(pkg, pkg)
	}

	var (
		selected = new([]string)
		action   = "none"
	)
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Options(options...).
				Title("Select packages").
				Value(selected).
				Filterable(true),
			huh.NewSelect[string]().
				Options(
					huh.NewOption("install", "install"),
					huh.NewOption("ignore", "ignore"),
					huh.NewOption("let me out", "none"),
				).
				Title("Select action").
				Value(&action),
		),
	).WithTheme(huh.ThemeCatppuccin())

	if err := form.Run(); err != nil {
		log.Fatal(err)
		return
	}

	switch action {
	case "install":
		installed, err := pm.Packages.InstallMissing(selected)
		fmt.Println(installed)
		if err != nil {
			log.Fatal(err)
			return
		}
	case "ignore":
		for _, pkg := range *selected {
			pm.Packages.ToIgnored(pkg)
		}
		pm.Packages.SaveMetafile()
	}
}
