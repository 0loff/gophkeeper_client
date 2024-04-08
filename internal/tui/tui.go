package tui

import (
	"context"
	"time"

	"github.com/0loff/gophkeeper_client/internal/app"
	"github.com/0loff/gophkeeper_client/pkg/encryptor"
	tcell "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	pb "github.com/0loff/gophkeeper_server/proto"
)

var (
	// TitleFooterView = "Navigate [ Tab / Shift-Tab ] · Focus [ Ctrl-F ] · Exit [ Ctrl-C ] \n Tables specific: Describe [ e ] · Preview [ p ]"
	TitleFooterView = "Register new user [ Ctrl + R ] · Sign in form [Ctrl + L] · View App Info [ Ctrl + I ] · Exit [ Ctrl + Q ] \n"
)

type Tui struct {
	App *app.App

	AppView *tview.Application
	Form    *tview.Form
	Pages   *tview.Pages
	Grid    *tview.Grid

	TextDataList  *tview.List
	CredsDataList *tview.List
	CardsDataList *tview.List
	BinDataList   *tview.List

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
			t.renderTextDataList()

			t.App.Credsdata = t.App.Requestor.NewRequest(context.Background(), t.App.JWT).GetCredsData()
			t.renderCredsDataList()

			t.App.CardsData = t.App.Requestor.NewRequest(context.Background(), t.App.JWT).GetCardsData()
			t.renderCardsDataList()

			t.App.Bindata = t.App.Requestor.NewRequest(context.Background(), t.App.JWT).GetBinData()
			t.renderBinDataList()

			t.Preview.Clear()
			t.ShowCreateDataModal()

		case status := <-statusCh:
			if status == "success" {
				t.App.Textdata = t.App.Requestor.NewRequest(context.Background(), t.App.JWT).GetTextData()
				t.renderTextDataList()

				t.App.Credsdata = t.App.Requestor.NewRequest(context.Background(), t.App.JWT).GetCredsData()
				t.renderCredsDataList()

				t.App.CardsData = t.App.Requestor.NewRequest(context.Background(), t.App.JWT).GetCardsData()
				t.renderCardsDataList()

				t.App.Bindata = t.App.Requestor.NewRequest(context.Background(), t.App.JWT).GetBinData()
				t.renderBinDataList()
			}
		}
	}
}

func (t *Tui) resetMessage() {
	t.queueUpdateDraw(func() {
		t.FooterText.SetText(TitleFooterView).SetTextColor(tcell.ColorGray)
	})
}

func (t *Tui) showMessage(msg string) {
	t.queueUpdateDraw(func() {
		t.FooterText.SetText(msg).SetTextColor(tcell.ColorGreen)
	})
	go time.AfterFunc(3*time.Second, t.resetMessage)
}

func (t *Tui) showError(msg string) {
	t.queueUpdateDraw(func() {
		t.FooterText.SetText(msg).SetTextColor(tcell.ColorRed)
	})
	go time.AfterFunc(3*time.Second, t.resetMessage)
}

func (t *Tui) renderTextDataList() {
	t.TextDataList.Clear()
	t.queueUpdateDraw(func() {
		for _, textdata := range t.App.Textdata.Data {
			t.TextDataList.AddItem(textdata.Metainfo, textdata.Text, rune('*'), func() {
				t.viewTextData(textdata)
			})
		}
	})
}

func (t *Tui) renderCredsDataList() {
	t.CredsDataList.Clear()
	t.queueUpdateDraw(func() {
		for _, credsdata := range t.App.Credsdata.Data {
			uname, err := encryptor.Decrypt(credsdata.Username, t.App.GetUserKey())
			if err != nil {
				panic(err)
			}

			t.CredsDataList.AddItem(credsdata.Metainfo, string(uname), rune('*'), func() {
				t.viewCredsData(credsdata)
			})
		}
	})
}

func (t *Tui) renderCardsDataList() {
	t.CardsDataList.Clear()
	t.queueUpdateDraw(func() {
		for _, cardsdata := range t.App.CardsData.Data {
			t.CardsDataList.AddItem(cardsdata.Metainfo, cardsdata.Pan, rune('*'), func() {
				t.viewCardsData(cardsdata)
			})
		}
	})
}

func (t *Tui) renderBinDataList() {
	t.BinDataList.Clear()
	t.queueUpdateDraw(func() {
		for _, bindata := range t.App.Bindata.Data {
			t.BinDataList.AddItem(bindata.Metainfo, string(bindata.Binary), rune('*'), func() {
				t.viewBinData(bindata)
			})
		}
	})
}

func (t *Tui) viewTextData(data *pb.TextdataEntry) {
	t.queueUpdateDraw(func() {
		t.Preview.Clear()
		t.Preview.
			SetRows(0, 0, 0, 0, 0, 0).SetColumns(0, 0, 0, 0, 0, 0, 0).
			AddItem(t.TextdataForm(data, "Update"), 2, 1, 2, 4, 0, 0, true)
	})
}

func (t *Tui) viewCredsData(data *pb.CredsdataEntry) {
	t.queueUpdateDraw(func() {
		t.Preview.Clear()
		t.Preview.
			SetRows(0, 0, 0, 0, 0).SetColumns(0, 0, 0, 0, 0).
			AddItem(t.CredsdataForm(data, "Update"), 2, 1, 1, 2, 0, 0, true)
	})
}

func (t *Tui) viewCardsData(data *pb.CardsdataEntry) {
	t.queueUpdateDraw(func() {
		t.Preview.Clear()
		t.Preview.
			SetRows(0, 0, 0, 0, 0, 0, 0, 0).SetColumns(0, 0, 0, 0, 0, 0, 0).
			AddItem(t.CardsdataForm(data, "Update"), 3, 2, 2, 2, 0, 0, true)
	})
}

func (t *Tui) viewBinData(data *pb.BindataEntry) {
	t.queueUpdateDraw(func() {
		t.Preview.Clear()
		t.Preview.
			SetRows(0, 0, 0, 0, 0, 0).SetColumns(0, 0, 0, 0, 0, 0, 0).
			AddItem(t.BindataForm(data, "Update"), 2, 1, 2, 4, 0, 0, true)
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

	t.TextDataList = tview.NewList().ShowSecondaryText(false)
	t.CredsDataList = tview.NewList().ShowSecondaryText(false)
	t.CardsDataList = tview.NewList().ShowSecondaryText(false)
	t.BinDataList = tview.NewList().ShowSecondaryText(false)

	t.Preview = tview.NewGrid()
	t.FooterText = tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText(TitleFooterView).SetTextColor(tcell.ColorGray)

	t.TextDataList.SetSelectedFunc(t.TextdataSelected)
	t.CredsDataList.SetSelectedFunc(t.CredsdataSelected)
	t.CardsDataList.SetSelectedFunc(t.CardsdataSelected)
	t.BinDataList.SetSelectedFunc(t.BindataSelected)

	navigate := tview.NewGrid().SetRows(0, 0, 0, 0).
		AddItem(t.TextDataList, 0, 0, 1, 1, 0, 0, true).
		AddItem(t.CredsDataList, 1, 0, 1, 1, 0, 0, true).
		AddItem(t.CardsDataList, 2, 0, 1, 1, 0, 0, true).
		AddItem(t.BinDataList, 3, 0, 1, 1, 0, 0, true)

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
			SetRows(0, 0, 0, 0, 0).SetColumns(0, 0, 0, 0, 0, 0, 0, 0, 0).
			AddItem(t.AuthForm(), 2, 3, 1, 2, 0, 0, true)
	})
}

func (t *Tui) ShowLoginForm() {
	t.Preview.Clear()
	t.queueUpdateDraw(func() {
		t.Preview.
			SetRows(0, 0, 0, 0, 0).SetColumns(0, 0, 0, 0, 0, 0, 0, 0, 0).
			AddItem(t.LoginForm(), 2, 3, 1, 2, 0, 0, true)
	})
}

func (t *Tui) ShowCreateTextDataForm() {
	t.Preview.Clear()

	t.queueUpdateDraw(func() {
		t.Preview.
			SetRows(0, 0, 0, 0, 0, 0).SetColumns(0, 0, 0, 0, 0, 0, 0).
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

func (t *Tui) ShowCreateCardsDataForm() {
	t.Preview.Clear()

	t.queueUpdateDraw(func() {
		t.Preview.
			SetRows(0, 0, 0, 0, 0, 0, 0, 0).SetColumns(0, 0, 0, 0, 0, 0, 0).
			AddItem(t.CardsdataForm(&pb.CardsdataEntry{}, "Create"), 3, 2, 2, 2, 0, 0, true)
	})
}

func (t *Tui) ShowCreateBinDataForm() {
	t.Preview.Clear()

	t.queueUpdateDraw(func() {
		t.Preview.
			SetRows(0, 0, 0, 0, 0, 0).SetColumns(0, 0, 0, 0, 0, 0, 0).
			AddItem(t.BindataForm(&pb.BindataEntry{}, "Create"), 2, 1, 2, 4, 0, 0, true)
	})
}

func (t *Tui) ShowCreateDataModal() {
	t.Preview.Clear()

	t.queueUpdateDraw(func() {
		t.Preview.AddItem(t.CreateDataModal(), 2, 2, 1, 1, 0, 0, true)
	})
}
