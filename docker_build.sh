#!/bin/bash

IMAGE_TAG=gcr.io/tyrionlannister-237214/dns-hammer

go build dnshammer.go
docker build -t $IMAGE_TAG:latest .
docker push $IMAGE_TAG:latest
