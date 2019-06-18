package main

import (
	"fleet-backend/customer-service/impl"
	"fleet-backend/customer-service/proto"
	"github.com/micro/go-micro"
)

func main() {

	service := micro.NewService(
		micro.Name("customer-service"),
	)

	repo, err := impl.NewCustomerServiceRepository()
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

	proto.RegisterCustomerServiceHandler(
		service.Server(),
		handler)

	if err := service.Run(); err != nil {
		panic(err)
	}
}
