package main

import (
	"github.com/AllenDang/gform"
)

func main() {
	gform.Init()

	mainWindow := gform.NewForm(nil)
	mainWindow.SetPos(300, 100)
	mainWindow.SetSize(500, 300)
	mainWindow.SetCaption("Controls Demo")

	btn := gform.NewPushButton(mainWindow)
	btn.SetPos(10, 10)
	btn.OnLBUp().Bind(btn_onclick)

	mainWindow.Show()

	gform.RunMainLoop()
}
