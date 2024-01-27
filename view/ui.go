package view

import (
	"fyne.io/fyne/v2/app"
	"github.com/Frank-Mayer/fyneflow"
)

type screenId string

const (
	screenLoading screenId = "loading"
	screenLogin   screenId = "login"
	screenError   screenId = "error"
)

var (
	flow  *fyneflow.Flow[screenId]
	myApp = app.New()
)

// ShowAndRun initializes the user interface.
func ShowAndRun() {
	my_window := myApp.NewWindow(myApp.Metadata().Name)

	flow = fyneflow.NewFlow[screenId](my_window)
	defer flow.Close()
	flow.Set(screenLoading, loadingDialog)
	flow.Set(screenLogin, loginDialog)
	flow.Set(screenError, errorDialog)

	my_window.ShowAndRun()
}
