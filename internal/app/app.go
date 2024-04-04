// App package - this is a package for creating an application instance
package app

import (
	pb "github.com/0loff/gophkeeper_server/proto"

	"github.com/0loff/gophkeeper_client/internal/requestor"
)

type App struct {
	Requestor *requestor.Requestor

	StatusCh chan string
	TokenCh  chan string
	JWT      string

	Textdata *pb.TextdataEntriesResponse
}

func NewApp() *App {
	app := &App{
		Requestor: requestor.NewRequestor(),
		StatusCh:  make(chan string),
		TokenCh:   make(chan string),
	}

	return app
}

func (a *App) LoadData() {

}

func (a *App) GetTextData() {

}
