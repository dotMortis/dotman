package ui

import (
	"dotman/internal/pacman"
	"fmt"

	"github.com/charmbracelet/huh"
)

func NewSingleGroupForm(fields ...huh.Field) *huh.Form {
	return huh.NewForm(huh.NewGroup(fields...)).WithTheme(huh.ThemeCatppuccin())
}

func NewPackagesSelectOptions(packages *pacman.Packages, numbered bool) *[]huh.Option[string] {
	options := make([]huh.Option[string], len(*packages))
	for i, pkg := range *packages {
		options[i] = huh.NewOption(fmt.Sprintf("%v %s", i+1, pkg), pkg)
	}
	return &options
}

func NewMultiSelectPackages(selected *[]string, options ...huh.Option[string]) *huh.MultiSelect[string] {
	return huh.NewMultiSelect[string]().
		Options(options...).
		Title("Select packages").
		Value(selected).
		Filterable(true).
		Height(min(len(options)+1, 10))
}

func PrintPackages(packages *pacman.Packages) {
	for _, item := range *packages {
		fmt.Println(item)
	}
}
