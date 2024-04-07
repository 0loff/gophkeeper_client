package tui

import (
	"context"

	"github.com/rivo/tview"

	pb "github.com/0loff/gophkeeper_server/proto"
)

func (t *Tui) TextdataForm(data *pb.TextdataEntry, btn string) *tview.Form {
	f := tview.NewForm().
		AddInputField("Metainfo", data.Metainfo, 87, nil, nil).
		AddTextArea("Text", data.Text, 87, 12, 0, nil).
		AddButton(btn, func() {
			switch btn {
			case "Create":
				t.TextdataCreation()
			case "Update":
				t.TextdataUpdating(data.ID)
			}
		}).
		SetButtonsAlign(tview.AlignCenter)

	if data.ID != 0 {
		f.AddButton("Delete", func() {
			t.TextdataDelete(data.ID)
		})
	}

	f.AddButton("Close", func() {
		t.Preview.Clear()
		t.ShowCreateDataModal()
	})

	f.SetBorder(true)
	f.SetTitle("Create Text Data Form")

	t.Form = f
	return f
}

func (t *Tui) TextdataCreation() {
	textField := t.Form.GetFormItemByLabel("Text")
	metainfoField := t.Form.GetFormItemByLabel("Metainfo")

	t.App.StatusCh <- t.App.Requestor.NewRequest(context.Background(), t.App.JWT).CreateTextData(
		context.Background(),
		textField.(*tview.TextArea).GetText(),
		metainfoField.(*tview.InputField).GetText(),
	)

	t.Preview.Clear()
	t.ShowCreateTextDataForm()
}

func (t *Tui) TextdataUpdating(id int64) {
	textField := t.Form.GetFormItemByLabel("Text")
	metainfoField := t.Form.GetFormItemByLabel("Metainfo")

	t.App.StatusCh <- t.App.Requestor.NewRequest(context.Background(), t.App.JWT).UpdateTextData(
		context.Background(),
		int(id),
		textField.(*tview.TextArea).GetText(),
		metainfoField.(*tview.InputField).GetText(),
	)
}

func (t *Tui) TextdataDelete(id int64) {
	t.App.StatusCh <- t.App.Requestor.NewRequest(context.Background(), t.App.JWT).DeleteTextData(
		context.Background(),
		int(id),
	)

	t.Preview.Clear()
	t.ShowCreateDataModal()
}
