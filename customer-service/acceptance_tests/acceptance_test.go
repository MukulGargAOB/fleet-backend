package acceptance_tests

import (
	"context"
	"fleet-backend/common"
	proto2 "fleet-backend/common/proto"
	"fleet-backend/customer-service/proto"
	"github.com/corbym/gocrest/has"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	client2 "github.com/micro/go-micro/client"
	"testing"
)

func TestDriverSignUp(t *testing.T) {
	//Test SignUp
	response, err := userServiceClient().SignUp(context.Background(), &proto.SignUpRequest{
		Name:             "Mukul",
		Password:         "12345",
		Email:            "mukul.garg@allonblock.com",
		FleetCompanyName: "ABC",
	})

	then.AssertThat(t, err, is.Nil())

	driver := response.Driver
	then.AssertThat(t, driver.Id == "", is.False())
	then.AssertThat(t, driver.FleetCompanyId == "", is.False())

	//GetDriverById
	response, err = userServiceClient().GetDriverById(context.Background(), &proto2.IdRequest{Id: driver.Id})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, response.Driver.Name, is.EqualTo("Mukul"))
	then.AssertThat(t, response.Driver.Email, is.EqualTo("mukul.garg@allonblock.com"))

	//GetDriversByFleetCompanyId
	usersResponse, err := userServiceClient().GetDriversByFleetCompanyId(context.Background(), &proto2.IdRequest{Id: response.Driver.FleetCompanyId})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, usersResponse.Drivers, has.Length(1))

	//Test CreateDriver
	createResponse, err := userServiceClient().CreateDriver(context.Background(), &proto.Driver{
		Name:           "Mukul Garg",
		Password:       "qwerty",
		Email:          "mukul@allonblock.com",
		FleetCompanyId: response.Driver.FleetCompanyId,
	})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, createResponse.Driver.Id == "", is.False())
	then.AssertThat(t, createResponse.Driver.FleetCompanyId == "", is.False())

	//UpdateDriver
	updateDriverReponse, err := userServiceClient().UpdateDriver(context.Background(), &proto.Driver{
		Name:           "Aman Garg",
		Password:       "qwerty",
		Email:          "aman.garg@allonblock.com",
		Id:             response.Driver.Id,
		FleetCompanyId: response.Driver.FleetCompanyId,
	})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, updateDriverReponse.Driver.Name, is.EqualTo("Aman Garg"))
	then.AssertThat(t, updateDriverReponse.Driver.Password, is.EqualTo("qwerty"))
	then.AssertThat(t, updateDriverReponse.Driver.Email, is.EqualTo("aman.garg@allonblock.com"))
}

func userServiceClient() proto.CustomerService {
	return proto.NewCustomerService("customer-service", client2.NewClient(common.UseConsul))
}

func TestCorporation(t *testing.T) {
	//CreateCorporation
	response, err := userServiceClient().CreateCorporation(context.Background(), &proto.Corporation{
		Name: "Mukul Co.",
	})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, response.Corporation.Id == "", is.False())
	then.AssertThat(t, response.Corporation.FleetCompanyId == "", is.False())
	then.AssertThat(t, response.Corporation.Name, is.EqualTo("Mukul Co."))

	//GetAllCorporationsByFleetCompanyId
	corporationsResponse, err := userServiceClient().GetAllCorporationsByFleetCompanyId(context.Background(), &proto2.IdRequest{Id: response.Corporation.FleetCompanyId})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, corporationsResponse.Corporations, has.Length(1))

	//GetCorporationById
	corpResponse, err := userServiceClient().GetCorporationById(context.Background(), &proto2.IdRequest{
		Id: response.Corporation.Id,
	})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, corpResponse.Corporation.Name, is.EqualTo("Mukul Co."))

	//CreateRegion
	regionResponse, err := userServiceClient().CreateRegion(context.Background(), &proto.Region{
		Name:          "Nadbai",
		CorporationId: response.Corporation.Id,
	})

	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, regionResponse.Region.CorporationId == "", is.False())

	//GetAllRegionsByCorporationId
	allRegionResponse, err := userServiceClient().GetAllRegionsByCorporationId(context.Background(), &proto2.IdRequest{Id: response.Corporation.Id})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, allRegionResponse.Regions, has.Length(1))

	//GetRegionById
	regResponse, err := userServiceClient().GetRegionById(context.Background(), &proto2.IdRequest{Id: regionResponse.Region.Id})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, regResponse.Region.Name, is.EqualTo("Nadbai"))

	//AddDistrict
	districtResponse, err := userServiceClient().CreateDistrict(context.Background(), &proto.District{
		Name:     "Nagar Road",
		RegionId: regionResponse.Region.Id,
	})

	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, districtResponse.District.RegionId == "", is.False())

	//GetAllDistrictsByRegionId
	allDistrictResponse, err := userServiceClient().GetAllDistrictsByRegionId(context.Background(), &proto2.IdRequest{Id: regionResponse.Region.Id})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, allDistrictResponse.Districts, has.Length(1))

	//GetDistrictById
	distResponse, err := userServiceClient().GetDistrictById(context.Background(), &proto2.IdRequest{Id: districtResponse.District.Id})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, distResponse.District.Name, is.EqualTo("Nagar Road"))

	//AddLocation
	locationResponse, err := userServiceClient().CreateLocation(context.Background(), &proto.Location{
		Name:       "Gupta House",
		DistrictId: districtResponse.District.Id,
	})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, locationResponse.Location.DistrictId == "", is.False())

	//GetAllLocationsByDistrictId
	allLocationResponse, err := userServiceClient().GetAllLocationsByDistrictId(context.Background(), &proto2.IdRequest{Id: districtResponse.District.Id})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, allLocationResponse.Locations, has.Length(1))

	//GetLocationById
	locResponse, err := userServiceClient().GetLocationById(context.Background(), &proto2.IdRequest{Id: locationResponse.Location.Id})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, locResponse.Location.Name, is.EqualTo("Gupta House"))
}
