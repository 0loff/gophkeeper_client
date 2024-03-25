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

func NewRequestor() *Requestor {
	return &Requestor{}
}

func (r *Requestor) SignUp(ctx context.Context, uname, pwd, email string) string {
	var header metadata.MD
	r.GrpcClient.UserAuth(ctx, &pb.UserAuthRequest{
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

func (r *Requestor) SignIn(ctx context.Context, email, pwd string) string {
	var header metadata.MD
	r.GrpcClient.UserLogin(ctx, &pb.UserLoginRequest{
		Email:    email,
		Password: pwd,
	}, grpc.Header(&header))

	token, ok := header["token"]
	if !ok {
		return ""
	}

	return token[0]
}

func (r *Requestor) GetTextData(ctx context.Context, token string) *pb.TextdataEntriesResponse {
	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	td, err := r.GrpcClient.TextdataGet(ctx, &emptypb.Empty{})
	if err != nil {
		logger.Log.Error("Error during receiving user text data", zap.Error(err))
	}

	return td

}
