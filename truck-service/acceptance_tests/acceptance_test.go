package acceptance_tests

import (
	"context"
	"fleet-backend/common"
	proto2 "fleet-backend/common/proto"
	"fleet-backend/truck-service/proto"
	"github.com/corbym/gocrest/has"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/google/uuid"
	client2 "github.com/micro/go-micro/client"
	"testing"
)

func truckServiceClient() proto.TruckService {
	return proto.NewTruckService("truck-service", client2.NewClient(common.UseConsul))
}

func TestTruck(t *testing.T) {
	////Test SignUp
	createTruckResponse, err := truckServiceClient().CreateTruck(context.Background(), &proto.Truck{
		LicensePlate:   "qwewfwef",
		ClockedInUser:  "asdf",
		Miles:          324.34,
		FleetCompanyId: uuid.New().String(),
		CorporationId:  uuid.New().String(),
		RegionId:       uuid.New().String(),
		DistrictId:     uuid.New().String(),
		LocationId:     uuid.New().String(),
	})

	then.AssertThat(t, err, is.Nil())

	truck := createTruckResponse.Truck
	then.AssertThat(t, truck.Id == "", is.False())
	then.AssertThat(t, truck.LicensePlate, is.EqualTo("qwewfwef"))
	then.AssertThat(t, truck.ClockedInUser, is.EqualTo("asdf"))
	then.AssertThat(t, truck.Miles == 324.34, is.True())
	then.AssertThat(t, truck.FleetCompanyId == "", is.False())
	then.AssertThat(t, truck.CorporationId == "", is.False())
	then.AssertThat(t, truck.RegionId == "", is.False())
	then.AssertThat(t, truck.DistrictId == "", is.False())
	then.AssertThat(t, truck.LocationId == "", is.False())

	////Test UpdateTruck
	updateTruckReponse, err := truckServiceClient().UpdateTruck(context.Background(), &proto.Truck{
		Id:             truck.Id,
		LicensePlate:   "JP",
		ClockedInUser:  "",
		Miles:          546.0,
		FleetCompanyId: truck.FleetCompanyId,
		CorporationId:  truck.CorporationId,
		RegionId:       truck.RegionId,
		DistrictId:     truck.DistrictId,
		LocationId:     truck.LocationId,
	})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, updateTruckReponse.Truck.LicensePlate, is.EqualTo("JP"))
	then.AssertThat(t, updateTruckReponse.Truck.ClockedInUser, is.EqualTo(""))
	then.AssertThat(t, updateTruckReponse.Truck.Miles == 546.0, is.True())
	then.AssertThat(t, updateTruckReponse.Truck.FleetCompanyId, is.EqualTo(truck.FleetCompanyId))
	then.AssertThat(t, updateTruckReponse.Truck.CorporationId, is.EqualTo(truck.CorporationId))
	then.AssertThat(t, updateTruckReponse.Truck.RegionId, is.EqualTo(truck.RegionId))
	then.AssertThat(t, updateTruckReponse.Truck.DistrictId, is.EqualTo(truck.DistrictId))
	then.AssertThat(t, updateTruckReponse.Truck.LocationId, is.EqualTo(truck.LocationId))

	////Test GetTruckById
	getTruckByIdResponse, err := truckServiceClient().GetTruckById(context.Background(), &proto2.IdRequest{Id: updateTruckReponse.Truck.Id})
	then.AssertThat(t, err, is.Nil())
	truck = getTruckByIdResponse.Truck
	then.AssertThat(t, truck.LicensePlate, is.EqualTo("JP"))
	then.AssertThat(t, truck.ClockedInUser, is.EqualTo(""))

	////Test GetDriversByFleetCompanyId
	trucksByFleetCompanyIdResponse, err := truckServiceClient().GetAllTrucksByFleetCompanyId(context.Background(), &proto2.IdRequest{Id: truck.FleetCompanyId})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, trucksByFleetCompanyIdResponse.Trucks, has.Length(1))

	////Test ClockIn
	clockInResponse, err := truckServiceClient().ClockIn(context.Background(), &proto.ClockOperation{
		DriverId: "mukulgarg",
		TruckId:  truck.Id,
	})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, clockInResponse.Truck.Id, is.EqualTo(truck.Id))
	then.AssertThat(t, clockInResponse.Truck.ClockedInUser, is.EqualTo("mukulgarg"))

	////Test ClockOut
	clockOutResponse, err := truckServiceClient().ClockOut(context.Background(), &proto.ClockOperation{
		DriverId: "mukulgarg",
		TruckId:  truck.Id,
	})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, clockOutResponse.Truck.Id, is.EqualTo(truck.Id))
	then.AssertThat(t, clockOutResponse.Truck.ClockedInUser, is.EqualTo(""))
}
