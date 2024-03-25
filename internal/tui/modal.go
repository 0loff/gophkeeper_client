package tui

import "github.com/rivo/tview"

func (t *Tui) ShowModal(msg string, btns []string) *tview.Modal {
	return tview.NewModal().
		SetText(msg).
		AddButtons(btns).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "ok" {
				t.Pages.ShowPage("Auth")
			}
		})
}
