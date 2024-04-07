package tui

import (
	"context"

	"github.com/rivo/tview"
)

func (t *Tui) LoginForm() *tview.Form {
	f := tview.NewForm().
		AddInputField("Email", "", 25, nil, nil).
		AddPasswordField("Password", "", 25, '*', nil).
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

	token := t.App.Requestor.NewRequest(context.Background(), t.App.JWT).SignIn(
		context.Background(),
		EmailField.(*tview.InputField).GetText(),
		PwdField.(*tview.InputField).GetText(),
	)

	if token == "" {
		t.showError("Email or password is incorrect")
		return
	}

	t.App.TokenCh <- token
}
