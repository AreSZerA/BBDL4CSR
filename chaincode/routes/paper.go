package routes

import (
	"chaincode/lib"
	"chaincode/utils"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"time"
)

func CreatePaper(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	now := time.Now().UnixNano()
	if len(args) < 5 {
		return shim.Error("function CreatePaper requires 5 arguments")
	}
	if args[0] == "" || args[1] == "" || args[2] == "" || args[3] == "" || args[4] == "" {
		return shim.Error("arguments should be nonempty")
	}
	email := args[0]
	title := args[1]
	abstract := args[2]
	authorsJSON := args[3]
	keywordsJSON := args[4]
	var authors []string
	err := json.Unmarshal([]byte(authorsJSON), &authors)
	if err != nil {
		return shim.Error(fmt.Sprintf("authors not in JSON: %s", err.Error()))
	}
	var keywords []string
	err = json.Unmarshal([]byte(keywordsJSON), &keywords)
	if err != nil {
		return shim.Error(fmt.Sprintf("keywords not in JSON: %s", err.Error()))
	}
	userResp := RetrieveUserByEmail(stub, []string{email})
	if userResp.Status != shim.OK {
		return shim.Error(fmt.Sprintf("cannot verify email: %s", userResp.Message))
	}
	if userResp.Payload == nil {
		return shim.Error(fmt.Sprintf("uploader %s not exists", email))
	}
	reviewerResp := retrieveAllSortedReviewers(stub, nil)
	if reviewerResp.Status != shim.OK {
		return shim.Error(fmt.Sprintf("failed to distribute reviewers: %s", reviewerResp.Message))
	}
	var reviewers []lib.User
	_ = json.Unmarshal(reviewerResp.Payload, &reviewers)
	if len(reviewers) < 3 {
		return shim.Error("not enough reviewers to distribute")
	}
	var reviewerEmails [3]string
	id := fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%s%d", email, now))))
	for i := 0; i < 3; i++ {
		reviewerEmails[i] = reviewers[i].Email
		_ = updateUserReviewingAdd(stub, []string{reviewers[i].Email})
		_ = createPeerReview(stub, []string{id, reviewers[i].Email})
	}
	paper := lib.Paper{
		ID:          id,
		Uploader:    email,
		UploadTime:  now,
		Title:       title,
		Abstract:    abstract,
		Authors:     authors,
		Keywords:    keywords,
		PeerReviews: reviewerEmails,
		Status:      lib.StatusReviewing,
	}
	payload, err := utils.PutLedger(stub, paper)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to put paper to ledger: %s", err.Error()))
	}
	return shim.Success(payload)
}

func UpdatePaperStatus(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) < 1 {
		return shim.Error("function UpdatePaperStatus requires 1 argument")
	}
	if args[0] == "" {
		return shim.Error("argument should be nonempty")
	}
	paperResp := RetrievePaperById(stub, args)
	if paperResp.Status != shim.OK {
		return shim.Error(fmt.Sprintf("failed to get paper information: %s", paperResp.Message))
	}
	var paper lib.Paper
	err := json.Unmarshal(paperResp.Payload, &paper)
	resp := RetrievePeerReviewsByPaperId(stub, []string{paper.ID})
	if resp.Status != shim.OK {
		return shim.Error(fmt.Sprintf("failed to retrieve peer review information: %s", resp.Message))
	}
	var peerReviews []lib.PeerReview
	err = json.Unmarshal(resp.Payload, &peerReviews)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to userialise peer review information: %s", err.Error()))
	}
	var statuses []string
	for _, peerReview := range peerReviews {
		statuses = append(statuses, peerReview.Status)
	}
	status := utils.GetStatus(statuses[0], statuses[1], statuses[2])
	if paper.Status == status {
		return shim.Success(nil)
	}
	paper.Status = status
	payload, err := utils.PutLedger(stub, paper)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to update ledger: %s", err.Error()))
	}
	return shim.Success(payload)
}

func RetrieveAllPapers(stub shim.ChaincodeStubInterface, _ []string) peer.Response {
	var papers []lib.Paper
	results, err := utils.GetAll(stub, lib.ObjectTypePaper)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to retrieve papers: %s", err.Error()))
	}
	for _, paperBytes := range results {
		var paper lib.Paper
		err = json.Unmarshal(paperBytes, &paper)
		if err != nil {
			return shim.Error(fmt.Sprintf("failed to unserialise paper: %s", err.Error()))
		}
		papers = append(papers, paper)
	}
	papersBytes, err := json.Marshal(papers)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to serialise papers: %s", err.Error()))
	}
	return shim.Success(papersBytes)
}

func retrievePapersByQuery(stub shim.ChaincodeStubInterface, query string) peer.Response {
	var papers []lib.Paper
	results, err := utils.GetByQuery(stub, query)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to retrieve by query %s: %s", query, err.Error()))
	}
	for _, paperBytes := range results {
		var paper lib.Paper
		err = json.Unmarshal(paperBytes, &paper)
		if err != nil {
			return shim.Error(fmt.Sprintf("failed to unserialise paper: %s", err.Error()))
		}
		papers = append(papers, paper)
	}
	papersBytes, err := json.Marshal(papers)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to serialise papers: %s", err.Error()))
	}
	return shim.Success(papersBytes)
}

func RetrieveAcceptedPapers(stub shim.ChaincodeStubInterface, _ []string) peer.Response {
	query := `{"selector":{"status":"accepted"}}`
	return retrievePapersByQuery(stub, query)
}

func RetrieveRejectedPapers(stub shim.ChaincodeStubInterface, _ []string) peer.Response {
	query := `{"selector":{"status":"rejected"}}`
	return retrievePapersByQuery(stub, query)
}

func RetrieveReviewingPapers(stub shim.ChaincodeStubInterface, _ []string) peer.Response {
	query := `{"selector":{"status":"reviewing"}}`
	return retrievePapersByQuery(stub, query)
}

func RetrievePapersByEmail(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) < 1 {
		return shim.Error("function RetrievePapersByEmail requires 1 argument")
	}
	if args[0] == "" {
		return shim.Error("argument should be nonempty")
	}
	query := `{"selector":{"uploader":"` + args[0] + `"}}`
	return retrievePapersByQuery(stub, query)
}

func RetrievePapersByTitle(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) < 1 {
		return shim.Error("function RetrievePapersByTitle requires 1 argument")
	}
	if args[0] == "" {
		return shim.Error("argument should be nonempty")
	}
	query := `{"selector":{"uploader":"` + args[0] + `"}}`
	return retrievePapersByQuery(stub, query)
}

func RetrievePaperById(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) < 1 {
		return shim.Error("function RetrievePaperById requires 1 argument")
	}
	if args[0] == "" {
		return shim.Error("argument should be nonempty")
	}
	query := `{"selector":{"id":"` + args[0] + `"}}`
	results, err := utils.GetByQuery(stub, query)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to retrieve by query %s: %s", query, err.Error()))
	}
	if results == nil {
		return shim.Success(nil)
	}
	var paper lib.Paper
	err = json.Unmarshal(results[0], &paper)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to unserialise paper: %s", err.Error()))
	}
	paperBytes, err := json.Marshal(paper)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to serialise papers: %s", err.Error()))
	}
	return shim.Success(paperBytes)
}
