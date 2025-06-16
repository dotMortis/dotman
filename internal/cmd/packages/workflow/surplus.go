package workflow

import (
	"dotman/internal/manager"
	"dotman/internal/ui"
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
)

func Surplus(pm *manager.PacmanManager, action SurplusAction) {
	packages := pm.Packages.Surplus(true)

	if len(*packages) == 0 && action != SurplusActionList {
		fmt.Println("No surplus packages found UwU")
		return
	}

	if action == SurplusActionList {
		ui.PrintPackages(packages)
		return
	}

	options := make([]huh.Option[string], len(*packages))
	for i, pkg := range *packages {
		options[i] = huh.NewOption(pkg, pkg)
	}

	var selected = new([]string)
	form := ui.NewSingleGroupForm(
		ui.NewMultiSelectPackages(selected, options...),
	)
	if err := form.Run(); err != nil {
		log.Fatal(err)
		return
	}
	RunAction(string(action), pm, selected)
}
