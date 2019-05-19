#!/bin/sh
cd ~/fabric-samples/basic-network
./stop.sh -d true
docker rm $(docker ps -a -q)
