package tui

import (
	"fmt"

	tcell "github.com/gdamore/tcell/v2"
)

func (t *Tui) setupKeyBoard() {
	t.AppView.SetInputCapture(func(e *tcell.EventKey) *tcell.EventKey {
		switch pressed_key := e.Rune(); pressed_key {
		case rune(tcell.KeyCtrlQ):
			t.AppView.Stop()
		case rune(tcell.KeyCtrlL):
			t.ShowLoginForm()
		case rune(tcell.KeyCtrlR):
			t.ShowAuthForm()
		case rune(tcell.KeyCtrlI):
			t.showMessage(
				fmt.Sprintf("App version: %s · Actual build commit: %s · Build date: %s",
					t.App.Info.Version, t.App.Info.Commit, t.App.Info.Date))
		case rune(tcell.KeyCtrlM):
			if t.App.JWT == "" {
				break
			}
			t.ShowCreateDataModal()
		}

		return e
	})
}
