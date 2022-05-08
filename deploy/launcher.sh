#!/usr/bin/env bash

ORDERER="orderer1.dl4csr.org:7050"
PK="/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/dl4csr.org/orderers/orderer1.dl4csr.org/msp/tlscacerts/tlsca.dl4csr.org-cert.pem"
CU_CHANNEL="claytonuniversitychannel"
GU_CHANNEL="garytonuniversitychannel"
CU_CLI="cli.clayton-university.dl4csr.org"
GU_CLI="cli.garyton-university.dl4csr.org"

INIT_ARGS='{"Args":["init"]}'
TEST_ARGS='{"Args":["ping"]}'

COUNT_STEP=1

function confirm() {
  read -r -n 1 -p "$1 [y/N] " input
  echo
  if [ "$input" != "y" ]; then
    echo "Operation canceled"
    echo
    exit 0
  fi
}

function print_step() {
  echo -e "\n\033[32m>>> Step ${COUNT_STEP}: $1\033[0m"
  ((COUNT_STEP++))
}

function print_helps() {
  echo 'Usage: launcher.sh [COMMAND]

The commands are:
  -h, --help       Show this help
  -d, --deploy     Deploy or redeploy the network, blockchain will be wiped then recreated
  -r, --remove     Remove the network, volumes will be deleted and data will be wiped
  -s, --start      Start the existing network, the blockchain will be restored
  -p, --pause      Stop the existing network, the blockchain still exists in containers
  -u, --upgrade    Upgrade chaincode, use with a new version number
  -i, --invoke     Function invocation test of the chaincode, use with function name and arguments'
}

function print_logo() {
  echo '
======================================================================================
   oooooooooo.    ooooo               .o      .oooooo.     .oooooo..o  ooooooooo.
   `888'"'"'   `Y8b   `888'"'"'             .d88     d8P'"'"'  `Y8b   d8P'"'"'    `Y8  `888   `Y88.
    888      888   888            .d'"'"'888    888           Y88bo.        888   .d88'"'"'
    888      888   888          .d'"'"'  888    888            `"Y8888o.    888ooo88P'"'"'
    888      888   888          88ooo888oo  888                `"Y88b   888`88b.
    888     d88'"'"'   888       o       888    `88b    ooo   oo     .d8P   888  `88b.
   o888bood8P'"'"'    o888ooooood8      o888o    `Y8bood8P'"'"'   8""88888P'"'"'   o888o  o888o
======================================================================================'
}

function deploy_network() {
  print_step "Deploy network"
  docker-compose up -d
  echo "Sleep 10 seconds for kafka cluster to complete booting..."
  sleep 10
}

function start_network() {
  print_step "Start network"
  docker-compose start
  echo "Sleep 10 seconds for network to complete initialization..."
  sleep 10
}

function stop_network() {
  print_step "Stop network"
  docker-compose stop
}

function create_channels() {
  print_step "Create channels"
  CU_TX="/opt/gopath/src/github.com/hyperledger/fabric/peer/channel-artifacts/$CU_CHANNEL.tx"
  GU_TX="/opt/gopath/src/github.com/hyperledger/fabric/peer/channel-artifacts/$GU_CHANNEL.tx"
  docker exec "$CU_CLI" peer channel create -o "$ORDERER" -c "$CU_CHANNEL" -f "$CU_TX" --tls --cafile "$PK"
  docker exec "$GU_CLI" peer channel create -o "$ORDERER" -c "$GU_CHANNEL" -f "$GU_TX" --tls --cafile "$PK"
}

function join_channels() {
  print_step "Join channels"
  docker exec "$CU_CLI" peer channel join -b "$CU_CHANNEL".block --tls --cafile "$PK"
  docker exec "$GU_CLI" peer channel join -b "$GU_CHANNEL".block --tls --cafile "$PK"
}

function generate_files() {
  print_step "Generates certifications, genesis block, and channel configuration transactions"
  mkdir ./channel-artifacts
  mkdir ./crypto-config
  export PATH=${PWD}/bin:${PWD}:$PATH
  cryptogen generate --config=./crypto-config.yaml
  configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./channel-artifacts/genesis.block -channelID genesischannel
  configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/"$CU_CHANNEL".tx -channelID "$CU_CHANNEL"
  configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/"$GU_CHANNEL".tx -channelID "$GU_CHANNEL"
  CLAYTON_UNIVERSITY_CA_PK=$(cd crypto-config/peerOrganizations/clayton-university.dl4csr.org/ca && ls *_sk)
  GARYTON_UNIVERSITY_CA_PK=$(cd crypto-config/peerOrganizations/garyton-university.dl4csr.org/ca && ls *_sk)
  export CLAYTON_UNIVERSITY_CA_PK
  export GARYTON_UNIVERSITY_CA_PK
}

function remove_files() {
  print_step "Remove generated files"
  sudo rm -rf ./channel-artifacts
  sudo rm -rf ./crypto-config
}

function remove_containers() {
  print_step "Remove containers"
  docker-compose down --remove-orphans
}

function install_chaincode() {
  print_step "Install chaincode"
  docker exec "$CU_CLI" peer chaincode install -n dl4csr -v "$1" -l golang -p chaincode --tls --cafile "$PK"
  docker exec "$GU_CLI" peer chaincode install -n dl4csr -v "$1" -l golang -p chaincode --tls --cafile "$PK"
}

function instantiate_chaincode() {
  print_step "Instantiate chaincode"
  docker exec "$CU_CLI" peer chaincode instantiate -o "$ORDERER" -C "$CU_CHANNEL" -n dl4csr -l golang -v "$1" -c "$INIT_ARGS" --tls --cafile "$PK"
  sleep 5
  echo "Chaincode for clayton-university.dl4csr.org has been instantiated"
  docker exec "$GU_CLI" peer chaincode instantiate -o "$ORDERER" -C "$GU_CHANNEL" -n dl4csr -l golang -v "$1" -c "$INIT_ARGS" --tls --cafile "$PK"
  sleep 5
  echo "Chaincode for garyton-university.dl4csr.org has been instantiated"
}

function upgrade_chaincode() {
  print_step "Upgrade chaincode"
  if [ "$1" == "" ]; then
    echo "Please select a version to update"
    return
  fi
  docker exec "$CU_CLI" peer chaincode upgrade -o "$ORDERER" -C "$CU_CHANNEL" -n dl4csr -l golang -v "$1" -c "$INIT_ARGS" --tls --cafile "$PK"
  sleep 5
  echo "Chaincode for clayton-university.dl4csr.org has been upgraded"
  docker exec "$GU_CLI" peer chaincode upgrade -o "$ORDERER" -C "$GU_CHANNEL" -n dl4csr -l golang -v "$1" -c "$INIT_ARGS" --tls --cafile "$PK"
  sleep 5
  echo "Chaincode for garyton-university.dl4csr.org has been upgraded"
}

function test_chaincode() {
  print_step "Test chaincode invocation"
  docker exec "$CU_CLI" peer chaincode invoke -C "$CU_CHANNEL" -n dl4csr -c "$TEST_ARGS" --tls --cafile "$PK"
  docker exec "$GU_CLI" peer chaincode invoke -C "$GU_CHANNEL" -n dl4csr -c "$TEST_ARGS" --tls --cafile "$PK"
}

function invoke_chaincode() {
  print_step "Invoke chaincode"
  if [ "$2" == "" ]; then
    echo "Usage: launcher.sh $1 [FUNCTION] [ARGUMENTS]"
    return
  fi
  echo "Arguments: $2"
  docker exec "$CU_CLI" peer chaincode invoke -C "$CU_CHANNEL" -n dl4csr -c '{"Args":['"$2"']}' --tls --cafile "$PK"
  docker exec "$GU_CLI" peer chaincode invoke -C "$GU_CHANNEL" -n dl4csr -c '{"Args":['"$2"']}' --tls --cafile "$PK"
}

case $1 in
"-h" | "--help" | "")
  print_helps
  ;;
"-d" | "--deploy")
  confirm "Data will lost if the network has been deployed before, still deploy?"
  print_logo
  remove_files
  remove_containers
  generate_files
  deploy_network
  create_channels
  join_channels
  install_chaincode "1.0.0"
  instantiate_chaincode "1.0.0"
  test_chaincode
  ;;
"-r" | "--remove")
  confirm "The network and data will be deleted and cannot be recovered, still remove?"
  remove_files
  remove_containers
  ;;
"-s" | "--start")
  start_network
  test_chaincode
  ;;
"-p" | "--pause")
  stop_network
  ;;
"-u" | "--upgrade")
  install_chaincode "$2"
  upgrade_chaincode "$2"
  test_chaincode
  ;;
"-i" | "--invoke")
  args=""
  index=1
  for arg in "$@"; do
    if [ "$index" != 1 ]; then
      args=$args"\"$arg\""
      if [ "$index" -lt $# ]; then
        args=$args","
      fi
    fi
    ((index = "$index" + 1))
  done
  invoke_chaincode "$1" "$args"
  ;;
*)
  echo "Unknown flag: $1"
  echo 'Use with flag "--help" or "-h" to show helps'
  ;;
esac
echo
