package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func loginDialog() fyne.CanvasObject {
	return container.NewVBox(
		// canvas.NewImageFromResource(resourceIconPng),
		widget.NewLabel("Hello World!"),
	)
}
