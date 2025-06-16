package workflow

import (
	"dotman/internal/manager"
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
)

func Ignore(pm *manager.PacmanManager) {
	ignored := pm.Packages.Ignored()

	if len(*ignored) == 0 {
		fmt.Println("No ignored packages found UwU")
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
					huh.NewOption("move to saved", "saved"),
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
	case "saved":
		for _, pkg := range *selected {
			pm.Packages.ToSaved(pkg)
		}
		pm.Packages.SaveMetafile()
	case "remove":
		for _, pkg := range *selected {
			pm.Packages.RemoveFromMetafile(pkg)
		}
		pm.Packages.SaveMetafile()
	}
}
