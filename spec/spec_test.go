package spec_test

import (
	"path/filepath"
	"runtime"

	. "github.com/genesis-community/testkit/v2/testing"
	. "github.com/onsi/ginkgo/v2"
)

var _ = BeforeSuite(func() {
	_, filename, _, _ := runtime.Caller(0)
	KitDir, _ = filepath.Abs(filepath.Join(filepath.Dir(filename), "../"))
})

var _ = Describe("RabbitMQ Kit", func() {

	Describe("RabbitMQ", func() {
		Test(Environment{
			Name:        "base",
			CloudConfig: "aws",
			CPI:         "aws",
			// Focus:       true,
		})
		Test(Environment{
			Name:        "base-allparams",
			CloudConfig: "aws",
			CPI:         "aws",
			Exodus:      "cf",
			// Focus:       true,
		})
		Test(Environment{
			Name:        "broker",
			CloudConfig: "aws",
			CPI:         "aws",
			Exodus:      "cf",
			// Focus:       true,
		})
		Test(Environment{
			Name:        "metrics-emitter",
			CloudConfig: "aws",
			CPI:         "aws",
			Exodus:      "cf",
			// Focus:       true,
		})
		Test(Environment{
			Name:        "external-rmq-lb",
			CloudConfig: "aws",
			CPI:         "aws",
			Exodus:      "cf",
			// Focus:       true,
		})
		Test(Environment{
			Name:        "route-registrar",
			CloudConfig: "aws",
			CPI:         "aws",
			Exodus:      "cf",
			// Focus:       true,
		})
		Test(Environment{
			Name:        "nats-tls",
			CloudConfig: "aws",
			CPI:         "aws",
			Exodus:      "cf",
			// Focus:       true,
		})
		Test(Environment{
			Name:        "no-tls",
			CloudConfig: "aws",
			CPI:         "aws",
			Exodus:      "cf",
			// Focus:       true,
		})
		Test(Environment{
			Name:        "stomp",
			CloudConfig: "aws",
			CPI:         "aws",
			Exodus:      "cf",
			// Focus:       true,
		})

	})
})
