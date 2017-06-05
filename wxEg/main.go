package main

import "github.com/dontpanic92/wxGo/wx"

func main() {
	wx.NewApp()
	f := wx.NewDialog(wx.NullWindow, -1, "Hello World")
	f.ShowModal()
	f.Destroy()
}
