// App package - this is a package for creating an application instance
package app

import (
	pb "github.com/0loff/gophkeeper_server/proto"

	"github.com/0loff/gophkeeper_client/internal/requestor"
)

type info struct {
	Version string
	Commit  string
	Date    string
}

type App struct {
	Requestor *requestor.Requestor

	StatusCh chan string
	TokenCh  chan string
	JWT      string

	Textdata  *pb.TextdataEntriesResponse
	Credsdata *pb.CredsdataEntriesResponse
	CardsData *pb.CardsdataEntriesResponse
	Bindata   *pb.BindataEntriesResponse

	Info info
}

func NewApp() *App {
	app := &App{
		Requestor: requestor.NewRequestor(),
		StatusCh:  make(chan string),
		TokenCh:   make(chan string),
	}

	return app
}
