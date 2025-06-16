package workflow

import (
	"dotman/internal/manager"
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
)

func Saved(pm *manager.PacmanManager) {
	ignored := pm.Packages.Saved()

	if len(*ignored) == 0 {
		fmt.Println("No saved packages found UwU")
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
					huh.NewOption("move to ignore", "ignore"),
					huh.NewOption("remove", "remove"),
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
	case "remove":
		for _, pkg := range *selected {
			pm.Packages.RemoveFromMetafile(pkg)
		}
		pm.Packages.SaveMetafile()
	}
}
