package Fyout

import (
	"os"

	"fyne.io/fyne"

	"github.com/Nv7-Github/Fyout/widgets"
)

var w *fyne.Window

// NewBuilder creates the builder UI
func NewBuilder(path string, window *fyne.Window) {
	w = window
	_, err := os.Stat(path)
	exists := os.IsExist(err)
	if !exists {
		os.Create(path)
		Widgets.Init(w)
		Widgets.UpdateUI()
	} else {
		Widgets.Init(w)
		Widgets.UpdateUI()
	}
}
