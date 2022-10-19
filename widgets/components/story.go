package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage/repository"
	"fyne.io/fyne/v2/widget"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"pikabu_api.go/rest/structs"
	"pikabu_api.go/widgets/types"
	"time"
)

func StoryBuild(data structs.PikabuStory, body types.Body) *fyne.Container {
	content := container.NewVBox()

	for _, part := range data.StoryData {
		if part.Type == "t" {
			markdown, _ := md.NewConverter("", true, nil).ConvertString(part.Data.(string))

			content.Add(widget.NewRichTextFromMarkdown(markdown))
		} else if part.Type == "i" {
			value := part.Data.(map[string]interface{})

			if uri, err := repository.ParseURI(value["large"].(string)); err == nil {
				image := canvas.NewImageFromURI(uri)
				image.FillMode = canvas.ImageFillOriginal

				content.Add(image)
			}
		}
	}

	tags := container.NewHBox()

	for _, tag := range data.Tags {
		tags.Add(widget.NewButton(tag, func() {

		}))
	}

	content.Add(tags)

	return container.NewVBox(
		widget.NewCard(data.StoryTitle, time.Unix(data.StoryTime, 0).Format("02.01.2006 15:04:05"), content),
	)
}
