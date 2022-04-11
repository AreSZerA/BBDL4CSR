#!/usr/bin/env bash

echo
echo "===================================================================================="
echo "  oooooooooo.    ooooo               .o      .oooooo.     .oooooo..o  ooooooooo. "
echo "  \`888'   \`Y8b   \`888'             .d88     d8P'  \`Y8b   d8P'    \`Y8  \`888   \`Y88. "
echo "   888      888   888            .d'888    888           Y88bo.        888   .d88' "
echo "   888      888   888          .d'  888    888            \`\"Y8888o.    888ooo88P'  "
echo "   888      888   888          88ooo888oo  888                \`\"Y88b   888\`88b.    "
echo "   888     d88'   888       o       888    \`88b    ooo   oo     .d8P   888  \`88b.  "
echo "  o888bood8P'    o888ooooood8      o888o    \`Y8bood8P'   8\"\"88888P'   o888o  o888o "
echo "===================================================================================="

STEP_BEG="\\033[32m"
STEP_END="\\033[0m"
COUNT_STEP=1

function print_step() {
  echo -e "\n${STEP_BEG}>>> Step ${COUNT_STEP}: $1${STEP_END}"
  ((COUNT_STEP++))
}

export PATH=${PWD}/bin:${PWD}:$PATH
export CLAYTON_UNIVERSITY_CA_PK=$(cd crypto-config/peerOrganizations/clayton-university.dl4csr.org/ca && ls *_sk)
export GARYTON_UNIVERSITY_CA_PK=$(cd crypto-config/peerOrganizations/garyton-university.dl4csr.org/ca && ls *_sk)

print_step "Clean environment"
mkdir -p ./channel-artifacts
mkdir -p ./crypto-config
sudo rm -rf ./channel-artifacts/*
sudo rm -rf ./crypto-config/*
docker-compose -f docker-compose-zookeeper.yaml down --remove-orphans

print_step "Generates certifications, genesis block, and channel configuration transactions"
cryptogen generate --config=./crypto-config.yaml
configtxgen \
  -profile TwoOrgsOrdererGenesis \
  -outputBlock ./channel-artifacts/genesis.block \
  -channelID genesischannel
configtxgen \
  -profile TwoOrgsChannel \
  -outputCreateChannelTx ./channel-artifacts/claytonuniversitychannel.tx \
  -channelID claytonuniversitychannel
configtxgen \
  -profile TwoOrgsChannel \
  -outputCreateChannelTx ./channel-artifacts/garytonuniversitychannel.tx \
  -channelID garytonuniversitychannel
export CLAYTON_UNIVERSITY_CA_PK=$(cd crypto-config/peerOrganizations/clayton-university.dl4csr.org/ca && ls *_sk)
export GARYTON_UNIVERSITY_CA_PK=$(cd crypto-config/peerOrganizations/garyton-university.dl4csr.org/ca && ls *_sk)

print_step "Start network"
docker-compose -f docker-compose-zookeeper.yaml up -d
docker-compose -f docker-compose-kafka.yaml up -d
docker-compose -f docker-compose-orderer.yaml up -d
docker-compose -f docker-compose-couchdb.yaml up -d
docker-compose -f docker-compose-peer.yaml up -d
docker-compose -f docker-compose-cli.yaml up -d
echo "Sleep 10 seconds for kafka cluster to complete booting..."
sleep 10

print_step "Create and join channels"
docker exec cli.clayton-university.dl4csr.org peer channel create \
  -o orderer1.dl4csr.org:7050 \
  -c claytonuniversitychannel \
  -f /opt/gopath/src/github.com/hyperledger/fabric/peer/channel-artifacts/claytonuniversitychannel.tx \
  --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/dl4csr.org/orderers/orderer1.dl4csr.org/msp/tlscacerts/tlsca.dl4csr.org-cert.pem
docker exec cli.garyton-university.dl4csr.org peer channel create \
  -o orderer1.dl4csr.org:7050 \
  -c garytonuniversitychannel \
  -f /opt/gopath/src/github.com/hyperledger/fabric/peer/channel-artifacts/garytonuniversitychannel.tx \
  --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/dl4csr.org/orderers/orderer1.dl4csr.org/msp/tlscacerts/tlsca.dl4csr.org-cert.pem

docker exec cli.clayton-university.dl4csr.org peer channel join \
  -b claytonuniversitychannel.block \
  --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/dl4csr.org/orderers/orderer1.dl4csr.org/msp/tlscacerts/tlsca.dl4csr.org-cert.pem
docker exec cli.garyton-university.dl4csr.org peer channel join \
  -b garytonuniversitychannel.block \
  --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/dl4csr.org/orderers/orderer1.dl4csr.org/msp/tlscacerts/tlsca.dl4csr.org-cert.pem

print_step "Install and instantiate chaincode"
docker exec cli.clayton-university.dl4csr.org peer chaincode install \
  -n dl4csr \
  -v 1.0.0 \
  -l golang \
  -p chaincode \
  --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/dl4csr.org/orderers/orderer1.dl4csr.org/msp/tlscacerts/tlsca.dl4csr.org-cert.pem
docker exec cli.garyton-university.dl4csr.org peer chaincode install \
  -n dl4csr \
  -v 1.0.0 \
  -l golang \
  -p chaincode \
  --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/dl4csr.org/orderers/orderer1.dl4csr.org/msp/tlscacerts/tlsca.dl4csr.org-cert.pem
docker exec cli.clayton-university.dl4csr.org peer chaincode instantiate \
  -o orderer1.dl4csr.org:7050 \
  -C claytonuniversitychannel \
  -n dl4csr \
  -l golang \
  -v 1.0.0 \
  -c '{"Args":["init"]}' \
  --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/dl4csr.org/orderers/orderer1.dl4csr.org/msp/tlscacerts/tlsca.dl4csr.org-cert.pem
echo "Sleep 5 seconds for chaincode to complete instantiating..."
sleep 5
docker exec cli.garyton-university.dl4csr.org peer chaincode instantiate \
  -o orderer1.dl4csr.org:7050 \
  -C garytonuniversitychannel \
  -n dl4csr \
  -l golang \
  -v 1.0.0 \
  -c '{"Args":["init"]}' \
  --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/dl4csr.org/orderers/orderer1.dl4csr.org/msp/tlscacerts/tlsca.dl4csr.org-cert.pem
echo "Sleep 5 seconds for chaincode to complete instantiating..."
sleep 5

print_step "Test chaincode invocation"
docker exec cli.clayton-university.dl4csr.org peer chaincode invoke \
  -C claytonuniversitychannel \
  -n dl4csr \
  -c '{"Args":["test"]}' \
  --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/dl4csr.org/orderers/orderer1.dl4csr.org/msp/tlscacerts/tlsca.dl4csr.org-cert.pem
docker exec cli.garyton-university.dl4csr.org peer chaincode invoke \
  -C garytonuniversitychannel \
  -n dl4csr \
  -c '{"Args":["test"]}' \
  --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/dl4csr.org/orderers/orderer1.dl4csr.org/msp/tlscacerts/tlsca.dl4csr.org-cert.pem