# Introduction

**Currently not deployed!**

Pseudo-real IoT data provided by a free REST API.

https://fakeiotapi.xyz

This project was made mostly for me to learn deployment of a full stack application, so it can be that
some things in backend and frontend code are not implemented in the best way possible.

# Implemented endpoints

```
GET /api/v1/farms -> Get all farms
GET /api/v1/farms/{farm-id} -> Get a single farm

GET /api/v1/employees -> Get all employees
GET /api/v1/employees/{employee-id} -> Get a single employee

GET /api/v1/fields -> Get all fields
GET /api/v1/fields/{field-id} -> Get a single field

GET /api/v1/sensors -> Get all sensors
GET /api/v1/sensors/{sensor-id} -> Get a single sensor

GET /api/v1/farms/{farm-id}/employees -> Get all employees for a farm
GET /api/v1/farms/{farm-id}/fields -> Get all fields for a farm
GET /api/v1/farms/{farm-id}/sensors -> Get all sensors for a farm

GET /api/v1/fields/{field-id}/sensors -> Get all sensors for a field

GET /api/v1/sensors/{sensor-id}/temperature/current -> Get current temperature for a sensor
GET /api/v1/sensors/{sensor-id}/humidity/current -> Get current humidity for a sensor
GET /api/v1/sensors/{sensor-id}/pressure/current -> Get current pressure for a sensor

GET /api/v1/sensors/{sensor-id}/temperature/history -> Get temperature history for a sensor
GET /api/v1/sensors/{sensor-id}/humidity/history -> Get humidity history for a sensor
GET /api/v1/sensors/{sensor-id}/pressure/history -> Get pressure history for a sensor
```

# Development environment

## Prerequisites

In parentheses are versions that I use.

go (1.20.1)  
npm (9.5.1)  
node (v19.8.1)  
Docker (20.10.22)  
Docker Compose (v2.15.1)

## Start Application

`git clone https://github.com/defilippomattia/fakeiotapi`  
`cd fakeiotapi`  
`docker-compose up --force-recreate "db-container" -d`  
`cd backend`  
`go run main.go`  
`cd frontend`  
`npm install`  
`in frontend/src/components/DocsPage/EndpointsTable.js set envDev to true -> NOT PRETTY`  
`npm run start`  
`http://localhost:3000`

## Start Tests

`cd backend`  
`go test ./...`

# Namecheap

Domain is on Namecahep, if Hosted Zone is created/recreated, change nameservers to the ones from AWS Route 53.

# Certificates

TODO: automate this

This is need only once or when certificate expires.

Instructions from: https://www.youtube.com/watch?v=5wzs-pcDQ3k&t=334s&ab_channel=TRY-ITDEV

1. Run command:  
   sudo -u ubuntu certbot certonly \  
    --manual \\  
    --preferred-challenges=dns \\  
    --email defilippomattia@gmail.com \\  
    --server https://acme-v02.api.letsencrypt.org/directory \\  
    --work-dir=/home/ubuntu/mycerts --config-dir=/home/ubuntu/mycerts --logs-dir=/home/ubuntu/mycerts \\  
    --agree-tos \\  
    -d fakeiotapi.xyz

2. During process (prompt with key value pair for TXT record will be shown - don't close it/click Y just yet):
3. Go to AWS Route 53, select fakeiotapi.xyz, click Create record, for name use key (usually \_acme-challenge), for type use TXT, for value use the value from the prompt
4. Go to https://www.digwebinterface.com/, for hostanme use key from prompt (usually \_acme-challenge.fakeiotapi.xyz), for type use TXT, click Dig and check if value is the one you entered
5. Wait for few minutes even if dig is OK, it can be that not every NS is propagated yet, If it is, click Enter in the terminal prompt
6. Certificates will be created in /home/ubuntu/mycerts/archive/fakeiotapi.xyz (if renewing, number will be increased by 1)
7. GitHub Actions will copy them to the server

# Improvements

- fakeiotapi-cicd.yml: sleep 10s wait for pg to start (better solution to check if pg is up)
- automate certificate creation and renewal
- deploy to AWS ECS
