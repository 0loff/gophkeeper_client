package tui

import "github.com/rivo/tview"

func (t *Tui) CreateDataModal() *tview.Modal {
	return tview.NewModal().
		SetText("What data do you want to keep").
		AddButtons([]string{
			"Create Text Data",
			"Create Credentials Data",
			"Create Card Data",
			"Create Binary Data",
		}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			switch buttonLabel {
			case "Create Text Data":
				t.ShowCreateTextDataForm()
			case "Create Credentials Data":
				t.ShowCreateCredsDataForm()
			case "Create Card Data":
				t.ShowCreateCardsDataForm()
			case "Create Binary Data":
				t.ShowCreateBinDataForm()
			}
		})
}
