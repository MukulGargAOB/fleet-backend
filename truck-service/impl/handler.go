package impl

import (
	"context"
	proto2 "fleet-backend/common/proto"
	"fleet-backend/truck-service/proto"
)

type Handler struct {
	Service Service
}

func (h Handler) CreateTruck(ctx context.Context, req *proto.Truck, res *proto.TruckResponse) error {
	if truck, err := h.Service.CreateTruck(ctx, req); err != nil {
		return err
	} else {
		res.Truck = truck
		return nil
	}
}

func (h Handler) UpdateTruck(ctx context.Context, req *proto.Truck, res *proto.TruckResponse) error {
	if truck, err := h.Service.UpdateTruck(ctx, req); err != nil {
		return err
	} else {
		res.Truck = truck
		return nil
	}
}

func (h Handler) GetTruckById(ctx context.Context, req *proto2.IdRequest, res *proto.TruckResponse) error {
	if truck, err := h.Service.GetTruckById(ctx, req.Id); err != nil {
		return err
	} else {
		res.Truck = truck
		return nil
	}
}

func (h Handler) GetAllTrucksByFleetCompanyId(ctx context.Context, req *proto2.IdRequest, res *proto.TrucksResponse) error {
	if trucks, err := h.Service.GetAllTrucksByFleetCompanyId(ctx, req.Id); err != nil {
		return err
	} else {
		res.Trucks = trucks
		return nil
	}
}

func (h Handler) ClockIn(ctx context.Context, req *proto.ClockOperation, res *proto.TruckResponse) error {
	if truck, err := h.Service.ClockIn(ctx, req); err != nil {
		return err
	} else {
		res.Truck = truck
		return nil
	}
}

func (h Handler) ClockOut(ctx context.Context, req *proto.ClockOperation, res *proto.TruckResponse) error {
	if truck, err := h.Service.ClockOut(ctx, req); err != nil {
		return err
	} else {
		res.Truck = truck
		return nil
	}
}
