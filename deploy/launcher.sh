#!/usr/bin/env bash

function show_helps() {
  echo 'Usage: launcher.sh [OPTION]

The options are:
  -h, --help       Show this help
  -d, --deploy     Deploy or redeploy the network, blockchain will be wiped then recreated
  -r, --remove     Remove the network, volumes will be deleted and data will be wiped
  -s, --start      Start the existing network, the blockchain will be restored
  -p, --pause      Stop the existing network, the blockchain still exists in containers
  -u, --upgrade    Upgrade chaincode, use with a new version number
  -i, --invoke     Function invocation test of the chaincode, use with function name and arguments
'
}

function deploy() {
  read -r -n 1 -p 'If the network has been deployed before, data will lost, still deploy? [y/N] ' FLAG
  echo
  if [ "$FLAG" != "y" ]; then
    echo "Operation canceled"
    exit 0
  fi
  echo '====================================================================================
  oooooooooo.    ooooo               .o      .oooooo.     .oooooo..o  ooooooooo.
  `888'"'"'   `Y8b   `888'"'"'             .d88     d8P'"'"'  `Y8b   d8P'"'"'    `Y8  `888   `Y88.
   888      888   888            .d'"'"'888    888           Y88bo.        888   .d88'"'"'
   888      888   888          .d'"'"'  888    888            `"Y8888o.    888ooo88P'"'"'
   888      888   888          88ooo888oo  888                `"Y88b   888`88b.
   888     d88'"'"'   888       o       888    `88b    ooo   oo     .d8P   888  `88b.
  o888bood8P'"'"'    o888ooooood8      o888o    `Y8bood8P'"'"'   8""88888P'"'"'   o888o  o888o
===================================================================================='

  STEP_BEG="\\033[32m"
  STEP_END="\\033[0m"
  COUNT_STEP=1

  function print_step() {
    echo -e "\n${STEP_BEG}>>> Step ${COUNT_STEP}: $1${STEP_END}"
    ((COUNT_STEP++))
  }

  export PATH=${PWD}/bin:${PWD}:$PATH

  print_step "Clean environment"
  sudo rm -rf ./channel-artifacts
  sudo rm -rf ./crypto-config
  mkdir -p ./channel-artifacts
  mkdir -p ./crypto-config
  docker-compose down --remove-orphans
  docker system prune --volumes -f

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
  CLAYTON_UNIVERSITY_CA_PK=$(cd crypto-config/peerOrganizations/clayton-university.dl4csr.org/ca && ls *_sk)
  GARYTON_UNIVERSITY_CA_PK=$(cd crypto-config/peerOrganizations/garyton-university.dl4csr.org/ca && ls *_sk)
  export CLAYTON_UNIVERSITY_CA_PK
  export GARYTON_UNIVERSITY_CA_PK

  print_step "Start network"
  docker-compose up -d
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
  echo "Chaincode for clayton-university.dl4csr.org has been instantiated"
  docker exec cli.garyton-university.dl4csr.org peer chaincode instantiate \
    -o orderer1.dl4csr.org:7050 \
    -C garytonuniversitychannel \
    -n dl4csr \
    -l golang \
    -v 1.0.0 \
    -c '{"Args":["init"]}' \
    --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/dl4csr.org/orderers/orderer1.dl4csr.org/msp/tlscacerts/tlsca.dl4csr.org-cert.pem
  echo "Chaincode for clayton-university.dl4csr.org has been instantiated"

  print_step "Test chaincode invocation"
  docker exec cli.clayton-university.dl4csr.org peer chaincode invoke \
    -C claytonuniversitychannel \
    -n dl4csr \
    -c '{"Args":["ping"]}' \
    --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/dl4csr.org/orderers/orderer1.dl4csr.org/msp/tlscacerts/tlsca.dl4csr.org-cert.pem
  docker exec cli.garyton-university.dl4csr.org peer chaincode invoke \
    -C garytonuniversitychannel \
    -n dl4csr \
    -c '{"Args":["ping"]}' \
    --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/dl4csr.org/orderers/orderer1.dl4csr.org/msp/tlscacerts/tlsca.dl4csr.org-cert.pem
}

function remove() {
  read -r -n 1 -p 'The network and data will be deleted and cannot be recovered, still remove? [y/N] ' FLAG
  echo
  if [ "$FLAG" != "y" ]; then
    echo "Operation canceled"
    exit 0
  fi
  echo "Removing network..."
  sudo rm -rf ./channel-artifacts
  sudo rm -rf ./crypto-config
  docker-compose down --remove-orphans
  docker system prune --volumes -f
}

function start() {
  echo 'Starting network...'
  docker-compose start
}

function pause() {
  echo 'Pausing network...'
  docker-compose stop
}

function upgrade() {
  docker exec cli.clayton-university.dl4csr.org peer chaincode install \
    -n dl4csr \
    -v "$1" \
    -l golang \
    -p chaincode \
    --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/dl4csr.org/orderers/orderer1.dl4csr.org/msp/tlscacerts/tlsca.dl4csr.org-cert.pem
  docker exec cli.clayton-university.dl4csr.org peer chaincode upgrade \
    -o orderer1.dl4csr.org:7050 \
    -C claytonuniversitychannel \
    -n dl4csr \
    -l golang \
    -v "$1" \
    -c '{"Args":["init"]}' \
    --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/dl4csr.org/orderers/orderer1.dl4csr.org/msp/tlscacerts/tlsca.dl4csr.org-cert.pem
  echo "Chaincode for clayton-university.dl4csr.org has been upgraded"
  docker exec cli.garyton-university.dl4csr.org peer chaincode install \
    -n dl4csr \
    -v "$1" \
    -l golang \
    -p chaincode \
    --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/dl4csr.org/orderers/orderer1.dl4csr.org/msp/tlscacerts/tlsca.dl4csr.org-cert.pem
  docker exec cli.garyton-university.dl4csr.org peer chaincode upgrade \
    -o orderer1.dl4csr.org:7050 \
    -C garytonuniversitychannel \
    -n dl4csr \
    -l golang \
    -v "$1" \
    -c '{"Args":["init"]}' \
    --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/dl4csr.org/orderers/orderer1.dl4csr.org/msp/tlscacerts/tlsca.dl4csr.org-cert.pem
  echo "Chaincode for garyton-university.dl4csr.org has been upgraded"
}

function invoke() {
  echo 'Testing chaincode invocation...'
  if [ "$2" == "" ]; then
    echo "Usage: launcher.sh $1 [FUNCTION] [ARGUMENTS]"
    exit 0
  fi
  docker exec cli.clayton-university.dl4csr.org peer chaincode invoke \
    -C claytonuniversitychannel \
    -n dl4csr \
    -c '{"Args":['"$2"']}' \
    --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/dl4csr.org/orderers/orderer1.dl4csr.org/msp/tlscacerts/tlsca.dl4csr.org-cert.pem
}

case $1 in
"-h" | "--help" | "")
  show_helps
  ;;
"-d" | "--deploy")
  deploy
  ;;
"-r" | "--remove")
  remove
  ;;
"-s" | "--start")
  start
  ;;
"-p" | "--pause")
  pause
  ;;
"-u" | "--upgrade")
  upgrade "$2"
  ;;
"-i" | "--invoke")
  builder=""
  index=1
  for arg in "$@"; do
    if [ "$index" != 1 ]; then
      builder=$builder"\"$arg\""
      if [ "$index" -lt $# ]; then
        builder=$builder","
      fi
    fi
    ((index = "$index" + 1))
  done
  invoke "$1" "$builder"
  ;;
*)
  echo "Unknown flag: $1"
  echo 'Use with flag "--help" or "-h" to show helps'
  ;;
esac
