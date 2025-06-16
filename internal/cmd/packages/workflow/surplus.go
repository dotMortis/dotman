package workflow

import (
	"dotman/internal/manager"
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
)

func Surplus(pm *manager.PacmanManager) {
	ignored := pm.Packages.Surplus(true)

	if len(*ignored) == 0 {
		fmt.Println("No surplus packages found UwU")
		return
	}

	options := make([]huh.Option[string], len(*ignored))
	for i, pkg := range *ignored {
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
					huh.NewOption("save", "save"),
					huh.NewOption("ignore", "ignore"),
					huh.NewOption("let me out", "none"),
				).
				Title("Select action").
				Value(&action),
		),
	)
	if err := form.Run(); err != nil {
		log.Fatal(err)
		return
	}

	switch action {
	case "ignore":
		for _, pkg := range *selected {
			pm.Packages.ToIgnored(pkg)
		}
		pm.Packages.SaveMetafile()
	case "save":
		for _, pkg := range *selected {
			pm.Packages.ToSaved(pkg)
		}
		pm.Packages.SaveMetafile()
	}
}
