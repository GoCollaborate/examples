package build_coordinator

import (
	"github.com/gavv/httpexpect"
	"net/http"
	"testing"
)

func TestRegisterService(t *testing.T) {
	var (
		// create httpexpect instance
		e = httpexpect.New(t, "http://localhost:8080")
	)
	// services POST
	e.POST("/services/{srvid}/registry").
		Expect().
		Status(http.StatusOK).JSON()
}

func TestSubscribeService(t *testing.T) {
	var (
		// create httpexpect instance
		e = httpexpect.New(t, "http://localhost:8080")
	)
	// services POST
	e.POST("/services/{srvid}/subscription").
		Expect().
		Status(http.StatusOK).JSON()
}

func TestDeRegisterService(t *testing.T) {
	var (
		// create httpexpect instance
		e = httpexpect.New(t, "http://localhost:8080")
	)
	// DELETE registry
	e.DELETE("/services/{srvid}/registry/single/{ip}/{port}").
		Expect().JSON()
}

func TestUnSubscribeService(t *testing.T) {
	var (
		// create httpexpect instance
		e = httpexpect.New(t, "http://localhost:8080")
	)
	// DELETE subscription
	e.DELETE("/services/{srvid}/subscription/single/{token}").
		Expect().JSON()
}
func TestBulkDeRegisterService(t *testing.T) {
	var (
		// create httpexpect instance
		e = httpexpect.New(t, "http://localhost:8080")
	)
	// DELETE registries
	e.DELETE("/services/{srvid}/registry").
		Expect().JSON()
}

func TestBulkUnSubscribeService(t *testing.T) {
	var (
		// create httpexpect instance
		e = httpexpect.New(t, "http://localhost:8080")
	)
	// DELETE subscriptions
	e.DELETE("/services/{srvid}/subscription").
		Expect().JSON()
}
