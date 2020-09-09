package Fyout

import (
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

// Widget allows you to recursively build the layouts
type Widget interface {
	Build() fyne.CanvasObject
	BuildTree() fyne.CanvasObject
	Delete(func())
	Clone() Widget
}

var root Widget
var widgets []Widget
var selected int
var widgetBtns []fyne.CanvasObject
var w *fyne.Window

// NewBuilder creates the builder UI
func NewBuilder(path string, window *fyne.Window) {
	w = window
	_, err := os.Stat(path)
	exists := os.IsExist(err)
	if !exists {
		os.Create(path)
		GenWidgets()
		UpdateUI()
	} else {
		GenWidgets()
		UpdateUI()
	}
}

// UpdateUI builds the editor
func UpdateUI() {
	tree := root.BuildTree()
	treescroll := widget.NewScrollContainer(tree)

	layout := root.Build()

	vbox := widget.NewVBox(widgetBtns...)
	scroll := widget.NewScrollContainer(vbox)

	vsplit := widget.NewVSplitContainer(treescroll, scroll)
	hsplit := widget.NewHSplitContainer(vsplit, layout)
	(*w).SetContent(hsplit)
}

// ChangeSelected changes the selected widget
func ChangeSelected(newselected int) {
	widgetBtns[selected].(*widget.Button).Disable()
	widgetBtns[newselected].(*widget.Button).Disable()
	selected = newselected
}
