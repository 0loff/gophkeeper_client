package tui

import (
	"context"

	"github.com/rivo/tview"
)

func (t *Tui) AuthForm() *tview.Form {
	f := tview.NewForm().
		AddInputField("Login", "", 25, nil, nil).
		AddPasswordField("Password", "", 25, '*', nil).
		AddInputField("Email", "", 25, nil, nil).
		AddButton("Register", t.registerAction).
		AddButton("Quit", func() {
			t.AppView.Stop()
		})

	f.SetBorder(true)
	f.SetTitle("Register form")

	t.Form = f
	return f
}

func (t *Tui) registerAction() {
	LoginField := t.Form.GetFormItemByLabel("Login")
	PwdField := t.Form.GetFormItemByLabel("Password")
	EmailField := t.Form.GetFormItemByLabel("Email")

	login := LoginField.(*tview.InputField).GetText()
	password := PwdField.(*tview.InputField).GetText()
	email := EmailField.(*tview.InputField).GetText()

	if login == "" || password == "" || email == "" {
		t.showError("All form fields must be filled")
		return
	}

	t.App.TokenCh <- t.App.Requestor.NewRequest(context.Background(), t.App.JWT).SignUp(
		context.Background(),
		LoginField.(*tview.InputField).GetText(),
		PwdField.(*tview.InputField).GetText(),
		EmailField.(*tview.InputField).GetText(),
	)
}
