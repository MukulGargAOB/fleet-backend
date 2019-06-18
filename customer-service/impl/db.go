package impl

import (
	"context"
	"fleet-backend/common"
	"fleet-backend/customer-service/proto"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const dbName = "customer"
const driverCollection = "drivers"
const clientCollection = "clients"
const corporationCollection = "corporation"
const regionCollection = "regions"
const districtCollection = "districts"
const locationCollection = "locations"

type CustomerServiceRepository struct {
	client *mgo.Session
	dbName string
}

func NewCustomerServiceRepository() (*CustomerServiceRepository, error) {
	if session, err := common.ConnectMongo(); err != nil {
		return nil, err
	} else {
		return &CustomerServiceRepository{
			client: session,
		}, err
	}
}

func (c *CustomerServiceRepository) driverCollection() *mgo.Collection {
	return c.client.DB(dbName).C(driverCollection)
}

func (c *CustomerServiceRepository) clientCollection() *mgo.Collection {
	return c.client.DB(dbName).C(clientCollection)
}

func (c *CustomerServiceRepository) corporationCollection() *mgo.Collection {
	return c.client.DB(dbName).C(corporationCollection)
}

func (c *CustomerServiceRepository) regionCollection() *mgo.Collection {
	return c.client.DB(dbName).C(regionCollection)
}

func (c *CustomerServiceRepository) districtCollection() *mgo.Collection {
	return c.client.DB(dbName).C(districtCollection)
}

func (c *CustomerServiceRepository) locationCollection() *mgo.Collection {
	return c.client.DB(dbName).C(locationCollection)
}

func (c *CustomerServiceRepository) Close() {
	c.client.Close()
}

func (c *CustomerServiceRepository) AddDriver(ctx context.Context, driver *proto.Driver) error {
	if err := c.driverCollection().Insert(driver); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *CustomerServiceRepository) AddClient(ctx context.Context, client *proto.FleetCompany) error {
	if err := c.clientCollection().Insert(client); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *CustomerServiceRepository) GetDriverById(ctx context.Context, id string) (*proto.Driver, error) {
	driver := &proto.Driver{}
	if err := c.driverCollection().Find(bson.M{"id": id}).One(driver); err != nil {
		return nil, err
	} else {
		return driver, nil
	}
}

func (c *CustomerServiceRepository) GetDriversByClientId(ctx context.Context, clientId string) ([]*proto.Driver, error) {
	var drivers []*proto.Driver
	if err := c.driverCollection().Find(bson.M{"fleetcompanyid": clientId}).All(&drivers); err != nil {
		return nil, err
	} else {
		return drivers, nil
	}
}

func (c *CustomerServiceRepository) UpdateDriverFromId(ctx context.Context, driver *proto.Driver) (*proto.Driver, error) {
	colQuerier := bson.M{"id": driver.Id}
	change := bson.M{"$set": bson.M{"name": driver.Name, "email": driver.Email, "password": driver.Password, "fleetcompanyid": driver.FleetCompanyId}}
	err := c.driverCollection().Update(colQuerier, change)
	if err != nil {
		return nil, err
	} else {
		return driver, nil
	}
}

func (c *CustomerServiceRepository) AddCorporation(ctx context.Context, corporation *proto.Corporation) error {
	if err := c.corporationCollection().Insert(corporation); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *CustomerServiceRepository) GetAllCorporationsByFleetCompanyId(ctx context.Context, fleetCompanyId string) ([]*proto.Corporation, error) {
	var corporations []*proto.Corporation
	if err := c.corporationCollection().Find(bson.M{"fleetcompanyid": fleetCompanyId}).All(&corporations); err != nil {
		return nil, err
	} else {
		return corporations, nil
	}
}

func (c *CustomerServiceRepository) GetCorporationById(ctx context.Context, corporationId string) (*proto.Corporation, error) {
	corporation := &proto.Corporation{}
	if err := c.corporationCollection().Find(bson.M{"id": corporationId}).One(corporation); err != nil {
		return nil, err
	} else {
		return corporation, nil
	}
}

func (c *CustomerServiceRepository) AddRegion(ctx context.Context, region *proto.Region) (*proto.Region, error) {
	if err := c.regionCollection().Insert(region); err != nil {
		return nil, err
	} else {
		return region, nil
	}
}

func (c *CustomerServiceRepository) GetAllRegionsByCorporationId(ctx context.Context, corporationId string) ([]*proto.Region, error) {
	var regions []*proto.Region
	if err := c.regionCollection().Find(bson.M{"corporationid": corporationId}).All(&regions); err != nil {
		return nil, err
	} else {
		return regions, nil
	}

}

func (c *CustomerServiceRepository) GetRegionById(ctx context.Context, regionId string) (*proto.Region, error) {
	var region *proto.Region
	if err := c.regionCollection().Find(bson.M{"id": regionId}).One(&region); err != nil {
		return nil, err
	} else {
		return region, nil
	}
}

func (c *CustomerServiceRepository) AddDistrict(ctx context.Context, district *proto.District) (*proto.District, error) {
	if err := c.districtCollection().Insert(district); err != nil {
		return nil, err
	} else {
		return district, nil
	}
}

func (c *CustomerServiceRepository) GetAllDistrictsByRegionId(ctx context.Context, regionId string) ([]*proto.District, error) {
	var districts []*proto.District
	if err := c.districtCollection().Find(bson.M{"regionid": regionId}).All(&districts); err != nil {
		return nil, err
	} else {
		return districts, nil
	}
}

func (c *CustomerServiceRepository) GetDistrictById(ctx context.Context, districtId string) (*proto.District, error) {
	var district *proto.District
	if err := c.districtCollection().Find(bson.M{"id": districtId}).One(&district); err != nil {
		return nil, err
	} else {
		return district, nil
	}
}

func (c *CustomerServiceRepository) AddLocation(ctx context.Context, location *proto.Location) (*proto.Location, error) {
	if err := c.locationCollection().Insert(location); err != nil {
		return nil, err
	} else {
		return location, nil
	}
}

func (c *CustomerServiceRepository) GetAllLocationsByDistrictId(ctx context.Context, districtId string) ([]*proto.Location, error) {
	var locations []*proto.Location
	if err := c.locationCollection().Find(bson.M{"districtid": districtId}).All(&locations); err != nil {
		return nil, err
	} else {
		return locations, nil
	}
}

func (c *CustomerServiceRepository) GetLocationById(ctx context.Context, locationId string) (*proto.Location, error) {
	location := &proto.Location{}
	if err := c.locationCollection().Find(bson.M{"id": locationId}).One(&location); err != nil {
		return nil, err
	} else {
		return location, nil
	}
}
