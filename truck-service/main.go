package main

import (
	"fleet-backend/truck-service/impl"
	"fleet-backend/truck-service/proto"
	"github.com/micro/go-micro"
)

func main() {

	service := micro.NewService(
		micro.Name("truck-service"),
	)

	repo, err := impl.NewTruckServiceRepository()
	if err != nil {
		panic(err)
	}
	defer repo.Close()

	svc := impl.Service{
		Repository: repo,
	}

	handler := impl.Handler{
		Service: svc,
	}

	service.Init()

	proto.RegisterTruckServiceHandler(
		service.Server(),
		handler)

	if err := service.Run(); err != nil {
		panic(err)
	}

}
