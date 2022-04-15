package retrieve

import (
	"chaincode/lib"
	"chaincode/utils"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func Papers(stub shim.ChaincodeStubInterface, _ []string) peer.Response {
	results, err := utils.GetAll(stub, lib.ObjectTypePaper)
	if err != nil {
		return shim.Error(err.Error())
	}
	var papers []lib.Paper
	for _, result := range results {
		var paper lib.Paper
		_ = json.Unmarshal(result, &paper)
		papers = append(papers, paper)
	}
	papersBytes, _ := json.Marshal(papers)
	return shim.Success(papersBytes)
}

func AcceptedPapersByTitle(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	title := args[0]
	query := `{"selector":{"$and":[{"$regex":{"paper_title":"^.*?` + title + `.*?$"}},{"paper_status":"accepted"}]},"sort":[{"paper_review_time":"desc"}]}`
	results, err := utils.GetByQuery(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	var papers []lib.Paper
	for _, result := range results {
		var paper lib.Paper
		_ = json.Unmarshal(result, &paper)
		papers = append(papers, paper)
	}
	papersBytes, _ := json.Marshal(papers)
	return shim.Success(papersBytes)
}

func AcceptedPapersByAuthor(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	author := args[0]
	query := `{"selector":{"$and":[{"$in":{"paper_author":"` + author + `"}},{"paper_status":"accepted"}]},"sort":[{"paper_review_time":"desc"}]}`
	results, err := utils.GetByQuery(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	var papers []lib.Paper
	for _, result := range results {
		var paper lib.Paper
		_ = json.Unmarshal(result, &paper)
		papers = append(papers, paper)
	}
	papersBytes, _ := json.Marshal(papers)
	return shim.Success(papersBytes)
}

func AcceptedPapersByKeyword(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	keyword := args[0]
	query := `{"selector":{"$and":[{"$in":{"paper_keywords":"` + keyword + `"}},{"paper_status":"accepted"}]},"sort":[{"paper_review_time":"desc"}]}`
	results, err := utils.GetByQuery(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	var papers []lib.Paper
	for _, result := range results {
		var paper lib.Paper
		_ = json.Unmarshal(result, &paper)
		papers = append(papers, paper)
	}
	papersBytes, _ := json.Marshal(papers)
	return shim.Success(papersBytes)
}

func AcceptedPapers(stub shim.ChaincodeStubInterface, _ []string) peer.Response {
	query := `{"selector":{"paper_status":"accepted"},"sort":[{"paper_upload_time":"desc"}]}`
	results, err := utils.GetByQuery(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	var papers []lib.Paper
	for _, result := range results {
		var paper lib.Paper
		_ = json.Unmarshal(result, &paper)
		papers = append(papers, paper)
	}
	papersBytes, _ := json.Marshal(papers)
	return shim.Success(papersBytes)
}

func RejectedPapers(stub shim.ChaincodeStubInterface, _ []string) peer.Response {
	query := `{"selector":{"paper_status":"rejected"},"sort":[{"paper_upload_time":"desc"}]}`
	results, err := utils.GetByQuery(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	var papers []lib.Paper
	for _, result := range results {
		var paper lib.Paper
		_ = json.Unmarshal(result, &paper)
		papers = append(papers, paper)
	}
	papersBytes, _ := json.Marshal(papers)
	return shim.Success(papersBytes)
}

func RejectedPapersByAuthor(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	author := args[0]
	query := `{"selector":{"$and":[{"$in":{"paper_author":"` + author + `"}},{"paper_status":"rejected"}]},"sort":[{"paper_review_time":"desc"}]}`
	results, err := utils.GetByQuery(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	var papers []lib.Paper
	for _, result := range results {
		var paper lib.Paper
		_ = json.Unmarshal(result, &paper)
		papers = append(papers, paper)
	}
	papersBytes, _ := json.Marshal(papers)
	return shim.Success(papersBytes)
}

func ReviewingPapers(stub shim.ChaincodeStubInterface, _ []string) peer.Response {
	query := `{"selector":{"paper_status":"reviewing"},"sort":[{"paper_upload_time":"desc"}]}`
	results, err := utils.GetByQuery(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	var papers []lib.Paper
	for _, result := range results {
		var paper lib.Paper
		_ = json.Unmarshal(result, &paper)
		papers = append(papers, paper)
	}
	papersBytes, _ := json.Marshal(papers)
	return shim.Success(papersBytes)
}

func ReviewingPapersByAuthor(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	author := args[0]
	query := `{"selector":{"$and":[{"$in":{"paper_author":"` + author + `"}},{"paper_status":"reviewing"}]},"sort":[{"paper_review_time":"desc"}]}`
	results, err := utils.GetByQuery(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	var papers []lib.Paper
	for _, result := range results {
		var paper lib.Paper
		_ = json.Unmarshal(result, &paper)
		papers = append(papers, paper)
	}
	papersBytes, _ := json.Marshal(papers)
	return shim.Success(papersBytes)
}

func PaperById(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	paperId := args[0]
	paperBytes, err := utils.GetByKeys(stub, lib.ObjectTypePaper, paperId)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(paperBytes)
}
