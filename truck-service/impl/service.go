package impl

import (
	"context"
	"fleet-backend/truck-service/proto"
	"github.com/google/uuid"
)

type Service struct {
	Repository *TruckServiceRepository
}

func (s Service) CreateTruck(ctx context.Context, request *proto.Truck) (*proto.Truck, error) {
	truck := &proto.Truck{
		Id:             uuid.New().String(),
		LicensePlate:   request.LicensePlate,
		ClockedInUser:  request.ClockedInUser,
		Miles:          request.Miles,
		FleetCompanyId: request.FleetCompanyId,
		CorporationId:  request.CorporationId,
		RegionId:       request.RegionId,
		DistrictId:     request.DistrictId,
		LocationId:     request.LocationId,
	}

	if err := s.Repository.AddTruck(ctx, truck); err != nil {
		return nil, err
	} else {
		return truck, nil
	}
}

func (s Service) UpdateTruck(ctx context.Context, truck *proto.Truck) (*proto.Truck, error) {
	if truck, err := s.Repository.UpdateTruckFromId(ctx, truck); err != nil {
		return nil, err
	} else {
		return truck, nil
	}
}

func (s Service) GetTruckById(ctx context.Context, truckId string) (*proto.Truck, error) {
	if truck, err := s.Repository.GetTruckById(ctx, truckId); err != nil {
		return nil, err
	} else {
		return truck, nil
	}
}

func (s Service) GetAllTrucksByFleetCompanyId(ctx context.Context, fleetCompanyId string) ([]*proto.Truck, error) {
	if trucks, err := s.Repository.GetAllTrucksByFleetCompanyId(ctx, fleetCompanyId); err != nil {
		return nil, err
	} else {
		return trucks, nil
	}
}

func (s Service) ClockIn(ctx context.Context, operation *proto.ClockOperation) (*proto.Truck, error) {
	if truck, err := s.Repository.GetTruckById(ctx, operation.TruckId); err != nil {
		return nil, err
	} else {
		if truck.ClockedInUser == "" {
			truck.ClockedInUser = operation.DriverId
			if truck, err := s.Repository.UpdateTruckFromId(ctx, truck); err != nil {
				return nil, err
			} else {
				return truck, nil
			}
		}
		return truck, nil
	}
}

func (s Service) ClockOut(ctx context.Context, operation *proto.ClockOperation) (*proto.Truck, error) {
	if truck, err := s.Repository.GetTruckById(ctx, operation.TruckId); err != nil {
		return nil, err
	} else {
		if truck.ClockedInUser == operation.DriverId {
			truck.ClockedInUser = ""
			if truck, err := s.Repository.UpdateTruckFromId(ctx, truck); err != nil {
				return nil, err
			} else {
				return truck, nil
			}
		}
		return truck, nil
	}
}
