package build_coordinator

import (
	"github.com/gavv/httpexpect"
	"net/http"
	"testing"
)

func TestGetService(t *testing.T) {
	var (
		// create httpexpect instance
		e = httpexpect.New(t, "http://localhost:8080")
	)
	// services GET
	e.GET("/services/{srvid}").
		Expect().JSON()
}

func TestGetServices(t *testing.T) {
	var (
		// create httpexpect instance
		e = httpexpect.New(t, "http://localhost:8080")
	)
	// services GET
	e.GET("/services").
		Expect().
		Status(http.StatusOK).JSON()
}

func TestPostServices(t *testing.T) {
	var (
		// create httpexpect instance
		e = httpexpect.New(t, "http://localhost:8080")
	)
	// services POST
	const (
		createServiceJSON = `{
		"data":[{
			"type":"service",
			"attributes": {
				"description": "test string",
				"parameters": [{
					"type":"string",
					"description": "test string",
					"constraints": [{
						"key": "maxLength",
						"value": 10
					}, {
						"key": "minLength",
						"value": 5
					}],
					"required": false
				},{
					"type": "integer",
	                "description": "test integer",
	                "constraints": [{
	                    "key": "maximum",
	                    "value": 1000
	                }, {
	                    "key": "minimum",
	                    "value": 100
	                }],
	                "required": true
				}, {
	                "type": "array",
	                "description": "test argument array",
	                "constraints": [{
	                    "key": "maxItems",
	                    "value": 1000
	                }, {
	                    "key": "uniqueItems",
	                    "value": true
	                }],
	                "required": true
            	}],
	            "registers": [],
	            "subscribers": [],
            	"mode": "LBModeRoundRobin",
	            "dependencies": [],
	            "version": "1.0",
	            "platform_version": "golang1.8.1"
			}
		}
	]}`
		schema = `{
			"data": "array",
			"included": "array",
			"links": "array"
		}`
	)

	repos := e.POST("/services").
		WithHeader("Content-Type", "application/json").
		WithJSON(createServiceJSON).
		Expect().
		Status(http.StatusCreated).JSON()
	repos.Schema(schema)
}

func TestAlterServices(t *testing.T) {
	var (
		// create httpexpect instance
		e = httpexpect.New(t, "http://localhost:8080")
	)
	// services PUT
	const (
		alterServiceJSON = `{
		"data":[{
			"type":"service",
			"attributes": {
				"description": "test string",
				"parameters": [],
	            "registers": [],
	            "subscribers": [],
            	"mode": "LBModeRoundRobin",
	            "dependencies": [],
	            "version": "1.0",
	            "platform_version": "golang1.8.1"
			}
		}
	]}`
		schema = `{
			"data": "array",
			"included": "array",
			"links": "array"
		}`
	)
	repos := e.PUT("/services").
		WithHeader("Content-Type", "application/json").
		WithJSON(alterServiceJSON).
		Expect().
		Status(http.StatusAccepted).JSON()
	repos.Schema(schema)
}

func TestDeleteService(t *testing.T) {
	var (
		// create httpexpect instance
		e = httpexpect.New(t, "http://localhost:8080")
	)
	// DELETE service
	e.DELETE("/services/{srvid}").
		Expect().JSON()
}
