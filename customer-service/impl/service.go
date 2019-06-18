package impl

import (
	"context"
	"fleet-backend/customer-service/proto"
	"github.com/google/uuid"
)

type Service struct {
	Repository *CustomerServiceRepository
}

func (s Service) SignUp(ctx context.Context, request *proto.SignUpRequest) (*proto.Driver, error) {

	client := &proto.FleetCompany{
		Id:   uuid.New().String(),
		Name: request.FleetCompanyName,
	}

	if err := s.Repository.AddClient(ctx, client); err != nil {
		return nil, err
	}

	driver := &proto.Driver{
		Name:           request.Name,
		Id:             uuid.New().String(),
		Email:          request.Email,
		Password:       request.Password,
		FleetCompanyId: client.Id,
	}

	if err := s.Repository.AddDriver(ctx, driver); err != nil {
		return nil, err
	} else {
		return driver, nil
	}
}

func (s Service) CreateDriver(ctx context.Context, request *proto.Driver) (*proto.Driver, error) {
	driver := &proto.Driver{
		Name:           request.Name,
		Id:             uuid.New().String(),
		Email:          request.Email,
		Password:       request.Password,
		FleetCompanyId: request.FleetCompanyId,
	}

	if err := s.Repository.AddDriver(ctx, driver); err != nil {
		return nil, err
	} else {
		return driver, nil
	}
}

func (s Service) GetDriverById(ctx context.Context, driverId string) (*proto.Driver, error) {
	if driver, err := s.Repository.GetDriverById(ctx, driverId); err != nil {
		return nil, err
	} else {
		return driver, nil
	}
}

func (s Service) GetAllDriversByClientId(ctx context.Context, clientId string) ([]*proto.Driver, error) {
	if drivers, err := s.Repository.GetDriversByClientId(ctx, clientId); err != nil {
		return nil, err
	} else {
		return drivers, nil
	}
}

func (s Service) UpdateDriverFromId(ctx context.Context, driver *proto.Driver) (*proto.Driver, error) {
	if driver, err := s.Repository.UpdateDriverFromId(ctx, driver); err != nil {
		return nil, err
	} else {
		return driver, nil
	}
}

func (s Service) CreateCorporation(ctx context.Context, corporation *proto.Corporation) (*proto.Corporation, error) {
	fleetcompany := &proto.FleetCompany{
		Id:   uuid.New().String(),
		Name: corporation.Name,
	}

	if err := s.Repository.AddClient(ctx, fleetcompany); err != nil {
		return nil, err
	}

	corp := &proto.Corporation{
		Name:           corporation.Name,
		Id:             uuid.New().String(),
		FleetCompanyId: fleetcompany.Id,
	}

	if err := s.Repository.AddCorporation(ctx, corp); err != nil {
		return nil, err
	} else {
		return corp, nil
	}
}

func (s Service) GetAllCorporationsByFleetCompanyId(ctx context.Context, fleetCompanyId string) ([]*proto.Corporation, error) {
	if corporations, err := s.Repository.GetAllCorporationsByFleetCompanyId(ctx, fleetCompanyId); err != nil {
		return nil, err
	} else {
		return corporations, nil
	}
}

func (s Service) GetCorporationById(ctx context.Context, corporationId string) (*proto.Corporation, error) {
	if corporation, err := s.Repository.GetCorporationById(ctx, corporationId); err != nil {
		return nil, err
	} else {
		return corporation, nil
	}
}

func (s Service) CreateRegion(ctx context.Context, request *proto.Region) (*proto.Region, error) {
	region := &proto.Region{
		Id:            uuid.New().String(),
		Name:          request.Name,
		CorporationId: request.CorporationId,
	}

	if region, err := s.Repository.AddRegion(ctx, region); err != nil {
		return nil, err
	} else {
		return region, nil
	}

}

func (s Service) GetAllRegionsByCorporationId(ctx context.Context, corporationId string) ([]*proto.Region, error) {
	if regions, err := s.Repository.GetAllRegionsByCorporationId(ctx, corporationId); err != nil {
		return nil, err
	} else {
		return regions, nil
	}
}

func (s Service) GetRegionById(ctx context.Context, regionId string) (*proto.Region, error) {
	if region, err := s.Repository.GetRegionById(ctx, regionId); err != nil {
		return nil, err
	} else {
		return region, nil
	}
}

func (s Service) CreateDistrict(ctx context.Context, request *proto.District) (*proto.District, error) {
	district := &proto.District{
		Id:       uuid.New().String(),
		Name:     request.Name,
		RegionId: request.RegionId,
	}

	if district, err := s.Repository.AddDistrict(ctx, district); err != nil {
		return nil, err
	} else {
		return district, nil
	}
}

func (s Service) GetAllDistrictsByRegionId(ctx context.Context, regionId string) ([]*proto.District, error) {
	if districts, err := s.Repository.GetAllDistrictsByRegionId(ctx, regionId); err != nil {
		return nil, err
	} else {
		return districts, nil
	}
}

func (s Service) GetDistrictById(ctx context.Context, districtId string) (*proto.District, error) {
	if district, err := s.Repository.GetDistrictById(ctx, districtId); err != nil {
		return nil, err
	} else {
		return district, nil
	}
}

func (s Service) CreateLocation(ctx context.Context, request *proto.Location) (*proto.Location, error) {
	location := &proto.Location{
		Id:         uuid.New().String(),
		Name:       request.Name,
		DistrictId: request.DistrictId,
	}

	if location, err := s.Repository.AddLocation(ctx, location); err != nil {
		return nil, err
	} else {
		return location, nil
	}
}

func (s Service) GetAllLocationsByDistrictId(ctx context.Context, districtId string) ([]*proto.Location, error) {
	if locations, err := s.Repository.GetAllLocationsByDistrictId(ctx, districtId); err != nil {
		return nil, err
	} else {
		return locations, nil
	}
}

func (s Service) GetLocationById(ctx context.Context, locationId string) (*proto.Location, error) {
	if location, err := s.Repository.GetLocationById(ctx, locationId); err != nil {
		return nil, err
	} else {
		return location, nil
	}
}
