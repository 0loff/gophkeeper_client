package tui

import (
	"context"

	pb "github.com/0loff/gophkeeper_server/proto"
	"github.com/rivo/tview"
)

func (t *Tui) CardsdataForm(data *pb.CardsdataEntry, btn string) *tview.Form {
	f := tview.NewForm().
		AddInputField("Pan", data.Pan, 30, nil, nil).
		AddInputField("Expiry date", data.Expiry, 10, nil, nil).
		AddInputField("Card holder", data.Holder, 30, nil, nil).
		AddInputField("Metainfo", data.Metainfo, 30, nil, nil).
		AddButton(btn, func() {
			switch btn {
			case "Create":
				t.CardsdataCreation()
			case "Update":
				t.CardsdataUpdating(data.ID)
			}
		}).
		SetButtonsAlign(tview.AlignCenter)

	if data.ID != 0 {
		f.AddButton("Delete", func() {
			t.CardsdataDelete(data.ID)
		})
	}

	f.AddButton("Close", func() {
		t.Preview.Clear()
		t.ShowCreateDataModal()
	})

	f.SetBorder(true)
	f.SetTitle("Create Credentials Data Form")

	t.Form = f
	return f
}

func (t *Tui) CardsdataCreation() {
	panField := t.Form.GetFormItemByLabel("Pan")
	expiryField := t.Form.GetFormItemByLabel("Expiry date")
	holderField := t.Form.GetFormItemByLabel("Card holder")
	metainfoField := t.Form.GetFormItemByLabel("Metainfo")

	t.App.StatusCh <- t.App.Requestor.NewRequest(context.Background(), t.App.JWT).CreateCardData(
		context.Background(),
		panField.(*tview.InputField).GetText(),
		expiryField.(*tview.InputField).GetText(),
		holderField.(*tview.InputField).GetText(),
		metainfoField.(*tview.InputField).GetText(),
	)

	t.Preview.Clear()
	t.ShowCreateCardsDataForm()
}

func (t *Tui) CardsdataUpdating(id int64) {
	panField := t.Form.GetFormItemByLabel("Pan")
	expiryField := t.Form.GetFormItemByLabel("Expiry date")
	holderField := t.Form.GetFormItemByLabel("Card holder")
	metainfoField := t.Form.GetFormItemByLabel("Metainfo")

	t.App.StatusCh <- t.App.Requestor.NewRequest(context.Background(), t.App.JWT).UpdateCardsData(
		context.Background(),
		int(id),
		panField.(*tview.InputField).GetText(),
		expiryField.(*tview.InputField).GetText(),
		holderField.(*tview.InputField).GetText(),
		metainfoField.(*tview.InputField).GetText(),
	)
}

func (t *Tui) CardsdataDelete(id int64) {
	t.App.StatusCh <- t.App.Requestor.NewRequest(context.Background(), t.App.JWT).DeleteCardsData(
		context.Background(),
		int(id),
	)

	t.Preview.Clear()
	t.ShowCreateDataModal()
}
