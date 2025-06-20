package workflow

import (
	"dotman/internal/manager"
	"dotman/internal/ui"
	"fmt"
	"log"
)

func Surplus(pm manager.Manager, action SurplusAction) {
	packages, err := pm.Packages().Surplus(true)
	if err != nil {
		log.Fatal(err)
		return
	}
	if len(*packages) == 0 && action != SurplusActionList {
		fmt.Println("No surplus packages found UwU")
		return
	}
	if action == SurplusActionList {
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
