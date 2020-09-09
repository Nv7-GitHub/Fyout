package Fyout

import (
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/theme"

	"github.com/Nv7-Github/Fyout/widgets"
)

var w *fyne.Window
var a *fyne.App

// NewBuilder creates the builder UI
func NewBuilder(path string, app *fyne.App) {
	a = app
	win := (*app).NewWindow("Fyout Layout Builder")
	win.Resize(fyne.NewSize(800, 600))
	w = &win
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
	MainMenu(w)
	(*w).Show()
}

// MainMenu creates the menu for the builder
func MainMenu(win *fyne.Window) {
	themeMenu := fyne.NewMenuItem("Theme", nil)
	themeMenu.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Dark Theme", func() {
			(*a).Settings().SetTheme(theme.DarkTheme())
			Widgets.UpdateUI()
		}),
		fyne.NewMenuItem("Light Theme", func() {
			(*a).Settings().SetTheme(theme.LightTheme())
			Widgets.UpdateUI()
		}),
	)
	settingsMenu := fyne.NewMenu("Settings", themeMenu)
	mainMenu := fyne.NewMainMenu(settingsMenu)
	(*w).SetMainMenu(mainMenu)
}
