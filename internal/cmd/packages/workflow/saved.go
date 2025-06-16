package workflow

import (
	"dotman/internal/manager"
	"dotman/internal/ui"
	"fmt"
	"log"
)

func Saved(pm *manager.PacmanManager, action SaveAction) {
	packages := pm.Packages.Saved()

	if len(*packages) == 0 && action != SaveActionList {
		fmt.Println("No saved packages found UwU")
		return
	}

	if action == SaveActionList {
		ui.PrintPackages(packages)
		return
	}

	var selected = new([]string)
	options := ui.NewPackagesSelectOptions(packages, false)
	form := ui.NewSingleGroupForm(
		ui.NewMultiSelectPackages(selected, *options...),
	)
	if err := form.Run(); err != nil {
		log.Fatal(err)
		return
	}
	RunAction(string(action), pm, selected)
}
