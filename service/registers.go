package service

import (
	"andrefsilveira/grpc-quick-start/internal/pb"
	"io"
)

type RegisterService struct {
	pb.UnimplementedRegisterServiceServer
}

func NewRegisterService() *RegisterService {
	return &RegisterService{}
}

func (c *RegisterService) CalculateDistance(stream pb.RegisterService_CreateRegisterStreamServer) error {
	registers := &pb.Registers{}
	for {
		register, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(registers)
		}
		if err != nil {
			return err
		}

		// distanceReult, err := haversine.distance(register.Latitude, register.Longidute)
	}
}
