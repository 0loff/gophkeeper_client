// App package - this is a package for creating an application instance
package app

import (
	pb "github.com/0loff/gophkeeper_server/proto"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/0loff/gophkeeper_client/internal/logger"
	"github.com/0loff/gophkeeper_client/internal/requestor"
)

type info struct {
	Version string
	Commit  string
	Date    string
}

type Claims struct {
	jwt.RegisteredClaims
	UserID uuid.UUID
	Key    []byte
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

func (a *App) ParseToken(authtoken string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(authtoken, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte("secretkey"), nil
	})

	if err != nil {
		logger.Log.Error("The value of the authentication token could not be parsed", zap.Error(err))
		return nil, err
	}

	if !token.Valid {
		logger.Log.Error("Invalid auth token param", zap.Error(err))
		return nil, err
	}

	return claims, nil
}

func (a *App) GetUserKey() []byte {
	claims, err := a.ParseToken(a.JWT)
	if err != nil {
		panic(err)
	}

	return claims.Key
}
