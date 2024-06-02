package service

import (
	"andrefsilveira/grpc-quick-start/internal/pb"
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

func (c *RegisterService) CalculateDistance(stream pb.RegisterService_CreateRegisterStreamServer) error {
	registers := &pb.Registers{}
	var checkpoint = 0.0
	var closer = false

	for {
		register, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(registers)
		}
		if err != nil {
			return err
		}
		lat, err := strconv.ParseFloat(register.Latitude, 64)
		lon, err := strconv.ParseFloat(register.Longidute, 64)
		result := haversine.Calculate(latitude, longitude, lat, lon)
		value := strconv.FormatFloat(result, 'f', -1, 64)
		if err != nil {
			return err
		}

		if result < checkpoint {
			checkpoint = result
			closer = true
		} else {
			closer = false
		}

		registers.Regiters = append(registers.Regiters, &pb.Register{
			Id:        uuid.New().String(),
			Latitude:  register.Latitude,
			Longitude: register.Longidute,
			Distance:  value,
			Closer:    closer,
		})

	}
}
