package route

import (
	"chaincode/lib"
	"chaincode/util"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func createPeerReview(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) < 2 {
		return shim.Error("function createPeerReview requires 2 arguments")
	}
	if args[0] == "" || args[1] == "" {
		return shim.Error("arguments should be nonempty")
	}
	paperID := args[0]
	reviewer := args[1]
	peerReview := lib.PeerReview{
		Paper:    paperID,
		Reviewer: reviewer,
		Status:   lib.StatusReviewing,
	}
	payload, err := util.PutLedger(stub, peerReview)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to put ledger: %s", err.Error()))
	}
	return shim.Success(payload)
}
