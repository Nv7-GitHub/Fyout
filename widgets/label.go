package widgets

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

// Label is a label
type Label struct {
	Title      string
	DeleteFunc func()

	Text string
}

// Build builds the layout
func (l *Label) Build() fyne.CanvasObject {
	return widget.NewLabel(l.Text)
}

// BuildTree builds the tree view
func (l *Label) BuildTree() fyne.CanvasObject {
	ac := widget.NewAccordionContainer()
	content := widget.NewVBox()

	optionsbtn := newContextMenuButton("Options", fyne.NewMenu("",
		fyne.NewMenuItem("Remove", l.DeleteFunc),
		fyne.NewMenuItem("Rename", func() {
			renameDialog(&l.Title)
			Save()
		}),
		fyne.NewMenuItem("Change Text", func() {
			renameDialog(&l.Text)
			Save()
		}),
	))
	content.Append(optionsbtn)

	item := widget.NewAccordionItem(l.Title, content)
	ac.Append(item)
	return ac
}

// Delete provides a function to give the DeleteFunc
func (l *Label) Delete(deletefunc func()) {
	l.DeleteFunc = deletefunc
}

// Clone duplicates the widget with the same data
func (l *Label) Clone() Widget {
	c := *l
	return &c
}

// Serialize creats a serialized form with the same data
func (l *Label) Serialize() WidgetSerialized {
	return &LabelSerialized{
		Title: l.Title,
		Text:  l.Text,
	}
}

// LabelSerialized is the serialized form of a button
type LabelSerialized struct {
	Title string
	Text  string
}

// Deserialize creates a deserialized version of the widget
func (l *LabelSerialized) Deserialize(deleteFunc func()) Widget {
	return &Label{
		Title:      l.Title,
		Text:       l.Text,
		DeleteFunc: deleteFunc,
	}
}
