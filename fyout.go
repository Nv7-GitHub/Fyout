package Fyout

import (
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/theme"

	"github.com/Nv7-Github/Fyout/widgets"
)

var w *fyne.Window
var a *fyne.App
var file *os.File

// NewBuilder creates the builder UI
func NewBuilder(path string, app *fyne.App) error {
	a = app
	win := (*app).NewWindow("Fyout Layout Builder")
	win.Resize(fyne.NewSize(800, 600))
	w = &win
	_, err := os.Stat(path)
	exists := !os.IsNotExist(err)
	if !exists {
		file, err = os.Create(path)
		if err != nil {
			return err
		}
		Widgets.Init(w, path)
		Widgets.Save()
		Widgets.Read()
		Widgets.UpdateUI()
	} else {
		Widgets.Init(w, path)
		Widgets.Read()
		Widgets.UpdateUI()
	}
	MainMenu(w)
	(*w).Show()
	return nil
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
