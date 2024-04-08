package tui

import (
	"context"

	"github.com/0loff/gophkeeper_client/pkg/encryptor"
	pb "github.com/0loff/gophkeeper_server/proto"
	"github.com/rivo/tview"
)

func (t *Tui) CardsdataForm(data *pb.CardsdataEntry, btn string) *tview.Form {
	var (
		pan    []byte
		exp    []byte
		holder []byte
		err    error
	)

	if data.Pan != nil {
		pan, err = encryptor.Decrypt(data.Pan, t.App.GetUserKey())
		if err != nil {
			t.showError("Cannot decrypt card pan")
			t.ShowCreateCardsDataForm()
		}
	}

	if data.Expiry != nil {
		exp, err = encryptor.Decrypt(data.Expiry, t.App.GetUserKey())
		if err != nil {
			t.showError("Cannot decrypt card expiry")
			t.ShowCreateCardsDataForm()
		}
	}

	if data.Holder != nil {
		holder, err = encryptor.Decrypt(data.Holder, t.App.GetUserKey())
		if err != nil {
			t.showError("Cannot decrypt card holder")
			t.ShowCreateCardsDataForm()
		}
	}

	f := tview.NewForm().
		AddInputField("Pan", string(pan), 30, nil, nil).
		AddInputField("Expiry date", string(exp), 10, nil, nil).
		AddInputField("Card holder", string(holder), 30, nil, nil).
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

	encPan, err := encryptor.Encrypt([]byte(panField.(*tview.InputField).GetText()), t.App.GetUserKey())
	if err != nil {
		t.showError("Cannot encrypt card pan")
		return
	}

	encExp, err := encryptor.Encrypt([]byte(expiryField.(*tview.InputField).GetText()), t.App.GetUserKey())
	if err != nil {
		t.showError("Cannot encrypt card expiry")
		return
	}

	encHolder, err := encryptor.Encrypt([]byte(holderField.(*tview.InputField).GetText()), t.App.GetUserKey())
	if err != nil {
		t.showError("Cannot encrypt card holder")
		return
	}

	t.App.StatusCh <- t.App.Requestor.NewRequest(context.Background(), t.App.JWT).CreateCardData(
		context.Background(),
		encPan,
		encExp,
		encHolder,
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

	encPan, err := encryptor.Encrypt([]byte(panField.(*tview.InputField).GetText()), t.App.GetUserKey())
	if err != nil {
		t.showError("Cannot encrypt card pan")
		return
	}

	encExp, err := encryptor.Encrypt([]byte(expiryField.(*tview.InputField).GetText()), t.App.GetUserKey())
	if err != nil {
		t.showError("Cannot encrypt card expiry")
		return
	}

	encHolder, err := encryptor.Encrypt([]byte(holderField.(*tview.InputField).GetText()), t.App.GetUserKey())
	if err != nil {
		t.showError("Cannot encrypt card holder")
		return
	}

	t.App.StatusCh <- t.App.Requestor.NewRequest(context.Background(), t.App.JWT).UpdateCardsData(
		context.Background(),
		int(id),
		encPan,
		encExp,
		encHolder,
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
