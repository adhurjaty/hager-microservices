#!/bin/bash

docker-compose down
docker-compose build
docker-compose up -d
docker-compose run -d user-service
curl -XPOST -H 'Content-Type: application/json' -d '{ "service": "go.micro.srv.user", "method": "UserService.Create", "request": { "Username": "adhurjaty", "Password": "testing123", "FirstName": "Anil", "Company": "Novacoast" } }' http://localhost:8080/rpc 