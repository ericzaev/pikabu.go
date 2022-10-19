package pages

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"pikabu_api.go/rest"
	"pikabu_api.go/widgets/components"
	"pikabu_api.go/widgets/types"
)

type Feed struct {
	body types.Body
}

func NewFeed(body types.Body) *Feed {
	return &Feed{body: body}
}

func (f *Feed) Build() *fyne.Container {
	var content *fyne.Container

	content = container.NewVBox()

	go func() {
		result := rest.PikabuFeed("hot_act")

		for _, story := range result.Data {
			content.Add(components.StoryBuild(story, f.body))
		}
	}()

	return content
}
