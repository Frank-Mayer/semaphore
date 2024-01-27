package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func loadingDialog() fyne.CanvasObject {
	return container.NewVBox(
		container.NewCenter(
			container.NewGridWrap(fyne.NewSize(100, 100),
				canvas.NewImageFromResource(resourceIconPng),
			),
		),
		widget.NewLabel("Semaphore"),
	)
}
