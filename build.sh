#! /bin/bash

#build web UI
cd ~/go/src/video_server/web
go install
cp ~/go/bin/web ~/go/bin/video_server_web_ui/web
cp -R ~/go/src/video_server/templates ~/go/bin/video_server_web_ui/