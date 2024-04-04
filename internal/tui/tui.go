package tui

import (
	"context"

	"github.com/0loff/gophkeeper_client/internal/app"
	tcell "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	pb "github.com/0loff/gophkeeper_server/proto"
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

	go tui.UserDataWorker(tui.App.TokenCh, tui.App.StatusCh)

	return tui
}

func (t *Tui) UserDataWorker(tokenCh, statusCh chan string) {
	for {
		select {
		case t.App.JWT = <-tokenCh:
			t.App.Textdata = t.App.Requestor.NewRequest(context.Background(), t.App.JWT).GetTextData()

			t.Preview.Clear()
			t.renderDataLists()

			t.ShowCreateDataModal()

		case status := <-statusCh:
			if status == "success" {
				t.App.Textdata = t.App.Requestor.NewRequest(context.Background(), t.App.JWT).GetTextData()

				t.renderDataLists()
			}
		}
	}
}

func (t *Tui) renderDataLists() {
	t.DataList.Clear()
	t.queueUpdateDraw(func() {
		for _, textdata := range t.App.Textdata.Data {
			t.DataList.AddItem(textdata.Metainfo, textdata.Text, rune('*'), func() {
				t.viewTextData(textdata)
			})
		}
	})
}

func (t *Tui) viewTextData(data *pb.TextdataEntry) {
	t.queueUpdateDraw(func() {
		t.Preview.Clear()
		t.Preview.
			AddItem(t.TextdataForm(data, "Update"), 1, 1, 2, 4, 0, 0, true)
	})
}

func (t *Tui) queueUpdate(f func()) {
	go func() {
		t.AppView.QueueUpdate(f)
	}()
}

func (t *Tui) queueUpdateDraw(f func()) {
	go func() {
		t.AppView.QueueUpdateDraw(f)
	}()
}

func (t *Tui) Init() {
	t.AppView = tview.NewApplication()
	t.Pages = tview.NewPages()

	t.DataList = tview.NewList().ShowSecondaryText(false).SetSecondaryTextColor(tcell.ColorDimGray)
	// t.Preview = tview.NewList().ShowSecondaryText(true).SetSecondaryTextColor(tcell.ColorDimGray)
	t.Preview = tview.NewGrid()
	t.FooterText = tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText(TitleFooterView).SetTextColor(tcell.ColorGray)

	t.DataList.SetSelectedFunc(t.dataSelected)

	navigate := tview.NewGrid().SetRows(0).
		AddItem(t.DataList, 0, 0, 1, 1, 0, 0, true)

	t.Grid = tview.NewGrid().
		SetRows(0, 2).
		SetColumns(40, 0).
		SetBorders(false).
		AddItem(navigate, 0, 0, 1, 1, 0, 0, true).
		AddItem(t.Preview, 0, 1, 1, 1, 0, 0, false).
		AddItem(t.FooterText, 1, 0, 1, 2, 0, 0, false)

	t.setupKeyBoard()

	if t.App.JWT == "" {
		t.ShowLoginForm()
	}

	if err := t.AppView.SetRoot(t.Grid, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func (t *Tui) ShowAuthForm() {
	t.Preview.Clear()
	t.queueUpdateDraw(func() {
		t.Preview.
			SetRows(0, 0, 0, 0, 0).SetColumns(0, 0, 0, 0, 0, 0, 0).
			AddItem(t.AuthForm(), 2, 2, 1, 2, 0, 0, true)
	})
}

func (t *Tui) ShowLoginForm() {
	t.Preview.Clear()
	t.queueUpdateDraw(func() {
		t.Preview.
			SetRows(0, 0, 0, 0, 0).SetColumns(0, 0, 0, 0, 0, 0, 0).
			AddItem(t.LoginForm(), 2, 2, 1, 2, 0, 0, true)
	})
}

func (t *Tui) ShowCreateTextDataForm() {
	t.Preview.Clear()

	t.queueUpdateDraw(func() {
		t.Preview.
			SetRows(0, 0, 0, 0, 0).SetColumns(0, 0, 0, 0, 0, 0, 0).
			AddItem(t.TextdataForm(&pb.TextdataEntry{}, "Create"), 2, 1, 2, 4, 0, 0, true)
	})
}

func (t *Tui) ShowCreateCredsDataForm() {
	t.Preview.Clear()

	t.queueUpdateDraw(func() {
		t.Preview.
			SetRows(0, 0, 0, 0, 0).SetColumns(0, 0, 0, 0, 0).
			AddItem(t.CredsdataForm(&pb.CredsdataEntry{}, "Create"), 2, 1, 1, 2, 0, 0, true)
	})
}

func (t *Tui) ShowCreateDataModal() {
	t.Preview.Clear()

	t.queueUpdateDraw(func() {
		t.Preview.AddItem(t.CreateDataModal(), 2, 2, 1, 2, 0, 0, true)
	})
}
