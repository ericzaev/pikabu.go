package pages

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"pikabu_api.go/widgets/types"
)

type Story struct {
	body types.Body
}

func NewStory(body types.Body) *Feed {
	return &Feed{body: body}
}

func (f *Story) Build() *fyne.Container {
	str := binding.NewString()
	str.Set("Hi!")

	return container.NewVBox(
		widget.NewButton("hide", func() {
			str.Set("World")
		}),
		widget.NewLabelWithData(str),
		widget.NewEntryWithData(str),
	)
}
