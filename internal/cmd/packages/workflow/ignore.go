package workflow

import (
	"dotman/internal/manager"
	"dotman/internal/ui"
	"fmt"
	"log"
)

func Ignored(pm manager.Manager, action IgnoreAction) {
	packages := pm.Packages().Ignored()

	if len(*packages) == 0 && action != IgnoreActionList {
		fmt.Println("No ignored packages found UwU")
		return
	}

	if action == IgnoreActionList {
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
