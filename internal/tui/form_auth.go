package tui

import (
	"context"

	"github.com/rivo/tview"
)

func (t *Tui) AuthForm() *tview.Form {
	f := tview.NewForm().
		AddInputField("Login", "", 20, nil, nil).
		AddPasswordField("Password", "", 20, '*', nil).
		AddInputField("Email", "", 20, nil, nil).
		AddButton("Register", t.registerAction)

	f.SetBorder(true)
	f.SetTitle("Register form")

	t.Form = f
	return f
}

func (t *Tui) registerAction() {
	LoginField := t.Form.GetFormItemByLabel("Login")
	PwdField := t.Form.GetFormItemByLabel("Password")
	EmailField := t.Form.GetFormItemByLabel("Email")

	t.App.TokenCh <- t.App.Requestor.NewRequest(context.Background(), t.App.JWT).SignUp(
		context.Background(),
		LoginField.(*tview.InputField).GetText(),
		PwdField.(*tview.InputField).GetText(),
		EmailField.(*tview.InputField).GetText(),
	)
}
