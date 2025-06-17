package workflow

import (
	"dotman/internal/manager"
	"dotman/internal/ui"
	"fmt"
	"log"
	"strconv"

	"github.com/charmbracelet/huh"
)

func Saved(pm manager.Manager, action SaveAction) {
	if action == SaveActionReorder {
		reorderSaved(pm)
		return
	}
	packages := pm.Packages().Saved()

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
	RunSliceAction(string(action), pm, selected)

}

func reorderSaved(pm manager.Manager) {
	for {
		packages := pm.Packages().Saved()

		if len(*packages) == 0 {
			fmt.Println("No saved packages found UwU")
			return
		}

		var (
			selected string
			position string
			length   = len(*packages)
		)
		options := ui.NewPackagesSelectOptions(packages, true)
		form := ui.NewSingleGroupForm(
			ui.NewSelectPackages(&selected, *options...),
			huh.NewInput().
				Title(fmt.Sprintf("New Position [1-%v], or -1 to quit", length)).
				Validate(func(value string) error {
					position, err := strconv.Atoi(position)
					if err != nil {
						return fmt.Errorf("position must be a number")
					}
					if position > length || (position < 1 && position != -1) {
						return fmt.Errorf("position must be between 1 and %v, or -1 to quit", length)
					}
					return nil

				}).
				Value(&position),
		)
		if err := form.Run(); err != nil {
			log.Fatal(err)
			return
		}
		positionInt, _ := strconv.Atoi(position)
		if positionInt == -1 {
			return
		}
		if err := RunReorderAction(pm, selected, positionInt-1); err != nil {
			log.Fatal(err)
			return
		}
	}
}
