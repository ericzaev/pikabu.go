package types

import "fyne.io/fyne/v2"

type Body interface {
	Add(*fyne.Container)
	Pop()
}
