package impl

import (
	"context"
	"fleet-backend/common"
	"fleet-backend/truck-service/proto"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const dbName = "customer"
const truckCollection = "truck"

//const clientCollection = "clients"
//const corporationCollection = "corporation"
//const regionCollection = "region"
//const districtCollection = "district"
//const locationCollection = "location"

type TruckServiceRepository struct {
	client *mgo.Session
	dbName string
}

func NewTruckServiceRepository() (*TruckServiceRepository, error) {
	if session, err := common.ConnectMongo(); err != nil {
		return nil, err
	} else {
		return &TruckServiceRepository{
			client: session,
		}, err
	}
}

func (c *TruckServiceRepository) truckCollection() *mgo.Collection {
	return c.client.DB(dbName).C(truckCollection)
}

func (c *TruckServiceRepository) Close() {
	c.client.Close()
}

func (c *TruckServiceRepository) AddTruck(ctx context.Context, truck *proto.Truck) error {
	if err := c.truckCollection().Insert(truck); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *TruckServiceRepository) UpdateTruckFromId(ctx context.Context, truck *proto.Truck) (*proto.Truck, error) {
	colQuerier := bson.M{"id": truck.Id}
	change := bson.M{"$set": bson.M{"licenseplate": truck.LicensePlate, "clockedinuser": truck.ClockedInUser, "miles": truck.Miles, "fleetcompanyid": truck.FleetCompanyId, "corporationid": truck.CorporationId, "regionid": truck.RegionId, "districtid": truck.DistrictId, "locationid": truck.LocationId}}
	err := c.truckCollection().Update(colQuerier, change)
	if err != nil {
		return nil, err
	} else {
		return truck, nil
	}
}

func (c *TruckServiceRepository) GetTruckById(ctx context.Context, truckId string) (*proto.Truck, error) {
	truck := &proto.Truck{}
	if err := c.truckCollection().Find(bson.M{"id": truckId}).One(truck); err != nil {
		return nil, err
	} else {
		return truck, nil
	}
}

func (c *TruckServiceRepository) GetAllTrucksByFleetCompanyId(ctx context.Context, fleetCompanyId string) ([]*proto.Truck, error) {
	var trucks []*proto.Truck
	if err := c.truckCollection().Find(bson.M{"fleetcompanyid": fleetCompanyId}).All(&trucks); err != nil {
		return nil, err
	} else {
		return trucks, nil
	}
}
