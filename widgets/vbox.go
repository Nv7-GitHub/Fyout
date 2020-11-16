package Widgets

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

// Vbox is a fyne VBox
type Vbox struct {
	Children   []Widget
	Title      string
	DeleteFunc func()
}

// Build builds the layout
func (v *Vbox) Build() fyne.CanvasObject {
	vbox := widget.NewVBox()

	for _, child := range v.Children {
		vbox.Append(child.Build())
	}

	return vbox
}

// BuildTree builds the tree view
func (v *Vbox) BuildTree() fyne.CanvasObject {
	ac := widget.NewAccordionContainer()
	content := widget.NewVBox()

	for _, child := range v.Children {
		childtree := child.BuildTree()
		rect := canvas.NewRectangle(theme.BackgroundColor())
		rect.SetMinSize(fyne.NewSize(20, 0))
		hbox := widget.NewHBox(rect, childtree)
		content.Append(hbox)
	}

	optionsbtn := newContextMenuButton("Options", fyne.NewMenu("",
		fyne.NewMenuItem("Add Selected", func() {
			child := widgets[selected].Clone()
			pos := len(v.Children)
			child.Delete(func() {
				copy(v.Children[pos:], v.Children[pos+1:])
				v.Children[len(v.Children)-1] = nil
				v.Children = v.Children[:len(v.Children)-1]
				UpdateUI()
			})
			v.Children = append(v.Children, child)
			UpdateUI()
		}),
		fyne.NewMenuItem("Remove", v.DeleteFunc),
		fyne.NewMenuItem("Rename", func() { renameDialog(&v.Title) }),
	))
	content.Append(optionsbtn)

	item := widget.NewAccordionItem(v.Title, content)
	ac.Append(item)
	return ac
}

// Delete provides a function to give the DeleteFunc
func (v *Vbox) Delete(deletefunc func()) {
	v.DeleteFunc = deletefunc
}

// Clone duplicates the widget with the same data
func (v *Vbox) Clone() Widget {
	c := *v
	return &c
}

// Serialize creats a serialized form with the same data
func (v *Vbox) Serialize() WidgetSerialized {
	childS := make([]WidgetSerialized, len(v.Children))
	for i, child := range v.Children {
		childS[i] = child.Serialize()
	}
	return &VboxSerialized{
		Title:    v.Title,
		Children: childS,
	}
}

// VboxSerialized a serialized vbox
type VboxSerialized struct {
	Title    string
	Children []WidgetSerialized
}

// Deserialize creates a deserialized version of the widget
func (v *VboxSerialized) Deserialize(deleteFunc func()) Widget {
	vbox := Vbox{
		Title:      v.Title,
		Children:   make([]Widget, len(v.Children)),
		DeleteFunc: deleteFunc,
	}
	for i, child := range v.Children {
		vbox.Children[i] = child.Deserialize(func() {
			copy(vbox.Children[i:], vbox.Children[i+1:])
			vbox.Children[len(vbox.Children)-1] = nil
			vbox.Children = vbox.Children[:len(vbox.Children)-1]
			UpdateUI()
		})
	}
	return &vbox
}
