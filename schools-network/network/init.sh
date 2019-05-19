#!/bin/sh
cryptogen generate --config=./crypto-config.yaml
sudo configtxgen -profile NICOrdererGenesis -outputBlock ./channel-artifacts/genesis.block
sudo configtxgen -profile SchoolOneOrgChannel -outputCreateChannelTx ./channel-artifacts/schooloneorg.tx -channelID schooloneorg
sudo configtxgen -profile SchoolTwoOrgChannel -outputCreateChannelTx ./channel-artifacts/schooltwoorg.tx -channelID schooltwoorg
sudo configtxgen -profile SchoolOneOrgChannel -outputAnchorPeersUpdate ./channel-artifacts/SchoolOneOrgCBSEAnchors.tx -channelID schooloneorg -asOrg CBSEOrg
sudo configtxgen -profile SchoolTwoOrgChannel -outputAnchorPeersUpdate ./channel-artifacts/SchoolTwoOrgCBSEAnchors.tx -channelID schooltwoorg -asOrg CBSEOrg
sudo configtxgen -profile SchoolOneOrgChannel -outputAnchorPeersUpdate ./channel-artifacts/SchoolOneOrgAnchors.tx -channelID schooloneorg -asOrg SchoolOneOrg
sudo configtxgen -profile SchoolTwoOrgChannel -outputAnchorPeersUpdate ./channel-artifacts/SchoolTwoOrgAnchors.tx -channelID schooltwoorg -asOrg SchoolTwoOrg