package Widgets

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

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
		_, exists := root.(*Root).Funcs[b.OnClick]
		if exists {
			root.(*Root).Funcs[b.OnClick]()
		}
	})

	return button
}

// BuildTree builds the tree view
func (b *Button) BuildTree() fyne.CanvasObject {
	ac := widget.NewAccordionContainer()
	content := widget.NewVBox()

	optionsbtn := newContextMenuButton("Options", fyne.NewMenu("",
		fyne.NewMenuItem("Remove", b.DeleteFunc),
		fyne.NewMenuItem("Rename", func() {
			renameDialog(&b.Title)
			Save()
		}),
		fyne.NewMenuItem("Change Text", func() {
			renameDialog(&b.Text)
			Save()
		}),
	))
	content.Append(optionsbtn)

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
