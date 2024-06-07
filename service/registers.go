package service

import (
	"andrefsilveira/grpc-quick-start/internal/pb"
	"context"
	"io"
	"strconv"

	"github.com/andrefsilveira1/go-haversine"
	"github.com/google/uuid"
)

type RegisterService struct {
	pb.UnimplementedRegisterServiceServer
}

func NewRegisterService() *RegisterService {
	return &RegisterService{}
}

var latitude = -5.8623992555733695
var longitude = -35.19111877919574

func (c *RegisterService) CreateRegister(ctx context.Context, in *pb.CreateRequest) (*pb.Register, error) {
	register := &pb.Register{}

	lat, err := strconv.ParseFloat(in.Latitude, 64)
	if err != nil {
		return nil, err
	}

	lon, err := strconv.ParseFloat(in.Longitude, 64)
	if err != nil {
		return nil, err
	}

	result := haversine.Calculate(latitude, longitude, lat, lon)
	value := strconv.FormatFloat(result, 'f', -1, 64)

	register = &pb.Register{
		Id:        uuid.New().String(),
		Latitude:  register.Latitude,
		Longitude: register.Longitude,
		Distance:  value,
	}

	return register, nil

}

func (c *RegisterService) CreateRegisterStream(stream pb.RegisterService_CreateRegisterStreamServer) error {
	registers := &pb.Registers{}

	for {
		register, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(registers)
		}

		if err != nil {
			return err
		}

		lat, err := strconv.ParseFloat(register.Latitude, 64)
		if err != nil {
			return err
		}

		lon, err := strconv.ParseFloat(register.Longitude, 64)
		if err != nil {
			return err
		}

		result := haversine.Calculate(latitude, longitude, lat, lon)
		value := strconv.FormatFloat(result, 'f', -1, 64)

		registers.Regiters = append(registers.Regiters, &pb.Register{
			Id:        uuid.New().String(),
			Latitude:  register.Latitude,
			Longitude: register.Longitude,
			Distance:  value,
		})

	}
}
func (c *RegisterService) CreateRegisterBidirectional(stream pb.RegisterService_CreateRegisterBidirectionalServer) error {
	for {
		register, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		lat, err := strconv.ParseFloat(register.Latitude, 64)
		if err != nil {
			return err
		}

		lon, err := strconv.ParseFloat(register.Longitude, 64)
		if err != nil {
			return err
		}

		result := haversine.Calculate(latitude, longitude, lat, lon)
		value := strconv.FormatFloat(result, 'f', -1, 64)

		err = stream.Send(&pb.Register{
			Id:        uuid.New().String(),
			Latitude:  register.Latitude,
			Longitude: register.Longitude,
			Distance:  value,
		})

		if err != nil {
			return err
		}
	}
}
