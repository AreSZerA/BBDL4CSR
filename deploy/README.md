# DL4CSR Deployment

## Files

The file tree shows as follows:

```text
.
├── bin
│   ├── configtxgen
│   ├── configtxlator
│   ├── cryptogen
│   ├── discover
│   ├── fabric-ca-client
│   ├── idemixgen
│   ├── orderer
│   └── peer
├── configtx.yaml
├── crypto-config.yaml
├── docker-compose-base.yaml
├── docker-compose.yaml
└── launcher.sh
```

The files are used for:

- Binary files in `bin/` are executable for generating files provided by Hyperledger Fabric.
- `configtx.yaml` is used for `bin/configtxgen` to create genesis blocks and channel configuration transactions.
- `crypto-config.yaml` defines the organizations and is used for `bin/cryptogen` to generate CA files for TLS.
- `docker-compose.yaml` is used for `docker-compose` to deploy and start network.
- `docker-compose-base.yaml` is the base file for the former one, in which the configurations are reused.
- `launcher.sh` is the bash script to deploy, remove, stop, start network and interact with the blockchain.
  Execute `./launcher.sh -h` in CLI, the help information will be displayed.

## Network Architecture

The dependence relationships can be graphed as below:

```text
                                                +------+
                                           +--> | ccli |
                                           |    +------+
              +----+                    +--+--+
          +---| k1 |<--+            +---| cp0 |---+
 +----+   |   +----+   |   +----+   |   +-----+   |    +-----+
 | z1 |<--+            +---| o1 |<--+             +--> | cdb |
 +----+   |   +----+   |   +----+   |   +-----+   |    +-----+
          +---| k2 |<--+            +---| cp1 |---+
 +----+   |   +----+   |   +----+   |   +-----+
 | z2 |<--+            +---| o2 |<--+
 +----+   |   +----+   |   +----+   |   +-----+
          +---| k3 |<--+            +---| gp0 |---+
 +----+   |   +----+   |   +----+   |   +-----+   |    +-----+
 | z3 |<--+            +---| o3 |<--+             +--> | gdb |
 +----+   |   +----+   |   +----+   |   +-----+   |    +-----+
          +---| k4 |<--+            +---| gp1 |<--+
              +----+                    +--+--+
                                           |    +------+
                                           +--> | gcli |
                                                +------+
```

The following table has listed the components of the network in detail:

| abbr. |                 name                  |              image               |   dependence    |
|:-----:|:-------------------------------------:|:--------------------------------:|:---------------:|
|  z1   |         zookeeper1.dl4csr.org         | hyperledger/fabric-zookeeper:0.4 |        -        |
|  z2   |         zookeeper2.dl4csr.org         | hyperledger/fabric-zookeeper:0.4 |        -        |
|  z3   |         zookeeper3.dl4csr.org         | hyperledger/fabric-zookeeper:0.4 |        -        |
|  k1   |           kafka1.dl4csr.org           |   hyperledger/fabric-kafka:0.4   |   z1, z2, z3    |
|  k2   |           kafka2.dl4csr.org           |   hyperledger/fabric-kafka:0.4   |   z1, z2, z3    |
|  k3   |           kafka3.dl4csr.org           |   hyperledger/fabric-kafka:0.4   |   z1, z2, z3    |
|  k4   |           kafka4.dl4csr.org           |   hyperledger/fabric-kafka:0.4   |   z1, z2, z3    |
|  o1   |          orderer1.dl4csr.org          |  hyperledger/fabric-orderer:1.4  | k1, k2, k3, k4  |
|  o2   |          orderer2.dl4csr.org          |  hyperledger/fabric-orderer:1.4  | k1, k2, k3, k4  |
|  o3   |          orderer3.dl4csr.org          |  hyperledger/fabric-orderer:1.4  | k1, k2, k3, k4  |
|  cdb  | couchdb.clayton-university.dl4csr.org |  hyperledger/fabric-couchdb:0.4  |        -        |
|  gdb  | couchdb.garyton-university.dl4csr.org |  hyperledger/fabric-couchdb:0.4  |        -        |
|  cp0  |  peer0.clayton-university.dl4csr.org  |   hyperledger/fabric-peer:1.4    | 01, 02, 03, cdb |
|  cp1  |  peer1.clayton-university.dl4csr.org  |   hyperledger/fabric-peer:1.4    | 01, 02, 03, cdb |
|  gp0  |  peer0.garyton-university.dl4csr.org  |   hyperledger/fabric-peer:1.4    | 01, 02, 03, gdb |
|  gp1  |  peer1.garyton-university.dl4csr.org  |   hyperledger/fabric-peer:1.4    | 01, 02, 03, gdb |
| ccli  |   cli.clayton-university.dl4csr.org   |   hyperledger/fabric-tools:1.4   |       cp0       |
| gcli  |   cli.garyton-university.dl4csr.org   |   hyperledger/fabric-tools:1.4   |       gp1       |
