#!/bin/bash

# fail when any one command fails
set -e

IMAGE_TAG=gcr.io/tyrionlannister-237214/dns-hammer

env GOOS=linux GOARCH=amd64 go build dnshammer.go
docker build -t $IMAGE_TAG:latest .
docker push $IMAGE_TAG:latest
