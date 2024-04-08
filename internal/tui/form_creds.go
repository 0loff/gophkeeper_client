package tui

import (
	"context"

	"github.com/0loff/gophkeeper_client/pkg/encryptor"
	pb "github.com/0loff/gophkeeper_server/proto"
	"github.com/rivo/tview"
)

func (t *Tui) CredsdataForm(data *pb.CredsdataEntry, btn string) *tview.Form {
	var uname []byte
	var pwd []byte
	var err error

	if data.Username != nil {
		uname, err = encryptor.Decrypt(data.Username, t.App.GetUserKey())
		if err != nil {
			t.showError("Cannot decrypt credentials username")
			t.ShowCreateCredsDataForm()
		}
	}

	if data.Password != nil {
		pwd, err = encryptor.Decrypt(data.Password, t.App.GetUserKey())
		if err != nil {
			t.showError("Cannot decrypt credentials password")
			t.ShowCreateCredsDataForm()
		}
	}

	f := tview.NewForm().
		AddInputField("Username", string(uname), 55, nil, nil).
		AddInputField("Password", string(pwd), 55, nil, nil).
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

	encUsername, err := encryptor.Encrypt([]byte(usernameField.(*tview.InputField).GetText()), t.App.GetUserKey())
	if err != nil {
		t.showError("Cannot encrypt username")
		return
	}

	encPassword, err := encryptor.Encrypt([]byte(passwordField.(*tview.InputField).GetText()), t.App.GetUserKey())
	if err != nil {
		t.showError("Cannot encrypt password")
		return
	}

	t.App.StatusCh <- t.App.Requestor.NewRequest(context.Background(), t.App.JWT).CreateCredsData(
		context.Background(),
		encUsername,
		encPassword,
		metainfoField.(*tview.InputField).GetText(),
	)
}

func (t *Tui) CredsdataUpdating(id int64) {
	usernameField := t.Form.GetFormItemByLabel("Username")
	passwordField := t.Form.GetFormItemByLabel("Password")
	metainfoField := t.Form.GetFormItemByLabel("Metainfo")

	encUname, err := encryptor.Encrypt([]byte(usernameField.(*tview.InputField).GetText()), t.App.GetUserKey())
	if err != nil {
		t.showError("Cannot encrypt username")
		return
	}

	encPwd, err := encryptor.Encrypt([]byte(passwordField.(*tview.InputField).GetText()), t.App.GetUserKey())
	if err != nil {
		t.showError("Cannot encrypt password")
		return
	}

	t.App.StatusCh <- t.App.Requestor.NewRequest(context.Background(), t.App.JWT).UpdateCredsData(
		context.Background(),
		int(id),
		encUname,
		encPwd,
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
