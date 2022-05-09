#!/usr/bin/env bash

# Order address, options:
# - orderer1.dl4csr.org:7050
# - orderer2.dl4csr.org:7050
# - orderer3.dl4csr.org:7050
ORDERER="orderer1.dl4csr.org:7050"

# Primary key for TLS, related to the selected order
PK="/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/dl4csr.org/orderers/orderer1.dl4csr.org/msp/tlscacerts/tlsca.dl4csr.org-cert.pem"

# Channel names
CU_CHANNEL="claytonuniversitychannel"
GU_CHANNEL="garytonuniversitychannel"

# CLI containers
CU_CLI="cli.clayton-university.dl4csr.org"
GU_CLI="cli.garyton-university.dl4csr.org"

# Arguments to initialize or invoke ping
INIT_ARGS='{"Args":["init"]}'
TEST_ARGS='{"Args":["ping"]}'

COUNT_STEP=1 # for counting steps

# confirm() blocks the program until user input.
function confirm() {
  # read user input
  read -r -n 1 -p "$1 [y/N] " input
  echo
  # if the user input is not "y", exit the script with code 0
  if [ "$input" != "y" ]; then
    echo "Operation canceled"
    echo
    exit 0
  fi
}

# print_step() prints step number and title.
function print_step() {
  # the first argument is the step title to print
  echo -e "\n\033[32m>>> Step ${COUNT_STEP}: $1\033[0m"
  # increment step number
  ((COUNT_STEP++))
}

# print_helps() displays the help information.
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

# print_logo() prints the ASCII painting of DL4CSR.
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

# deploy_network() starts the network.
function deploy_network() {
  print_step "Deploy network"
  # read the ./docker-compose.yaml file to deploy network
  # the flag -d is used to hide debug logs
  docker-compose up -d
  echo "Sleep 10 seconds for containers to complete booting..."
  sleep 10
}

# start_network() starts the existing network.
function start_network() {
  print_step "Start network"
  # start the existing network according to the ./docker-compose.yaml file
  docker-compose start
  echo "Sleep 10 seconds for containers to complete booting..."
  sleep 10
}

# stop_network() stops the network.
function stop_network() {
  print_step "Stop network"
  # stop containers in the network in accordance with the docker-compose.yaml file
  docker-compose stop
}

# create_channels() creates channels via CLI containers.
function create_channels() {
  print_step "Create channels"
  CU_TX="/opt/gopath/src/github.com/hyperledger/fabric/peer/channel-artifacts/$CU_CHANNEL.tx"
  GU_TX="/opt/gopath/src/github.com/hyperledger/fabric/peer/channel-artifacts/$GU_CHANNEL.tx"
  # -o for orderer address, -c for channel name, -f for transaction file, --tls to enable TLS, and -cafile for CA file
  docker exec "$CU_CLI" peer channel create -o "$ORDERER" -c "$CU_CHANNEL" -f "$CU_TX" --tls --cafile "$PK"
  docker exec "$GU_CLI" peer channel create -o "$ORDERER" -c "$GU_CHANNEL" -f "$GU_TX" --tls --cafile "$PK"
}

# join_channels() joins channels with genesis blocks.
function join_channels() {
  print_step "Join channels"
  # -b for genesis block, --tls to enable TLS, and -cafile for CA file
  docker exec "$CU_CLI" peer channel join -b "$CU_CHANNEL".block --tls --cafile "$PK"
  docker exec "$GU_CLI" peer channel join -b "$GU_CHANNEL".block --tls --cafile "$PK"
}

# generate_files() generates CA files, genesis block, and channel configuration transactions.
function generate_files() {
  print_step "Generates certifications, genesis block, and channel configuration transactions"
  # create directories
  mkdir ./channel-artifacts
  mkdir ./crypto-config
  # export binary tools
  export PATH=${PWD}/bin:${PWD}:$PATH
  # generate CA files according to the ./crypto-config.yaml file
  cryptogen generate --config=./crypto-config.yaml
  # generate genesis block in ./channel-artifacts/ named genesis.block
  configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./channel-artifacts/genesis.block -channelID genesischannel
  # create channel configuration transactions in ./channel-artifacts/
  configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/"$CU_CHANNEL".tx -channelID "$CU_CHANNEL"
  configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/"$GU_CHANNEL".tx -channelID "$GU_CHANNEL"
  # export CA files which will be used in ./docker-compose.yaml when deploying the network
  CLAYTON_UNIVERSITY_CA_PK=$(cd crypto-config/peerOrganizations/clayton-university.dl4csr.org/ca && ls *_sk)
  GARYTON_UNIVERSITY_CA_PK=$(cd crypto-config/peerOrganizations/garyton-university.dl4csr.org/ca && ls *_sk)
  export CLAYTON_UNIVERSITY_CA_PK
  export GARYTON_UNIVERSITY_CA_PK
}

# remove_files() removes the ./channel-artifacts/ and ./crypto-config/ recursively.
function remove_files() {
  print_step "Remove generated files"
  sudo rm -rf ./channel-artifacts
  sudo rm -rf ./crypto-config
}

# remove_containers() stops and removes the containers.
function remove_containers() {
  print_step "Remove containers"
  # remove the containers according to the ./docker-compose.yaml file
  docker-compose down --remove-orphans
}

# install_chaincode() installs chaincode via CLI.
function install_chaincode() {
  print_step "Install chaincode"
  # -n for chaincode name, -v for version number, -l for language, --tls to enable TLS, and -cafile for CA file
  # the version number is the first argument received by this function
  docker exec "$CU_CLI" peer chaincode install -n dl4csr -v "$1" -l golang -p chaincode --tls --cafile "$PK"
  docker exec "$GU_CLI" peer chaincode install -n dl4csr -v "$1" -l golang -p chaincode --tls --cafile "$PK"
}

# instantiate_chaincode() instantiates chaincode.
function instantiate_chaincode() {
  print_step "Instantiate chaincode"
  # -o for orderer address, -C for channel name, -n for chaincode name, -l for language, -v for version number
  # -c for arguments, --tls to enable TLS, and -cafile for CA file
  # the version number is the first argument received by this function
  docker exec "$CU_CLI" peer chaincode instantiate -o "$ORDERER" -C "$CU_CHANNEL" -n dl4csr -l golang -v "$1" -c "$INIT_ARGS" --tls --cafile "$PK"
  sleep 5
  echo "Chaincode for clayton-university.dl4csr.org has been instantiated"
  docker exec "$GU_CLI" peer chaincode instantiate -o "$ORDERER" -C "$GU_CHANNEL" -n dl4csr -l golang -v "$1" -c "$INIT_ARGS" --tls --cafile "$PK"
  sleep 5
  echo "Chaincode for garyton-university.dl4csr.org has been instantiated"
}

# upgrade_chaincode() upgrades chaincode.
function upgrade_chaincode() {
  print_step "Upgrade chaincode"
  if [ "$1" == "" ]; then
    echo "Please select a version to update"
    return
  fi
  # -o for orderer address, -C for channel name, -n for chaincode name, -l for language, -v for version number
  # -c for arguments, --tls to enable TLS, and -cafile for CA file
  # the version number is the first argument received by this function
  docker exec "$CU_CLI" peer chaincode upgrade -o "$ORDERER" -C "$CU_CHANNEL" -n dl4csr -l golang -v "$1" -c "$INIT_ARGS" --tls --cafile "$PK"
  sleep 5
  echo "Chaincode for clayton-university.dl4csr.org has been upgraded"
  docker exec "$GU_CLI" peer chaincode upgrade -o "$ORDERER" -C "$GU_CHANNEL" -n dl4csr -l golang -v "$1" -c "$INIT_ARGS" --tls --cafile "$PK"
  sleep 5
  echo "Chaincode for garyton-university.dl4csr.org has been upgraded"
}

# test_chaincode() invokes "ping" to test connectivity and chaincode availability.
function test_chaincode() {
  print_step "Test chaincode invocation"
  # -C for channel name, -n for chaincode name, -c for arguments, --tls to enable TLS, and -cafile for CA file
  docker exec "$CU_CLI" peer chaincode invoke -C "$CU_CHANNEL" -n dl4csr -c "$TEST_ARGS" --tls --cafile "$PK"
  docker exec "$GU_CLI" peer chaincode invoke -C "$GU_CHANNEL" -n dl4csr -c "$TEST_ARGS" --tls --cafile "$PK"
}

# invoke_chaincode() invokes chaincode with function name and arguments.
function invoke_chaincode() {
  print_step "Invoke chaincode"
  if [ "$2" == "" ]; then
    echo "Usage: launcher.sh $1 [FUNCTION] [ARGUMENTS]"
    return
  fi
  echo "Arguments: $2"
  # -C for channel name, -n for chaincode name, -c for arguments, --tls to enable TLS, and -cafile for CA file
  docker exec "$CU_CLI" peer chaincode invoke -C "$CU_CHANNEL" -n dl4csr -c '{"Args":['"$2"']}' --tls --cafile "$PK"
  docker exec "$GU_CLI" peer chaincode invoke -C "$GU_CHANNEL" -n dl4csr -c '{"Args":['"$2"']}' --tls --cafile "$PK"
}

# Entrance of the script.
case $1 in
"-h" | "--help" | "") # display help information
  print_helps
  ;;
"-d" | "--deploy") # deploy network
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
"-r" | "--remove") # remove network
  confirm "The network and data will be deleted and cannot be recovered, still remove?"
  remove_files
  remove_containers
  ;;
"-s" | "--start") # resume network
  start_network
  test_chaincode
  ;;
"-p" | "--pause") # stop network
  stop_network
  ;;
"-u" | "--upgrade") # upgrade chaincode
  install_chaincode "$2"
  upgrade_chaincode "$2"
  test_chaincode
  ;;
"-i" | "--invoke") # invoke chaincode functions
  args=""
  index=1
  # traverse the arguments
  for arg in "$@"; do
    if [ "$index" != 1 ]; then
      # add quotes
      args=$args"\"$arg\""
      if [ "$index" -lt $# ]; then
        # add coma
        args=$args","
      fi
    fi
    ((index = "$index" + 1))
  done
  invoke_chaincode "$1" "$args"
  ;;
*) # undefined flags
  echo "Unknown flag: $1"
  echo 'Use with flag "--help" or "-h" to show helps'
  ;;
esac
echo
