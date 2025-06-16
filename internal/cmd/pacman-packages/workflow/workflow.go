package workflow

import (
	"dotman/internal/manager"
	"fmt"
)

type SaveAction string
type IgnoreAction string
type MissingAction string
type SurplusAction string

const (
	SaveActionIgnore  SaveAction = "ignore"
	SaveActionRemove  SaveAction = "remove"
	SaveActionList    SaveAction = "list"
	SaveActionReorder SaveAction = "reorder"

	IgnoreActionSave   IgnoreAction = "save"
	IgnoreActionRemove IgnoreAction = "remove"
	IgnoreActionList   IgnoreAction = "list"

	MissingActionInstall      MissingAction = "install"
	MissingActionForceInstall MissingAction = "force-install"
	MissingActionIgnore       MissingAction = "ignore"
	MissingActionRemove       MissingAction = "remove"
	MissingActionList         MissingAction = "list"

	SurplusActionSave        SurplusAction = "save"
	SurplusActionIgnore      SurplusAction = "ignore"
	SurplusActionForceIgnore SurplusAction = "force-ignore"
	SurplusActionList        SurplusAction = "list"
)

func RunSliceAction(action string, pm *manager.PacmanManager, selected *[]string) {
	switch action {
	case "install":
		installed, err := pm.Packages.InstallMissing(selected, false)
		fmt.Println(installed)
		if err != nil {
			fmt.Println("Failed to install packages:", err)
		}
	case "force-install":
		installed, err := pm.Packages.InstallMissing(selected, true)
		fmt.Println(installed)
		if err != nil {
			fmt.Println("Failed to install packages:", err)
		}
	case "ignore":
		for _, pkg := range *selected {
			if err := pm.Packages.ToIgnored(pkg, false); err != nil {
				fmt.Println("Failed to ignore package:", err)
			}
		}
		pm.Packages.SaveMetafile()
	case "force-ignore":
		for _, pkg := range *selected {
			if err := pm.Packages.ToIgnored(pkg, true); err != nil {
				fmt.Println("Failed to ignore package:", err)
			}
		}
		pm.Packages.SaveMetafile()
	case "remove":
		for _, pkg := range *selected {
			pm.Packages.RemoveFromMetafile(pkg)
		}
		pm.Packages.SaveMetafile()
	case "save":
		for _, pkg := range *selected {
			if err := pm.Packages.ToSaved(pkg); err != nil {
				fmt.Println("Failed to save package:", err)
			}
		}
		pm.Packages.SaveMetafile()
	}
}

func RunReorderAction(pm *manager.PacmanManager, selected string, index int) error {
	if err := pm.Packages.ToSavedIndex(selected, index); err != nil {
		return fmt.Errorf("failed to reorder package: %w", err)
	}
	pm.Packages.SaveMetafile()
	return nil
}
