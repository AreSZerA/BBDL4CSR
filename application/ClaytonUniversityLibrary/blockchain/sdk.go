package blockchain

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"log"
)

const (
	channelName   = "claytonuniversitychannel"
	chaincodeName = "dl4csr"
	orgName       = "clayton-university.dl4csr.org"
	peerName      = "peer0.clayton-university.dl4csr.org"
	userName      = "Admin"
	configFile    = "conf/sdk.yaml"
)

var sdk *fabsdk.FabricSDK

func init() {
	var err error
	sdk, err = fabsdk.New(config.FromFile(configFile))
	if err != nil {
		log.Panicf("Failed to initialise SDK: %s", err.Error())
	}
}

// Execute for updating
func Execute(functionName string, args ...[]byte) (channel.Response, error) {
	ctx := sdk.ChannelContext(channelName, fabsdk.WithOrg(orgName), fabsdk.WithUser(userName))
	cli, err := channel.New(ctx)
	if err != nil {
		return channel.Response{}, err
	}
	req := channel.Request{ChaincodeID: chaincodeName, Fcn: functionName, Args: args}
	resp, err := cli.Execute(req, channel.WithTargetEndpoints(peerName))
	if err != nil {
		return channel.Response{}, err
	}
	return resp, nil
}

// Query for retrieving only
func Query(functionName string, args ...[]byte) (channel.Response, error) {
	ctx := sdk.ChannelContext(channelName, fabsdk.WithOrg(orgName), fabsdk.WithUser(userName))
	cli, err := channel.New(ctx)
	if err != nil {
		return channel.Response{}, err
	}
	req := channel.Request{ChaincodeID: chaincodeName, Fcn: functionName, Args: args}
	resp, err := cli.Query(req, channel.WithTargetEndpoints(peerName))
	if err != nil {
		return channel.Response{}, err
	}
	return resp, nil
}
