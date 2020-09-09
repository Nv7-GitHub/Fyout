package Widgets

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

// Vbox is a fyne VBox
type Button struct {
	Title      string
	DeleteFunc func()

	Onclick func()
	Text    string
}

// Build builds the layout
func (b *Button) Build() fyne.CanvasObject {
	button := widget.NewButton(b.Text, func() {
		if b.Onclick != nil {
			b.Onclick()
		}
	})

	return button
}

// BuildTree builds the tree view
func (b *Button) BuildTree() fyne.CanvasObject {
	ac := widget.NewAccordionContainer()
	content := widget.NewVBox()

	removebtn := widget.NewButton("Remove", b.DeleteFunc)
	content.Append(removebtn)

	renamebtn := widget.NewButton("Rename", func() {
		b.Title = "Renamed"
		UpdateUI()
	})
	content.Append(renamebtn)

	item := widget.NewAccordionItem(b.Title, content)
	ac.Append(item)
	return ac
}

func (b *Button) Delete(deletefunc func()) {
	b.DeleteFunc = deletefunc
}

func (b *Button) Clone() Widget {
	c := *b
	return &c
}
