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

func (r Request) DeleteTextData(ctx context.Context, id int) string {
	cb, err := r.Client.TextdataDelete(r.Ctx, &pb.TextDataDeleteRequest{
		ID: int64(id),
	})
	if err != nil {
		logger.Log.Error("Cannot delete user text data", zap.Error(err))
		return "fail"
	}

	return cb.Status
}

func (r Request) GetCredsData() *pb.CredsdataEntriesResponse {

	cd, err := r.Client.CredsdataGet(r.Ctx, &emptypb.Empty{})
	if err != nil {
		logger.Log.Error("Error during receiving user credentials data", zap.Error(err))
	}

	return cd
}

func (r Request) CreateCredsData(ctx context.Context, username, password []byte, metainfo string) string {
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

func (r Request) UpdateCredsData(ctx context.Context, id int, username, password []byte, metainfo string) string {
	cb, err := r.Client.CredsdataUpdate(r.Ctx, &pb.CredsdataUpdateRequest{
		ID:       int64(id),
		Username: username,
		Password: password,
		Metainfo: metainfo,
	})
	if err != nil {
		logger.Log.Error("Cannot update user data", zap.Error(err))
		return "fail"
	}

	return cb.Status
}

func (r Request) DeleteCredsData(ctx context.Context, id int) string {
	cb, err := r.Client.CredsdataDelete(r.Ctx, &pb.CredsdataDeleteRequest{
		ID: int64(id),
	})
	if err != nil {
		logger.Log.Error("Cannot delete user credentials data", zap.Error(err))
		return "fail"
	}

	return cb.Status
}

func (r Request) GetCardsData() *pb.CardsdataEntriesResponse {

	cd, err := r.Client.CardsdataGet(r.Ctx, &emptypb.Empty{})
	if err != nil {
		logger.Log.Error("Error during receiving user cards data", zap.Error(err))
	}

	return cd
}

func (r Request) CreateCardData(ctx context.Context, pan, exp, holder []byte, metainfo string) string {
	cb, err := r.Client.CardsdataCreate(r.Ctx, &pb.CardsdataStoreRequest{
		Pan:      pan,
		Expiry:   exp,
		Holder:   holder,
		Metainfo: metainfo,
	})
	if err != nil {
		logger.Log.Error("Cannot create user card data", zap.Error(err))
		return "fail"
	}

	return cb.Status
}

func (r Request) UpdateCardsData(ctx context.Context, id int, pan, exp, holder []byte, metainfo string) string {
	cb, err := r.Client.CardsdataUpdate(r.Ctx, &pb.CardsdataUpdateRequest{
		ID:       int64(id),
		Pan:      pan,
		Expiry:   exp,
		Holder:   holder,
		Metainfo: metainfo,
	})
	if err != nil {
		logger.Log.Error("Cannot update user card data", zap.Error(err))
		return "fail"
	}

	return cb.Status
}

func (r Request) DeleteCardsData(ctx context.Context, id int) string {
	cb, err := r.Client.CardsdataDelete(r.Ctx, &pb.CardsdataDeleteRequest{
		ID: int64(id),
	})
	if err != nil {
		logger.Log.Error("Cannot delete user cards data", zap.Error(err))
		return "fail"
	}

	return cb.Status
}

func (r Request) GetBinData() *pb.BindataEntriesResponse {

	cd, err := r.Client.BindataGet(r.Ctx, &emptypb.Empty{})
	if err != nil {
		logger.Log.Error("Error during receiving user binary data", zap.Error(err))
	}

	return cd
}

func (r Request) CreateBinData(ctx context.Context, binary, metainfo string) string {
	cb, err := r.Client.BindataCreate(r.Ctx, &pb.BindataStoreRequest{
		Binary:   []byte(binary),
		Metainfo: metainfo,
	})
	if err != nil {
		logger.Log.Error("Cannot create user binary data", zap.Error(err))
		return "fail"
	}

	return cb.Status
}

func (r Request) UpdateBinData(ctx context.Context, id int, binary, metainfo string) string {
	cb, err := r.Client.BindataUpdate(r.Ctx, &pb.BindataUpdateRequest{
		ID:       int64(id),
		Binary:   []byte(binary),
		Metainfo: metainfo,
	})
	if err != nil {
		logger.Log.Error("Cannot update user binary data", zap.Error(err))
		return "fail"
	}

	return cb.Status
}

func (r Request) DeleteBinData(ctx context.Context, id int) string {
	cb, err := r.Client.BindataDelete(r.Ctx, &pb.BindataDeleteRequest{
		ID: int64(id),
	})
	if err != nil {
		logger.Log.Error("Cannot delete user binary data", zap.Error(err))
		return "fail"
	}

	return cb.Status
}
