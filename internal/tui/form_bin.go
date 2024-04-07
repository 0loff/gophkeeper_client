package tui

import (
	"context"

	pb "github.com/0loff/gophkeeper_server/proto"
	"github.com/rivo/tview"
)

func (t *Tui) BindataForm(data *pb.BindataEntry, btn string) *tview.Form {
	f := tview.NewForm().
		AddTextArea("Binary data", string(data.Binary), 87, 12, 0, nil).
		AddInputField("Metainfo", data.Metainfo, 84, nil, nil).
		AddButton(btn, func() {
			switch btn {
			case "Create":
				t.BindataCreation()
			case "Update":
				t.BindataUpdating(data.ID)
			}
		}).
		SetButtonsAlign(tview.AlignCenter)

	if data.ID != 0 {
		f.AddButton("Delete", func() {
			t.BindataDelete(data.ID)
		})
	}

	f.AddButton("Close", func() {
		t.Preview.Clear()
		t.ShowCreateDataModal()
	})

	f.SetBorder(true)
	f.SetTitle("Create Binary Data Form")

	t.Form = f
	return f
}

func (t *Tui) BindataCreation() {
	binaryField := t.Form.GetFormItemByLabel("Binary data")
	metainfoField := t.Form.GetFormItemByLabel("Metainfo")

	t.App.StatusCh <- t.App.Requestor.NewRequest(context.Background(), t.App.JWT).CreateBinData(
		context.Background(),
		binaryField.(*tview.TextArea).GetText(),
		metainfoField.(*tview.InputField).GetText(),
	)

	t.Preview.Clear()
	t.ShowCreateBinDataForm()
}

func (t *Tui) BindataUpdating(id int64) {
	binaryField := t.Form.GetFormItemByLabel("Binary data")
	metainfoField := t.Form.GetFormItemByLabel("Metainfo")

	t.App.StatusCh <- t.App.Requestor.NewRequest(context.Background(), t.App.JWT).UpdateBinData(
		context.Background(),
		int(id),
		binaryField.(*tview.TextArea).GetText(),
		metainfoField.(*tview.InputField).GetText(),
	)
}

func (t *Tui) BindataDelete(id int64) {
	t.App.StatusCh <- t.App.Requestor.NewRequest(context.Background(), t.App.JWT).DeleteBinData(
		context.Background(),
		int(id),
	)

	t.Preview.Clear()
	t.ShowCreateDataModal()
}
