package Widgets

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

// Hbox is a fyne Hbox
type Hbox struct {
	Children   []Widget
	Title      string
	DeleteFunc func()
}

// Build builds the layout
func (h *Hbox) Build() fyne.CanvasObject {
	Hbox := widget.NewHBox()

	for _, child := range h.Children {
		Hbox.Append(child.Build())
	}

	return Hbox
}

// BuildTree builds the tree view
func (h *Hbox) BuildTree() fyne.CanvasObject {
	ac := widget.NewAccordionContainer()
	content := widget.NewVBox()

	for _, child := range h.Children {
		childtree := child.BuildTree()
		rect := canvas.NewRectangle(theme.BackgroundColor())
		rect.SetMinSize(fyne.NewSize(20, 0))
		hbox := widget.NewHBox(rect, childtree)
		content.Append(hbox)
	}

	optionsbtn := newContextMenuButton("Options", fyne.NewMenu("",
		fyne.NewMenuItem("Add Selected", func() {
			child := widgets[selected].Clone()
			pos := len(h.Children)
			child.Delete(func() {
				copy(h.Children[pos:], h.Children[pos+1:])
				h.Children[len(h.Children)-1] = nil
				h.Children = h.Children[:len(h.Children)-1]
				UpdateUI()
			})
			h.Children = append(h.Children, child)
			UpdateUI()
		}),
		fyne.NewMenuItem("Remove", h.DeleteFunc),
		fyne.NewMenuItem("Rename", func() { renameDialog(&h.Title) }),
	))
	content.Append(optionsbtn)

	item := widget.NewAccordionItem(h.Title, content)
	ac.Append(item)
	return ac
}

// Delete provides a function to give the DeleteFunc
func (h *Hbox) Delete(deletefunc func()) {
	h.DeleteFunc = deletefunc
}

// Clone duplicates the widget with the same data
func (h *Hbox) Clone() Widget {
	c := *h
	return &c
}

// Serialize creats a serialized form with the same data
func (h *Hbox) Serialize() WidgetSerialized {
	childS := make([]WidgetSerialized, len(h.Children))
	for i, child := range h.Children {
		childS[i] = child.Serialize()
	}
	return &VboxSerialized{
		Title:    h.Title,
		Children: childS,
	}
}

// HboxSerialized a serialized vbox
type HboxSerialized struct {
	Title    string
	Children []WidgetSerialized
}

// Deserialize creates a deserialized version of the widget
func (h *HboxSerialized) Deserialize(deleteFunc func()) Widget {
	hbox := Hbox{
		Title:      h.Title,
		Children:   make([]Widget, len(h.Children)),
		DeleteFunc: deleteFunc,
	}
	for i, child := range h.Children {
		hbox.Children[i] = child.Deserialize(func() {
			copy(hbox.Children[i:], hbox.Children[i+1:])
			hbox.Children[len(hbox.Children)-1] = nil
			hbox.Children = hbox.Children[:len(hbox.Children)-1]
			UpdateUI()
		})
	}
	return &hbox
}
