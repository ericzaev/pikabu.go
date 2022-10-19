package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"pikabu_api.go/widgets/pages"
)

type Main struct {
	top  *fyne.Container
	root *fyne.Container
	body *fyne.Container
}

func NewMain(root *fyne.Container) *Main {
	return &Main{root: root}
}

func (m *Main) Add(content *fyne.Container) {
	m.body.Add(content)

	m.topAction()
	m.bodyAction()
}

func (m *Main) Pop() {
	m.body.Remove(m.body.Objects[len(m.body.Objects)-1])

	m.topAction()
	m.bodyAction()
}

func (m *Main) topAction() {
	if m.top.Hidden && len(m.body.Objects) > 1 {
		m.top.Show()
	} else if !m.top.Hidden {
		m.top.Hide()
	}
}

func (m *Main) bodyAction() {
	for index, content := range m.body.Objects {
		if index == len(m.body.Objects)-1 {
			content.Show()
		} else {
			content.Hide()
		}
	}
}

func (m *Main) Build() {
	m.top = container.NewVBox(widget.NewButton("Back", func() {
		m.Pop()
	}))
	m.body = container.NewVBox()

	m.Add(pages.NewFeed(m).Build())

	m.root.Add(container.NewVBox(m.top, m.body))
}
