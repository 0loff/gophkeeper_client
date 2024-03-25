package main

import (
	"log"

	pb "github.com/0loff/gophkeeper_server/proto"

	"github.com/0loff/gophkeeper_client/internal/app"
	"github.com/0loff/gophkeeper_client/internal/interceptor"
	"github.com/0loff/gophkeeper_client/internal/tui"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	Run(app.NewApp())
}

func Run(app *app.App) {
	conn, err := grpc.Dial(":3200", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(interceptor.AuthInterceptor))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	app.Request.GrpcClient = pb.NewGophkeeperClient(conn)

	t := tui.NewAppView(app)
	t.Init()
}
