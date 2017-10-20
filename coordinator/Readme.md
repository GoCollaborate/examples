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
go run main.go -mode=cdnt
```
The restful services api is available at: 
```
http://localhost:8080/services
```

### Extended Basic Types
#### Parameter
```
{
	"type": <ParamType>,
	"description": <String>,
	"constraints": [<Constraint>],
	"required": <Bool>
}
```
#### Card
```
{
	"ip": <String>,
	"port": <Int>,
	"alive": <Bool>,
	"api": <String>
}
```
#### Mode

| Name        | Value           | Description  |
|:------------- |:-------------| :----- |
| ClbtModeNormal      | 0 | Working mode |
| ClbtModeOnlyRegister | 1 | Collaborator only registers service, service is not accessible until it has been changed to ClbtModeNormal |
| ClbtModeOnlySubscribe | 2 | Subscriber only subscribes to collaborator service at coordinator, no service redirection will be provided |
| LBModeRandom | 3 | Assign tasks as per weighted probability |
| LBModeLeastActive | 4 | Assign tasks to least active responders |
| LBModeRoundRobin | 5 | Assign tasks sequentially based on the order of collaborator |
| LBModeIPHash | 6 | Assign tasks based on the hash value of subscriber IP |

#### ParamType

| Name        | Value           | 
|:------------- |:-------------|
| ArgTypeInteger | "integer" |
| ArgTypeNumber | "number" | 
| ArgTypeString | "string" |
| ArgTypeObject | "object" | 
| ArgTypeBoolean | "boolean" | 
| ArgTypeNull | "null" | 
| ArgTypeArray | "array" | 

#### Constraint
```
{
	"key": <ConstraintKey>,
	"value": <Interface{}>
}
```
##### ConstraintKey
| Name        | Value           | 
|:------------- |:-------------|
| ConstraintTypeMax | "maximum" |
| ConstraintTypeMin | "minimum" | 
| ConstraintTypeXMin | "exclusiveMinimum" |
| ConstraintTypeXMax | "exclusiveMaximum" | 
| ConstraintTypeUniqueItems | "uniqueItems" | 
| ConstraintTypeMaxProperties | "maxProperties" | 
| ConstraintTypeMinProperties | "minProperties" | 
| ConstraintTypeMaxLength | "maxLength" |
| ConstraintTypeMinLength | "minLength" | 
| ConstraintTypePattern | "pattern" |
| ConstraintTypeMaxItems | "maxItems" | 
| ConstraintTypeMinItems | "minItems" | 
| ConstraintTypeEnum | "enum" | 
| ConstraintTypeAllOf | "allOf" | 
| ConstraintTypeAnyOf | "anyOf" | 
| ConstraintTypeOneOf | "oneOf" | 

#### Registry
```
{
	"cards": [<Card>]
}
```

#### Subscription
```
{
	"token": <String>
}
```

#### Heartbeat
```
{
	"card": <Card>
}
```

### Create Service
- **POST**: `/services`

- **Headers**:
	- Content-Type: `application/json`

**Body** (required):
```
{
	"data": [{
		"type": "service",
		"attributes": {
			"description": <String>,
			"parameters": [<Parameter>],
			"registers": [<Card>],
			"subscribers": [<String>],
			"mode": <Mode>,
			"load_balance_mode": <Mode>,
			"dependencies": [<String>],
			"version": <String>,
			"platform_version": <String>
		}
	}]
}
```

### Get Services
#### Bulk
- **GET**: `/services`
#### Single
- **GET**: `/services/{serviceid}`

### Register Service
The service provider should register their endpoint as per they expose for external access.

- **POST**: `/services/{serviceid}/registry`

- **Headers**:
	- Content-Type: `application/json`

**Body** (required):
```
{
	"data": [{
		"id": <String/"">,
		"type": "registry",
		"attributes": <Registry>
	}]
}
```

### Subscribe Service
The service consumer should subscribe their interest as per they request for.

- **POST**: `/services/{serviceid}/subscription`

- **Headers**:
	- Content-Type: `application/json`

**Body**:
```
{
	"data": [{
		"id": <String/"">,
		"type": "subscription",
		"attributes": <Subscription>
	}]
}
```

### Deregister Service
The service provider should deregister their endpoint as per they terminate the provision.

#### Single Deletion
- **DELETE**: `/services/{serviceid}/registry/{ip}/{port}`

#### Bulk Deletion
- **DELETE**: `/services/{serviceid}/registry`

### Unsubscribe Service
The service consumer should unsubscribe their usage as per they terminate the dependencies.

#### Single Deletion
- **DELETE**: `/services/{serviceid}/subscription/{token}`

#### Bulk Deletion
- **DELETE**: `/services/{serviceid}/subscription`

### Delete Service
Delete a service if it is no longer required.

- **DELETE**: `/services/{serviceid}`

### Heart Beat
Send server heartbeats to Coordinator.

- **POST**: `/services/heartbeat`

- **Headers**:
	- Content-Type: `application/json`

**Body**:
```
{
	"data": [{
		"id": <String/"">,
		"type": "subscription",
		"attributes": <Heartbeat>
	}]
}
```

### Query Service
Client launch request to call a service

- **GET**: `/query/{srvid}/{token}`
