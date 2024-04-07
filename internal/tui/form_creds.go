package tui

import (
	"context"

	pb "github.com/0loff/gophkeeper_server/proto"
	"github.com/rivo/tview"
)

func (t *Tui) CredsdataForm(data *pb.CredsdataEntry, btn string) *tview.Form {
	f := tview.NewForm().
		AddInputField("Username", data.Username, 55, nil, nil).
		AddInputField("Password", data.Password, 55, nil, nil).
		AddInputField("Metainfo", data.Metainfo, 55, nil, nil).
		AddButton(btn, func() {
			switch btn {
			case "Create":
				t.CredsdataCreation()
			case "Update":
				t.CredsdataUpdating(data.ID)
			}
		}).
		SetButtonsAlign(tview.AlignCenter)

	if data.ID != 0 {
		f.AddButton("Delete", func() {
			t.CredsdataDelete(data.ID)
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

func (t *Tui) CredsdataCreation() {
	usernameField := t.Form.GetFormItemByLabel("Username")
	passwordField := t.Form.GetFormItemByLabel("Password")
	metainfoField := t.Form.GetFormItemByLabel("Metainfo")

	t.App.StatusCh <- t.App.Requestor.NewRequest(context.Background(), t.App.JWT).CreateCredsData(
		context.Background(),
		usernameField.(*tview.InputField).GetText(),
		passwordField.(*tview.InputField).GetText(),
		metainfoField.(*tview.InputField).GetText(),
	)
}

func (t *Tui) CredsdataUpdating(id int64) {
	usernameField := t.Form.GetFormItemByLabel("Username")
	passwordField := t.Form.GetFormItemByLabel("Password")
	metainfoField := t.Form.GetFormItemByLabel("Metainfo")

	t.App.StatusCh <- t.App.Requestor.NewRequest(context.Background(), t.App.JWT).UpdateCredsData(
		context.Background(),
		int(id),
		usernameField.(*tview.InputField).GetText(),
		passwordField.(*tview.InputField).GetText(),
		metainfoField.(*tview.InputField).GetText(),
	)
}

func (t *Tui) CredsdataDelete(id int64) {
	t.App.StatusCh <- t.App.Requestor.NewRequest(context.Background(), t.App.JWT).DeleteCredsData(
		context.Background(),
		int(id),
	)

	t.Preview.Clear()
	t.ShowCreateDataModal()
}
