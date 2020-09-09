package main

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

	addchildbtn := widget.NewButton("Add Selected", func() {
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
	})
	content.Append(addchildbtn)

	removebtn := widget.NewButton("Remove", v.DeleteFunc)
	content.Append(removebtn)

	renamebtn := widget.NewButton("Rename", func() {
		v.Title = "Renamed"
		UpdateUI()
	})
	content.Append(renamebtn)

	item := widget.NewAccordionItem(v.Title, content)
	ac.Append(item)
	return ac
}

func (v *Vbox) Delete(deletefunc func()) {
	v.DeleteFunc = deletefunc
}

func (v *Vbox) Clone() Widget {
	c := *v
	return &c
}
