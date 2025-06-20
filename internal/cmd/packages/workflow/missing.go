package workflow

import (
	"dotman/internal/manager"
	"dotman/internal/ui"
	"fmt"
	"log"
)

func Missing(pm manager.Manager, action MissingAction) {
	packages, err := pm.Packages().Uninstalled()
	if err != nil {
		log.Fatal(err)
		return
	}
	if len(*packages) == 0 && action != MissingActionList {
		fmt.Println("No missing packages found UwU")
		return
	}
	if action == MissingActionList {
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
	RunSliceAction(string(action), pm, selected)
}
