package tui

import (
	pb "github.com/0loff/gophkeeper_server/proto"
)

func (t *Tui) dataSelected(index int, mainText string, secondaryText string, shortcut rune) {
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
