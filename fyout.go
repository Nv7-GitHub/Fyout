package fyout

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
func NewBuilder(path string, app *fyne.App) {
	a = app
	win := (*app).NewWindow("Fyout Layout Builder")
	win.Resize(fyne.NewSize(800, 600))
	w = &win
	_, err := os.Stat(path)
	exists := !os.IsNotExist(err)
	if !exists {
		widgets.Init(w, path)
		widgets.Save()
		widgets.UpdateUI()
	} else {
		widgets.Init(w, path)
		widgets.Read()
		widgets.UpdateUI()
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
			widgets.UpdateUI()
		}),
		fyne.NewMenuItem("Light Theme", func() {
			(*a).Settings().SetTheme(theme.LightTheme())
			widgets.UpdateUI()
		}),
	)
	settingsMenu := fyne.NewMenu("Settings", themeMenu)
	mainMenu := fyne.NewMainMenu(settingsMenu)
	(*w).SetMainMenu(mainMenu)
}

// LoadLayout calls LoadLayout in the widget package
func LoadLayout(path string, funcs map[string]func()) fyne.CanvasObject {
	return widgets.LoadLayout(path, funcs)
}

// AddWidget calls AddWidget in the widget package
func AddWidget(name string, widget widgets.Widget) {
	widgets.AddWidget(name, widget)
}
