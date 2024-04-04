package tui

import (
	"context"

	"github.com/rivo/tview"
)

func (t *Tui) LoginForm() *tview.Form {
	f := tview.NewForm().
		AddInputField("Email", "", 20, nil, nil).
		AddPasswordField("Password", "", 20, '*', nil).
		AddButton("Login", t.loginAction).
		AddButton("Quit", func() {
			t.AppView.Stop()
		})

	f.SetBorder(true)
	f.SetTitle("Login form")

	t.Form = f
	return f
}

func (t *Tui) loginAction() {
	EmailField := t.Form.GetFormItemByLabel("Email")
	PwdField := t.Form.GetFormItemByLabel("Password")

	t.App.TokenCh <- t.App.Requestor.NewRequest(context.Background(), t.App.JWT).SignIn(
		context.Background(),
		EmailField.(*tview.InputField).GetText(),
		PwdField.(*tview.InputField).GetText(),
	)
}
