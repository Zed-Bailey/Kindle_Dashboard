#!/bin/sh
cd "dashboard"

echo "building application for kindle 4"
env GOARCH=arm GOARM=7 GOOS=linux go build -o dashboard
echo "built!"
echo "executable path = bin/dashboard"

mv dashboard ../bin/
