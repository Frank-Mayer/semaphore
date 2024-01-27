package view

import (
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/charmbracelet/log"
)

func errorDialog() fyne.CanvasObject {
	err, _ := flow.UseStateStr("error", "").Get()
	log.Debug("Displaying error ui", "error", err)
	vbox := container.NewVBox(
		canvas.NewText("Da lief was schief :(", color.RGBA{R: 255, G: 0, B: 0, A: 255}),
	)
	if err != "" {
		// split error message into multiple lines
		for _, line := range strings.Split(err, "\n") {
			vbox.Add(canvas.NewText(line, color.White))
		}
	}
	return vbox
}

func renderError(err error) fyne.CanvasObject {
	if err := flow.UseStateStr("error", "").Set(err.Error()); err != nil {
		log.Error("Failed to set error", "error", err)
	}
	return errorDialog()
}

func displayError(err error) {
	if err == nil {
		log.Error("called displayError with nil error")
		return
	}
	log.Error("Error", "error", err)
	if err := flow.UseStateStr("error", "").Set(err.Error()); err != nil {
		log.Error("Failed to set error", "error", err)
	}
	_ = flow.GoTo(screenError)
}
