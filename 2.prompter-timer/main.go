package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.NewWithID("com.corporate.promptertimer")
	myApp.Settings().SetTheme(theme.DarkTheme())

	myWindow := myApp.NewWindow("Prompter Timer")
	myWindow.SetFullScreen(true)

	var duration = 5 * time.Minute
	var remaining = duration

	timerText := canvas.NewText(formatDuration(remaining), theme.ForegroundColor())
	timerText.TextStyle = fyne.TextStyle{Bold: true}
	timerText.Alignment = fyne.TextAlignCenter
	timerText.TextSize = 400

	input := widget.NewEntry()
	input.SetPlaceHolder("Minutes (e.g. 5)")
	input.SetText("5")

	startBtn := widget.NewButtonWithIcon("Start", theme.MediaPlayIcon(), nil)
	stopBtn := widget.NewButtonWithIcon("Stop", theme.MediaStopIcon(), nil)
	resetBtn := widget.NewButtonWithIcon("Reset", theme.ViewRefreshIcon(), nil)

	buttonSize := fyne.NewSize(120, 50)
	startBtn.Resize(buttonSize)
	stopBtn.Resize(buttonSize)
	resetBtn.Resize(buttonSize)

	buttons
}
