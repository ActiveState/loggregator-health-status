# loggregator_health_status

Small utility the allows you to identify the status of loggregator API.

First you need to use CF client to target the api:
```
cf api --skip-ssl-validation https://api.192.168.4.31.xip.io
cf login
cf app APP_NAME --guid  #skip this if you already know the app guid
7102d156-bb28-4a12-b81a-5cee5b3e7116
```
and that it.

```
$ go run main.go --appGuid=7102d156-bb28-4a12-b81a-5cee5b3e7116
```

and the expected results :

```
           status     query
           200 OK     info
    404 Not Found     v2/apps/7102d156-bb28-4a12-b81a-5cee5b3e7116/instances/0
    404 Not Found     v2/config/environment_variable_groups/running
 401 Unauthorized     v2/organizations/ebcb6f21-5e5c-494d-b3a0-df335aff36ad/spaces?inline-relations-depth=1
           200 OK     v2/service_plans
 401 Unauthorized     v2/config/feature_flags
 401 Unauthorized     v2/service_plan_visibilities
 401 Unauthorized     v2/quota_definitions
 401 Unauthorized     v2/service_brokers
 401 Unauthorized     v2/organizations
    404 Not Found     thisisExampleOF/5ab38238-bd0b-448e-aef6-7015e86536f7
```

To add more endpoints updated the ```endpoints.csv``` like below:
```
v2/app/APPGUID/instances/INSTANCE
```

Available variables for query uri:

```
	APPGUID   
	INSTANCE  
	ORGGUID   
	SPACEGUID 
```

