package tui

import (
	"github.com/0loff/gophkeeper_client/pkg/encryptor"
	pb "github.com/0loff/gophkeeper_server/proto"
)

func (t *Tui) TextdataSelected(index int, mainText string, secondaryText string, shortcut rune) {
	var currentData *pb.TextdataEntry

	for _, data := range t.App.Textdata.Data {
		if mainText == data.Metainfo && secondaryText == data.Text {
			currentData = &pb.TextdataEntry{
				ID:       data.ID,
				Text:     data.Text,
				Metainfo: data.Metainfo,
			}
		}
	}

	t.viewTextData(currentData)
}

func (t *Tui) CredsdataSelected(index int, mainText string, secondaryText string, shortcut rune) {
	var currentData *pb.CredsdataEntry

	for _, data := range t.App.Credsdata.Data {
		uname, err := encryptor.Decrypt(data.Username, t.App.GetUserKey())
		if err != nil {
			panic(err)
		}

		if mainText == data.Metainfo && secondaryText == string(uname) {
			currentData = &pb.CredsdataEntry{
				ID:       data.ID,
				Username: data.Username,
				Password: data.Password,
				Metainfo: data.Metainfo,
			}
		}
	}

	t.viewCredsData(currentData)
}

func (t *Tui) CardsdataSelected(index int, mainText string, secondaryText string, shortcut rune) {
	var currentData *pb.CardsdataEntry

	for _, data := range t.App.CardsData.Data {
		if mainText == data.Metainfo && secondaryText == data.Pan {
			currentData = &pb.CardsdataEntry{
				ID:       data.ID,
				Pan:      data.Pan,
				Expiry:   data.Expiry,
				Holder:   data.Holder,
				Metainfo: data.Metainfo,
			}
		}
	}

	t.viewCardsData(currentData)
}

func (t *Tui) BindataSelected(index int, mainText string, secondaryText string, shortcut rune) {
	var currentData *pb.BindataEntry

	for _, data := range t.App.Bindata.Data {
		if mainText == data.Metainfo && secondaryText == string(data.Binary) {
			currentData = &pb.BindataEntry{
				ID:       data.ID,
				Binary:   data.Binary,
				Metainfo: data.Metainfo,
			}
		}
	}

	t.viewBinData(currentData)
}
