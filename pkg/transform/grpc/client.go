package grpc

import (
	"context"
	"os"
	"time"

	age_pb "github.com/mike955/zebra/api/age"
	bank_pb "github.com/mike955/zebra/api/bank"
	cellphone_pb "github.com/mike955/zebra/api/cellphone"
	email_pb "github.com/mike955/zebra/api/email"
	flake_pb "github.com/mike955/zebra/api/flake"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var gRPCClientMap = map[string]interface{}{}

func NewFlakeRpc(flakeAddr string) (client flake_pb.FlakeServiceClient, err error) {
	if gRPCClientMap["flake"] == nil {
		if os.Getenv("FlAKE_ADDR") != "" {
			flakeAddr = os.Getenv("FlAKE_ADDR")
		}
		conn, err := grpc.DialContext(context.Background(), flakeAddr, grpc.WithUnaryInterceptor(CllentUnaryInterceptor), grpc.WithStreamInterceptor(CllentStreamInterceptor), grpc.WithInsecure())
		if err != nil {
			return nil, err
		}
		gRPCClientMap["flake"] = flake_pb.NewFlakeServiceClient(conn)
	}
	client = gRPCClientMap["flake"].(flake_pb.FlakeServiceClient)
	return
}

func NewAgeRpc(ageAddr string) (client age_pb.AgeServiceClient, err error) {
	if gRPCClientMap["age"] == nil {
		if os.Getenv("AGE_ADDR") != "" {
			ageAddr = os.Getenv("AGE_ADDR")
		}
		conn, err := grpc.DialContext(context.Background(), ageAddr, grpc.WithUnaryInterceptor(CllentUnaryInterceptor), grpc.WithStreamInterceptor(CllentStreamInterceptor), grpc.WithInsecure())
		if err != nil {
			return nil, err
		}
		client = age_pb.NewAgeServiceClient(conn)
		gRPCClientMap["age"] = client
	}
	client = gRPCClientMap["age"].(age_pb.AgeServiceClient)
	return
}

func NewEmailRpc(emailAddr string) (client email_pb.EmailServiceClient, err error) {
	if gRPCClientMap["email"] == nil {
		if os.Getenv("EMAIL_ADDR") != "" {
			emailAddr = os.Getenv("EMAIL_ADDR")
		}
		conn, err := grpc.DialContext(context.Background(), emailAddr, grpc.WithUnaryInterceptor(CllentUnaryInterceptor), grpc.WithStreamInterceptor(CllentStreamInterceptor), grpc.WithInsecure())
		if err != nil {
			return nil, err
		}
		client = email_pb.NewEmailServiceClient(conn)
		gRPCClientMap["email"] = client
	}
	client = gRPCClientMap["email"].(email_pb.EmailServiceClient)
	return
}

func NewBankRpc(bankAddr string) (client bank_pb.BankServiceClient, err error) {
	if gRPCClientMap["bank"] == nil {
		if os.Getenv("BANK_ADDR") != "" {
			bankAddr = os.Getenv("BANK_ADDR")
		}
		conn, err := grpc.DialContext(context.Background(), bankAddr, grpc.WithUnaryInterceptor(CllentUnaryInterceptor), grpc.WithStreamInterceptor(CllentStreamInterceptor), grpc.WithInsecure())
		if err != nil {
			return nil, err
		}
		client = bank_pb.NewBankServiceClient(conn)
		gRPCClientMap["bank"] = client
	}
	client = gRPCClientMap["bank"].(bank_pb.BankServiceClient)
	return
}

func NewCellphoneRpc(cellphoneAddr string) (client cellphone_pb.CellphoneServiceClient, err error) {
	if gRPCClientMap["cellphone"] == nil {
		if os.Getenv("CELLPHONE_ADDR") != "" {
			cellphoneAddr = os.Getenv("CELLPHONE_ADDR")
		}
		conn, err := grpc.DialContext(context.Background(), cellphoneAddr, grpc.WithUnaryInterceptor(CllentUnaryInterceptor), grpc.WithStreamInterceptor(CllentStreamInterceptor), grpc.WithInsecure())
		if err != nil {
			return nil, err
		}
		client = cellphone_pb.NewCellphoneServiceClient(conn)
		gRPCClientMap["cellphone"] = client
	}
	client = gRPCClientMap["cellphone"].(cellphone_pb.CellphoneServiceClient)
	return
}

func CllentUnaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	ctx = metadata.AppendToOutgoingContext(ctx, "traceId", ctx.Value("traceId").(string))
	ctx = metadata.AppendToOutgoingContext(ctx, "x_real_ip", ctx.Value("x_real_ip").(string))
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	end := time.Now()
	logger := ctx.Value("logger").(*logrus.Entry)
	logger.Infof("info: rpc call,method: %s start time: %s, end time: %s, err: %v", method, start.Format("Basic"), end.Format(time.RFC3339), err)
	return err
}

func CllentStreamInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "traceId", ctx.Value("traceId").(string))
	ctx = metadata.AppendToOutgoingContext(ctx, "x_real_ip", ctx.Value("x_real_ip").(string))
	logger := ctx.Value("logger").(*logrus.Entry)
	start := time.Now()
	s, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}
	return newWrappedStream(ctx, s, start, logger), nil
}

type wrappedStream struct {
	ctx context.Context
	grpc.ClientStream
	start  time.Time
	logger *logrus.Entry
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	w.logger.Infof("Receive a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	w.logger.Infof("Send a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.SendMsg(m)
}

func newWrappedStream(ctx context.Context, s grpc.ClientStream, start time.Time, logger *logrus.Entry) grpc.ClientStream {
	return &wrappedStream{ctx, s, start, logger}
}
