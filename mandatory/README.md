# Leaf Mandatory

## Mandatory Object of Leaf project
Can be put and retrieved from/to context. consist of basic context data (user agent, os, request-id, request-language, device, user, etc.)

## Usage
Import the package
```go
import (
    ...
    "github.com/paulusrobin/leaf-utilities/mandatory"
    ...
)
```

To Retrieve mandatory from context:
```go
func Process(ctx context.Context) {
    ...
    mandatory := leafMandatort.FromContext(ctx)
    ...
    if mandatory.User().IsLogin() {
        ...
    } else {
        ...
    }
}
```

## Mandatory Functions:
- ```TraceID()``` - To get TraceID / RequestID, ```return string```
- ```IpAddresses()``` - To get Ip Addresses, return ```[]string```
- ```Language()``` - To request language, return ```string```
- ```Authorization()``` - To get Authorization Object, return ```leafMandatory.Authorization```
    - ```Authorization()``` - To get Full Authorization, return ```string```
    - ```Token()``` - To get Token Authorization, return ```string```
    - ```ApiKey()``` - To get ApiKey, return ```string```
    - ```ServiceID()``` - To get ServiceID, return ```string```
    - ```ServiceSecret()``` - To get ServiceSecret, return ```string```
- ```DeviceType()``` - To get Device Category Request, return ```leafMandatory.DeviceType```
    - ```ID()``` - To get DeviceType ID, return ```int```
    - ```Name()``` - To get DeviceType Name, return ```string```
    - ```Code()``` - To get DeviceType Code, return ```string```
- ```Device()``` - To get Device Data, return ```leafMandatory.Device```
    - ```AppVersion()``` - To get Application version, return ```string```
    - ```DeviceID()``` - To get Device Unique ID, return ```string```
    - ```Family()``` - To get Device Family, return ```string```
    - ```Brand()``` - To get Device Brand, return ```string```
    - ```Model()``` - To get Device Model, return ```string```
- ```UserAgent()``` - To get User Agent Data, return ```leafMandatory.UserAgent```
    - ```Value()``` - To get user agent Value, return ```string```
    - ```Family()``` - To get user agent Family, return ```string```
    - ```Major()``` - To get user agent Major version, return ```string```
    - ```Minor()``` - To get user agent Minor version, return ```string```
    - ```Patch()``` - To get user agent Patch version, return ```string```
- ```OS()``` - To get Operating System, return ```leafMandatory.OS```
    - ```Name()``` - To get OS Name, return ```string```
    - ```Version()``` - To get OS Version , return ```string```
    - ```Family()``` - To get OS Family, return ```string```
    - ```Major()``` - To get OS Major Version, return ```string```
    - ```Minor()``` - To get OS Minor Version, return ```string```
    - ```Patch()``` - To get OS Patch Version, return ```string```
    - ```PatchMinor()``` - To get OS Patch Minor Version, return ```string```
- ```User()``` - To get User Object, return ```leafMandatory.User```
    - ```ID()``` - To get User ID, return ```uint64```
    - ```Email()``` - To get User Email, return ```string```
    - ```IsLogin()``` - To get User Status is Login or not, return ```boolean```
- ```Valid()``` - To check mandatory object is valid or not, return ```boolean```
- ```IsUserLogin()``` - To get user login status, return ```boolean```
- ```IsMobileApp()``` - To get request is from mobile apps or not, return ```boolean```
- ```IsWebApp()``` - To get request is from web apps or not, return ```boolean```