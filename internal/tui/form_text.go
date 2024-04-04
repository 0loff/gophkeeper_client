package tui

import (
	"context"

	"github.com/rivo/tview"

	pb "github.com/0loff/gophkeeper_server/proto"
)

func (t *Tui) TextdataForm(data *pb.TextdataEntry, btn string) *tview.Form {
	f := tview.NewForm().
		AddInputField("Metainfo", data.Metainfo, 87, nil, nil).
		AddTextArea("Text", data.Text, 87, 14, 0, nil).
		AddButton(btn, func() {
			switch btn {
			case "Create":
				t.TextdataCreation()
			case "Update":
				t.TextdataUpdating(data.ID)
			}
		}).
		SetButtonsAlign(tview.AlignCenter)

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
