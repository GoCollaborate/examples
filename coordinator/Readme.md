# GoCollaborate
## Example for Coordinator Mode
### Entry
```go
package main

import (
	"github.com/GoCollaborate"
)

func main() {
	collaborate.Run()
}

```
### Run
Running GoCollaborate server in Coordinator mode is very simple, you just go:
```sh
go run main.go -svrmode=cdnt
```
The restful services api is available at: 
```
http://localhost:8080/services
```
### Create Service
**POST**: `/services`

**Content-Type**: `application/json`

**Body**:
```
{
	"data": [{
		"type": "service",
		"attributes": {
			"description": "test description",
			"parameters": [{
				"type": "string",
				"description": "test argument string",
				"constraints": [{
					"key": "maxLength",
					"value": 10
				}, {
					"key": "minLength",
					"value": 5
				}],
				"required": false
			}, {
				"type": "integer",
				"description": "test argument integer",
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
			"mode": "RPCServerModeNormal",
			"load_balance_mode": "RPCServerModeRandomLoadBalance",
			"dependencies": [],
			"version": "1.0",
			"platform_version": "golang1.8.1"
		}
	}]
}
```

### Get Services
#### Bulk
**GET**: `/services`
#### Single
**GET**: `/services/{serviceid}`

### Register Service
The service provider should register their endpoint as per they expose for external access.

**POST**: `/services/{serviceid}/registry`

**Content-Type**: "application/json"

**Body**:
```
{
	"data": [{
		"id": "",
		"type": "registry",
		"attributes": {
			"Cards": [{
				"ip": "192.168.0.1",
				"port": 4444,
				"alive": true
			}]
		}
	}]
}
```

### Subscribe Service
The service consumer should subscribe their interest as per they request for.

**POST**: `/services/{serviceid}/subscription`

**Content-Type**: "application/json"

**Body**:
```
{
	"data": [{
		"id": "",
		"type": "subscription"
	}]
}
```

### Deregister Service
The service provider should deregister their endpoint as per they terminate the provision.

#### Single Deletion
**DELETE**: `/services/{serviceid}/registry/{ip}/{port}`

**Content-Type**: "application/json"

#### Bulk Deletion
**DELETE**: `/services/{serviceid}/registry`

**Content-Type**: "application/json"

### Unsubscribe Service
The service consumer should unsubscribe their usage as per they terminate the dependencies.

#### Single Deletion
**DELETE**: `/services/{serviceid}/subscription/{token}`

**Content-Type**: "application/json"

#### Bulk Deletion
**DELETE**: `/services/{serviceid}/subscription`

**Content-Type**: "application/json"

### Delete Service
Delete a service if it is no longer required.

**DELETE**: `/services/{serviceid}`

**Content-Type**: "application/json"
