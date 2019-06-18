package impl

import (
	"context"
	proto2 "fleet-backend/common/proto"
	"fleet-backend/customer-service/proto"
)

type Handler struct {
	Service Service
}

func (h Handler) SignUp(ctx context.Context, req *proto.SignUpRequest, res *proto.DriverResponse) error {
	if driver, err := h.Service.SignUp(ctx, req); err != nil {
		return err
	} else {
		res.Driver = driver
		return nil
	}
}

func (h Handler) CreateDriver(ctx context.Context, req *proto.Driver, res *proto.DriverResponse) error {
	if driver, err := h.Service.CreateDriver(ctx, req); err != nil {
		return err
	} else {
		res.Driver = driver
		return nil
	}
}

func (h Handler) UpdateDriver(ctx context.Context, req *proto.Driver, res *proto.DriverResponse) error {
	if driver, err := h.Service.UpdateDriverFromId(ctx, req); err != nil {
		return err
	} else {
		res.Driver = driver
		return nil
	}
}

func (h Handler) GetDriverById(ctx context.Context, req *proto2.IdRequest, res *proto.DriverResponse) error {
	if driver, err := h.Service.GetDriverById(ctx, req.Id); err != nil {
		return err
	} else {
		res.Driver = driver
		return nil
	}
}

func (h Handler) GetDriversByFleetCompanyId(ctx context.Context, req *proto2.IdRequest, res *proto.DriversResponse) error {
	if drivers, err := h.Service.GetAllDriversByClientId(ctx, req.Id); err != nil {
		return err
	} else {
		res.Drivers = drivers
		return nil
	}
}

func (h Handler) CreateCorporation(ctx context.Context, req *proto.Corporation, res *proto.CorporationResponse) error {
	if corporation, err := h.Service.CreateCorporation(ctx, req); err != nil {
		return err
	} else {
		res.Corporation = corporation
		return nil
	}
}

func (h Handler) GetAllCorporationsByFleetCompanyId(ctx context.Context, req *proto2.IdRequest, res *proto.CorporationsResponse) error {
	if corporations, err := h.Service.GetAllCorporationsByFleetCompanyId(ctx, req.Id); err != nil {
		return err
	} else {
		res.Corporations = corporations
		return nil
	}
}

func (h Handler) GetCorporationById(ctx context.Context, req *proto2.IdRequest, res *proto.CorporationResponse) error {
	if corporation, err := h.Service.GetCorporationById(ctx, req.Id); err != nil {
		return err
	} else {
		res.Corporation = corporation
		return nil
	}
}

func (h Handler) CreateRegion(ctx context.Context, req *proto.Region, res *proto.RegionResponse) error {
	if region, err := h.Service.CreateRegion(ctx, req); err != nil {
		return err
	} else {
		res.Region = region
		return nil
	}

}

func (h Handler) GetAllRegionsByCorporationId(ctx context.Context, req *proto2.IdRequest, res *proto.RegionsResponse) error {
	if regions, err := h.Service.GetAllRegionsByCorporationId(ctx, req.Id); err != nil {
		return err
	} else {
		res.Regions = regions
		return nil
	}

}

func (h Handler) GetRegionById(ctx context.Context, req *proto2.IdRequest, res *proto.RegionResponse) error {
	if region, err := h.Service.GetRegionById(ctx, req.Id); err != nil {
		return err
	} else {
		res.Region = region
		return nil
	}
}

func (h Handler) CreateDistrict(ctx context.Context, req *proto.District, res *proto.DistrictResponse) error {
	if district, err := h.Service.CreateDistrict(ctx, req); err != nil {
		return err
	} else {
		res.District = district
		return nil
	}

}

func (h Handler) GetAllDistrictsByRegionId(ctx context.Context, req *proto2.IdRequest, res *proto.DistrictsResponse) error {
	if districts, err := h.Service.GetAllDistrictsByRegionId(ctx, req.Id); err != nil {
		return err
	} else {
		res.Districts = districts
		return nil
	}
}

func (h Handler) GetDistrictById(ctx context.Context, req *proto2.IdRequest, res *proto.DistrictResponse) error {
	if district, err := h.Service.GetDistrictById(ctx, req.Id); err != nil {
		return err
	} else {
		res.District = district
		return nil
	}

}

func (h Handler) CreateLocation(ctx context.Context, req *proto.Location, res *proto.LocationResponse) error {
	if location, err := h.Service.CreateLocation(ctx, req); err != nil {
		return err
	} else {
		res.Location = location
		return nil
	}

}

func (h Handler) GetAllLocationsByDistrictId(ctx context.Context, req *proto2.IdRequest, res *proto.LocationsResponse) error {
	if locations, err := h.Service.GetAllLocationsByDistrictId(ctx, req.Id); err != nil {
		return err
	} else {
		res.Locations = locations
		return nil
	}
}

func (h Handler) GetLocationById(ctx context.Context, req *proto2.IdRequest, res *proto.LocationResponse) error {
	if location, err := h.Service.GetLocationById(ctx, req.Id); err != nil {
		return err
	} else {
		res.Location = location
		return nil
	}
}
