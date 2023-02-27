#!/bin/bash

# Run Project
docker-compose up --build -d
sleep 5

# Query results
curl 127.0.0.1:8080/secret | jq '.secretCode'
curl 127.0.0.1:8080/health | jq '.'

# Destroy
docker-compose down