package requestor

import (
	"context"

	"github.com/0loff/gophkeeper_client/internal/logger"
	pb "github.com/0loff/gophkeeper_server/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Requestor struct {
	GrpcClient pb.GophkeeperClient
}

type Request struct {
	Client pb.GophkeeperClient
	Ctx    context.Context
}

func NewRequestor() *Requestor {
	return &Requestor{}
}

func (r *Requestor) NewRequest(ctx context.Context, token string) Request {
	if token != "" {
		md := metadata.New(map[string]string{"token": token})
		ctx = metadata.NewOutgoingContext(ctx, md)
	}

	return Request{
		Client: r.GrpcClient,
		Ctx:    ctx,
	}
}

func (r Request) SignUp(ctx context.Context, uname, pwd, email string) string {
	var header metadata.MD
	r.Client.UserAuth(ctx, &pb.UserAuthRequest{
		Login:    uname,
		Password: pwd,
		Email:    email,
	}, grpc.Header(&header))

	token, ok := header["token"]
	if !ok {
		return ""
	}

	return token[0]
}

func (r Request) SignIn(ctx context.Context, email, pwd string) string {
	var header metadata.MD
	r.Client.UserLogin(ctx, &pb.UserLoginRequest{
		Email:    email,
		Password: pwd,
	}, grpc.Header(&header))

	token, ok := header["token"]
	if !ok {
		return ""
	}

	return token[0]
}

func (r Request) GetTextData() *pb.TextdataEntriesResponse {

	td, err := r.Client.TextdataGet(r.Ctx, &emptypb.Empty{})
	if err != nil {
		logger.Log.Error("Error during receiving user text data", zap.Error(err))
	}

	return td
}

func (r Request) CreateTextData(ctx context.Context, text, metainfo string) string {
	cb, err := r.Client.TextdataCreate(r.Ctx, &pb.TextDataStoreRequest{
		Text:     text,
		Metainfo: metainfo,
	})
	if err != nil {
		logger.Log.Error("Cannot store user data", zap.Error(err))
	}

	return cb.Status
}

func (r Request) UpdateTextData(ctx context.Context, id int, text, metainfo string) string {
	cb, err := r.Client.TextdataUpdate(r.Ctx, &pb.TextDataUpdateRequest{
		ID:       int64(id),
		Text:     text,
		Metainfo: metainfo,
	})
	if err != nil {
		logger.Log.Error("Cannot update user data", zap.Error(err))
		return "fail"
	}

	return cb.Status
}

func (r Request) CreateCredsData(ctx context.Context, username, password, metainfo string) string {
	cb, err := r.Client.CredsdataCreate(r.Ctx, &pb.CredsdataStoreRequest{
		Username: username,
		Password: password,
		Metainfo: metainfo,
	})
	if err != nil {
		logger.Log.Error("Cannot store user credentianls", zap.Error(err))
	}

	return cb.Status
}
