#! /bin/bash

# Build web and other services

cd ~/go/src/video_server/api
env GOOS=linux GOARCH=amd64 go build -o ../bin/api

cd ~/go/src/video_server/scheduler
env GOOS=linux GOARCH=amd64 go build -o ../bin/scheduler

cd ~/go/src/video_server/streamserver
env GOOS=linux GOARCH=amd64 go build -o ../bin/streamserver

cd ~/go/src/video_server/web
env GOOS=linux GOARCH=amd64 go build -o ../bin/web