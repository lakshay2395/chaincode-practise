#!/bin/bash
# sleep 60
cd /opt/channel-artifacts/channel-artifacts/schooloneorg/
peer channel create -c schooloneorg -f /opt/channel-artifacts/channel-artifacts/schooloneorg/schooloneorg.tx -o orderer.nic.in:7050
peer channel join -b /opt/channel-artifacts/channel-artifacts/schooloneorg/schooloneorg.block
# cd /opt/channel-artifacts/channel-artifacts/schooltwoorg/
# peer channel create -c schooltwoorg -f /opt/channel-artifacts/channel-artifacts/schooltwoorg/schooltwoorg.tx -o orderer.nic.in:7050
# peer channel join -b /opt/channel-artifacts/channel-artifacts/schooltwoorg/schooltwoorg.block
peer chaincode install -p github.com/chaincode -n scc -v 0
peer chaincode instantiate -C schooloneorg -c '{"Args":["init"]}' -n scc -v 0 -P "AND('SchoolOneOrg.Admin')"
# peer chaincode instantiate -C schooltwoorg -c '{"Args":["init"]}' -n scc -v 0
# peer chaincode invoke -n scc -C schooloneorg -c '{"Args":["AddStudentDetails","Lakshay Bhambri","23-06-1995","MALE","Vinay Bhambri","Vimmi Bhambri"]}'
# f7a2a460-a6e1-480c-ae9f-a5fff1a75f89
# peer chaincode invoke -n scc -C schooloneorg -c '{"Args":["GetStudentDetails","73da98c8-7101-4e06-907a-b5033f077491"]}'

# peer channel fetch config ./c2.block -c schooloneorg -o orderer.nic.in:7050
# peer channel join -b ./c2.block