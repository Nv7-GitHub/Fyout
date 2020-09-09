package Widgets

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

var root Widget
var widgets []Widget
var selected int
var widgetBtns []fyne.CanvasObject
var w *fyne.Window

// Widget allows you to recursively build the layouts
type Widget interface {
	Build() fyne.CanvasObject
	BuildTree() fyne.CanvasObject
	Delete(func())
	Clone() Widget
}

// Root is the root element
type Root struct {
	Child Widget
}

// Build builds the layout
func (r *Root) Build() fyne.CanvasObject {
	if r.Child != nil {
		return r.Child.Build()
	}
	return canvas.NewRectangle(theme.BackgroundColor())
}

// BuildTree builds the layout for the tree view
func (r *Root) BuildTree() fyne.CanvasObject {
	if r.Child == nil {
		return widget.NewButton("Add Selected", func() {
			cap := widgets[selected].Clone()
			r.Child = cap
			r.Child.Delete(func() {
				r.Child = nil
				UpdateUI()
			})
			UpdateUI()
		})
	}

	return r.Child.BuildTree()
}

func (r *Root) Delete(func()) {}

func (r *Root) Clone() Widget { return *new(Widget) }

// GenWidgets generates all the widgets
func GenWidgets() {
	root = new(Root)
	widgets = []Widget{
		&Vbox{
			Title:    "NewVBox",
			Children: make([]Widget, 0),
		},
		&Hbox{
			Title:    "NewHBox",
			Children: make([]Widget, 0),
		},
		&Button{
			Title: "NewButton",
			Text:  "NewButton",
		},
	}

	widgetNames := []string{"VBox", "HBox", "Button"}

	for i, val := range widgetNames {
		cap := i
		widgetBtns = append(widgetBtns, widget.NewButton(val, func() { ChangeSelected(cap) }))
	}

	selected = 0
	ChangeSelected(0)
}

// ChangeSelected changes the selected widget
func ChangeSelected(newselected int) {
	widgetBtns[selected].(*widget.Button).Enable()
	widgetBtns[newselected].(*widget.Button).Disable()
	selected = newselected
}

// Init receives the window and initializes widgets
func Init(win *fyne.Window) {
	w = win
	GenWidgets()
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
