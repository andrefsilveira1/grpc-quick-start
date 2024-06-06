package service

import (
	"context"
	"io"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	"andrefsilveira/grpc-quick-start/internal/pb"

	"github.com/stretchr/testify/assert"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterRegisterServiceServer(s, NewRegisterService())
	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestCreateRegisterStream(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewRegisterServiceClient(conn)

	stream, err := client.CreateRegisterStream(ctx)
	if err != nil {
		t.Fatalf("CreateRegisterStream failed: %v", err)
	}

	// Send a few register requests
	registers := []struct {
		Latitude  string
		Longitude string
	}{
		{"-5.8624", "-35.1911"},
		{"-5.8625", "-35.1912"},
		{"-5.8626", "-35.1913"},
	}

	for _, reg := range registers {
		if err := stream.Send(&pb.CreateRequest{
			Latitude:  reg.Latitude,
			Longitude: reg.Longitude,
		}); err != nil {
			t.Fatalf("Failed to send a request: %v", err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		t.Fatalf("Failed to receive response: %v", err)
	}

	assert.Equal(t, len(registers), len(reply.Registers))
	for _, reg := range reply.Registers {
		assert.NotEmpty(t, reg.Id)
		assert.NotEmpty(t, reg.Distance)
	}
}

func TestCreateRegisterBidirectional(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewRegisterServiceClient(conn)

	stream, err := client.CreateRegisterBidirectional(ctx)
	if err != nil {
		t.Fatalf("CreateRegisterBidirectional failed: %v", err)
	}

	// Send a few register requests and receive responses
	registers := []struct {
		Latitude  string
		Longitude string
	}{
		{"-5.8624", "-35.1911"},
		{"-5.8625", "-35.1912"},
		{"-5.8626", "-35.1913"},
	}

	for _, reg := range registers {
		if err := stream.Send(&pb.CreateRequest{
			Latitude:  reg.Latitude,
			Longitude: reg.Longitude,
		}); err != nil {
			t.Fatalf("Failed to send a request: %v", err)
		}

		reply, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatalf("Failed to receive a response: %v", err)
		}

		assert.NotEmpty(t, reply.Id)
		assert.NotEmpty(t, reply.Distance)
	}
	stream.CloseSend()
}
