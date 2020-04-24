package main

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/canhdoan/shippy-service-vessel/proto/vessel"
	"github.com/micro/go-micro"
)

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

type VesselRepository struct {
	vessels []*pb.Vessel
}

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}

	return nil, errors.New("No vessel found by that spec")
}

type service struct {
	repo Repository
}

func (s *service) FindAvaliable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}

	res.Vessel = vessel
	return nil
}

func main() {
	vessels := []*pb.Vessel{
		&pb.Vessel{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
		&pb.Vessel{Id: "vessel002", Name: "Boaty McBoatface", MaxWeight: 100000, Capacity: 1500},
		&pb.Vessel{Id: "vessel003", Name: "Boaty McBoatface", MaxWeight: 300000, Capacity: 2500},
		&pb.Vessel{Id: "vessel004", Name: "Boaty McBoatface", MaxWeight: 400000, Capacity: 700},
		&pb.Vessel{Id: "vessel005", Name: "Boaty McBoatface", MaxWeight: 500000, Capacity: 800},
	}
	repo := &VesselRepository{vessels}

	sv := micro.NewService(
		micro.Name("shippy.service.vessel"),
	)
	sv.Init()

	// register implement
	pb.RegisterVesselServiceHandler(sv.Server(), &service{repo})
	if err := sv.Run(); err != nil {
		fmt.Println(err)
	}
}
