package common

import (
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry/consul"
)

var UseConsul = func(options *client.Options) {
	options.Registry = consul.NewRegistry()
}
