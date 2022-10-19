package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/joho/godotenv"
	"pikabu_api.go/widgets"
)

func main() {
	godotenv.Load()

	instance := app.New()

	window := instance.NewWindow("Container")
	content := container.NewVBox()

	widgets.NewMain(content).Build()

	window.Resize(fyne.NewSize(640, 480))
	window.SetContent(container.NewScroll(content))
	window.ShowAndRun()
}
