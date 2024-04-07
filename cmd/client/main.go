package main

import (
	"log"

	pb "github.com/0loff/gophkeeper_server/proto"

	"github.com/0loff/gophkeeper_client/internal/app"
	"github.com/0loff/gophkeeper_client/internal/tui"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	buildVersion string = "N/A"
	buildDate    string = "N/A"
	buildCommit  string = "N/A"
)

func main() {
	Run(app.NewApp())
}

func Run(app *app.App) {

	app.Info.Version = buildVersion
	app.Info.Date = buildDate
	app.Info.Commit = buildCommit

	conn, err := grpc.Dial(":3211", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	app.Requestor.GrpcClient = pb.NewGophkeeperClient(conn)

	t := tui.NewAppView(app)
	t.Init()
}
