#!/bin/bash

go build -o catalogue-service cmd/catalogue-service/*.go
./catalogue-service 
# -dbname=catalogue_service_db -dbuser=chunhou -cache=false -production=false
# migrate create -ext sql -dir internal/db/migrations create_products_table