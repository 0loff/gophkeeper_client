package tui

import (
	"github.com/0loff/gophkeeper_client/internal/app"
	tcell "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	// TitleFooterView = "Navigate [ Tab / Shift-Tab ] · Focus [ Ctrl-F ] · Exit [ Ctrl-C ] \n Tables specific: Describe [ e ] · Preview [ p ]"
	TitleFooterView = "Register new user [ Ctrl + R ] · Sign in form [Ctrl + L] · Focus [ Ctrl-F ] · Exit [ Ctrl + Q ] \n Tables specific: Describe [ e ] · Preview [ p ]"
)

type Tui struct {
	App *app.App

	AppView *tview.Application
	Form    *tview.Form
	Pages   *tview.Pages
	Grid    *tview.Grid

	DataList   *tview.List
	Preview    *tview.Grid
	FooterText *tview.TextView
}

func NewAppView(a *app.App) *Tui {
	tui := &Tui{
		App: a,
	}

	return tui
}

func (tui *Tui) queueUpdate(f func()) {
	go func() {
		tui.AppView.QueueUpdate(f)
	}()
}

func (tui *Tui) queueUpdateDraw(f func()) {
	go func() {
		tui.AppView.QueueUpdateDraw(f)
	}()
}

func (t *Tui) Init() {
	t.AppView = tview.NewApplication()
	t.Pages = tview.NewPages()

	t.DataList = tview.NewList().ShowSecondaryText(true).SetSecondaryTextColor(tcell.ColorDimGray)
	// t.Preview = tview.NewList().ShowSecondaryText(true).SetSecondaryTextColor(tcell.ColorDimGray)
	t.Preview = tview.NewGrid().SetRows(0, 0, 0, 0, 0).SetColumns(0, 0, 0, 0, 0)
	t.FooterText = tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText(TitleFooterView).SetTextColor(tcell.ColorGray)

	navigate := tview.NewGrid().SetRows(0).
		AddItem(t.DataList, 0, 0, 1, 1, 0, 0, true)

	t.Grid = tview.NewGrid().
		SetRows(0, 2).
		SetColumns(40, 0).
		SetBorders(true).
		AddItem(navigate, 0, 0, 1, 1, 0, 0, true).
		AddItem(t.Preview, 0, 1, 1, 1, 0, 0, false).
		AddItem(t.FooterText, 1, 0, 1, 2, 0, 0, false)

	t.AppView.SetInputCapture(t.inputActions)

	// t.Pages.AddPage("Modal", t.ShowModal(), true, false)
	// t.Pages.AddPage("Auth", t.AuthView(), true, false)

	// modal := tview.NewModal().
	// 	SetText("This is a new modal window")

	// t.pages.AddPage("Login", modal, true, false)

	// if t.App.Token == "" {

	// 		t.ShowAuthForm()
	// 	})
	// }

	if err := t.AppView.SetRoot(t.Grid, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func (t *Tui) inputActions(e *tcell.EventKey) *tcell.EventKey {
	switch pressed_key := e.Rune(); pressed_key {
	case rune(tcell.KeyCtrlQ):
		t.AppView.Stop()
	case rune(tcell.KeyCtrlL):
		t.ShowLoginForm()
	case rune(tcell.KeyCtrlR):
		t.ShowAuthForm()
	}

	return e
}

func (t *Tui) ShowAuthForm() {
	t.Preview.Clear()
	t.queueUpdateDraw(func() {
		t.Preview.AddItem(t.AuthForm(), 2, 2, 1, 1, 0, 0, true)
	})
}

func (t *Tui) ShowLoginForm() {
	t.Preview.Clear()
	t.queueUpdateDraw(func() {
		t.Preview.AddItem(t.LoginForm(), 2, 2, 1, 1, 0, 0, true)
	})
}
