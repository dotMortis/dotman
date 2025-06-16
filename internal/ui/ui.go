package ui

import (
	"dotman/internal/pacman"
	"fmt"
	"math"

	"github.com/charmbracelet/huh"
)

func NewSingleGroupForm(fields ...huh.Field) *huh.Form {
	return huh.NewForm(huh.NewGroup(fields...)).WithTheme(huh.ThemeCatppuccin())
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
	reset := "\033[0m"
	total := len(*packages)

	for i, item := range *packages {
		// Calculate position in the color cycle (0 to 1)
		pos := float64(i) / float64(total)

		// Use sine waves to create smooth transitions
		r := int(math.Sin(pos*2*math.Pi)*127 + 128)
		g := int(math.Sin((pos*2*math.Pi+2*math.Pi/3))*127 + 128)
		b := int(math.Sin((pos*2*math.Pi+4*math.Pi/3))*127 + 128)

		// Create the color escape sequence
		color := fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
		fmt.Printf("%s%s%s\n", color, item, reset)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
