package Widgets

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

var onclicks map[string]func()

// Button is a button
type Button struct {
	Title      string
	DeleteFunc func()

	OnClick string
	Text    string
}

// Build builds the layout
func (b *Button) Build() fyne.CanvasObject {
	button := widget.NewButton(b.Text, func() {
		_, exists := onclicks[b.OnClick]
		if exists {
			onclicks[b.OnClick]()
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

func (b *Button) Serialize() WidgetSerialized {
	return &ButtonSerialized{
		Title:   b.Title,
		OnClick: b.OnClick,
		Text:    b.Text,
	}
}

// ButtonSerialized is the serialized form of a button
type ButtonSerialized struct {
	Title string

	OnClick string
	Text    string
}

func (b *ButtonSerialized) Deserialize(deleteFunc func()) Widget {
	return &Button{
		Title:      b.Title,
		OnClick:    b.OnClick,
		Text:       b.Text,
		DeleteFunc: deleteFunc,
	}
}
