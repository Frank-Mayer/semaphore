package main

import (
	"github.com/Frank-Mayer/semaphore/model/data"
	"github.com/Frank-Mayer/semaphore/view"
	"github.com/charmbracelet/log"
)

func main() {
	if err := data.Load(); err != nil {
		log.Fatal(err)
	}
	view.ShowAndRun()
	if err := data.Save(); err != nil {
		log.Fatal(err)
	}
}
