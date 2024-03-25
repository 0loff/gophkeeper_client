// App package - this is a package for creating an application instance
package app

import (
	"context"

	pb "github.com/0loff/gophkeeper_server/proto"

	"github.com/0loff/gophkeeper_client/internal/requestor"
)

type App struct {
	Request *requestor.Requestor

	TokenCh chan string
	JWT     string

	Textdata *pb.TextdataEntriesResponse
}

func NewApp() *App {
	app := &App{
		Request: requestor.NewRequestor(),
		TokenCh: make(chan string),
	}

	go app.UserDataWorker(app.TokenCh)

	return app
}

func (a *App) UserDataWorker(ch chan string) {
	for {
		select {
		case a.JWT = <-ch:
			a.Textdata = a.Request.GetTextData(context.Background(), a.JWT)
		}
	}
}

func (a *App) LoadData() {

}

func (a *App) GetTextData() {

}
