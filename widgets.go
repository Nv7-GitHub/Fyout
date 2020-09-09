package Fyout

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

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
		&Button{
			Title: "NewButton",
			Text:  "NewButton",
		},
	}

	widgetNames := []string{"VBox", "Button"}

	for i, val := range widgetNames {
		cap := i
		widgetBtns = append(widgetBtns, widget.NewButton(val, func() { ChangeSelected(cap) }))
	}

	selected = 0
	ChangeSelected(0)
}
