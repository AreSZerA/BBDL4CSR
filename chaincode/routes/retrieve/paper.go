package retrieve

import (
	"chaincode/lib"
	"chaincode/utils"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func queryPapers(stub shim.ChaincodeStubInterface, query string) peer.Response {
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

func PapersByQuery(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	query := args[0]
	return queryPapers(stub, query)
}

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

func PapersSortByTitle(stub shim.ChaincodeStubInterface, _ []string) peer.Response {
	query := `{"sort":[{"paper_title":"asc"}]}`
	return queryPapers(stub, query)
}

func PapersSortByUploadTime(stub shim.ChaincodeStubInterface, _ []string) peer.Response {
	query := `{"sort":[{"paper_upload_time":"desc"}]}`
	return queryPapers(stub, query)
}

func AcceptedPapersSortByTitle(stub shim.ChaincodeStubInterface, _ []string) peer.Response {
	query := `{"selector":{"paper_status":"accepted"},"sort":[{"paper_title":"asc"}]}`
	return queryPapers(stub, query)
}

func AcceptedPapersSortByUploadTime(stub shim.ChaincodeStubInterface, _ []string) peer.Response {
	query := `{"selector":{"paper_status":"accepted"},"sort":[{"paper_upload_time":"asc"}]}`
	return queryPapers(stub, query)
}

func AcceptedPapersSortByPublishTime(stub shim.ChaincodeStubInterface, _ []string) peer.Response {
	query := `{"selector":{"paper_status":"accepted"},"sort":[{"paper_publish_time":"asc"}]}`
	return queryPapers(stub, query)
}

func RejectedPapersSortByTitle(stub shim.ChaincodeStubInterface, _ []string) peer.Response {
	query := `{"selector":{"paper_status":"rejected"},"sort":[{"paper_title":"asc"}]}`
	return queryPapers(stub, query)
}

func RejectedPapersSortByUploadTime(stub shim.ChaincodeStubInterface, _ []string) peer.Response {
	query := `{"selector":{"paper_status":"rejected"},"sort":[{"paper_upload_time":"desc"}]}`
	return queryPapers(stub, query)
}

func ReviewingPapersSortByTitle(stub shim.ChaincodeStubInterface, _ []string) peer.Response {
	query := `{"selector":{"paper_status":"reviewing"},"sort":[{"paper_title":"asc"}]}`
	return queryPapers(stub, query)
}

func ReviewingPapersSortByUploadTime(stub shim.ChaincodeStubInterface, _ []string) peer.Response {
	query := `{"selector":{"paper_status":"reviewing"},"sort":[{"paper_upload_time":"desc"}]}`
	return queryPapers(stub, query)
}

func AcceptedPapersByTitleSortByTitle(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	title := args[0]
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_title":{"$regex":".*?%s.*?"}},{"paper_status":"accepted"}]},"sort":[{"paper_title":"asc"}]}`, title)
	return queryPapers(stub, query)
}

func AcceptedPapersByTitleSortByPublishTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	title := args[0]
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_title":{"$regex":".*?%s.*?"}},{"paper_status":"accepted"}]},"sort":[{"paper_publish_time":"desc"}]}`, title)
	return queryPapers(stub, query)
}

func AcceptedPapersByAuthorSortByTitle(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	author := args[0]
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_authors":{"$in":["%s"]}},{"paper_status":"accepted"}]},"sort":[{"paper_title":"asc"}]}`, author)
	return queryPapers(stub, query)
}

func AcceptedPapersByAuthorSortByPublishTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	author := args[0]
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_authors":{"$in":["%s"]}},{"paper_status":"accepted"}]},"sort":[{"paper_publish_time":"desc"}]}`, author)
	return queryPapers(stub, query)
}

func AcceptedPapersByKeywordSortByTitle(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	keyword := args[0]
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_keywords":{"$in":["%s"]}},{"paper_status":"accepted"}]},"sort":[{"paper_title":"asc"}]}`, keyword)
	return queryPapers(stub, query)
}

func AcceptedPapersByKeywordSortByPublishTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	keyword := args[0]
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_keywords":{"$in":["%s"]}},{"paper_status":"accepted"}]},"sort":[{"paper_publish_time":"desc"}]}`, keyword)
	return queryPapers(stub, query)
}

func AcceptedPapersByUploaderSortByTitle(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	uploader := args[0]
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_uploader":"%s"},{"paper_status":"accepted"}]},"sort":[{"paper_title":"asc"}]}`, uploader)
	return queryPapers(stub, query)
}

func AcceptedPapersByUploaderSortByUploadTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	uploader := args[0]
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_uploader":"%s"},{"paper_status":"accepted"}]},"sort":[{"paper_upload_time":"desc"}]}`, uploader)
	return queryPapers(stub, query)
}

func AcceptedPapersByUploaderSortByPublishTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	uploader := args[0]
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_uploader":"%s"},{"paper_status":"accepted"}]},"sort":[{"paper_publish_time":"desc"}]}`, uploader)
	return queryPapers(stub, query)
}

func RejectedPapersByUploaderSortByTitle(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	uploader := args[0]
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_uploader":"%s"},{"paper_status":"rejected"}]},"sort":[{"paper_title":"asc"}]}`, uploader)
	return queryPapers(stub, query)
}

func RejectedPapersByUploaderSortByUploadTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	uploader := args[0]
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_uploader":"%s"},{"paper_status":"rejected"}]},"sort":[{"paper_upload_time":"desc"}]}`, uploader)
	return queryPapers(stub, query)
}

func RejectedPapersByUploaderSortByPublishTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	uploader := args[0]
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_uploader":"%s"},{"paper_status":"rejected"}]},"sort":[{"paper_publish_time":"desc"}]}`, uploader)
	return queryPapers(stub, query)
}

func ReviewingPapersByUploaderSortByTitle(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	uploader := args[0]
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_uploader":"%s"},{"paper_status":"reviewing"}]},"sort":[{"paper_title":"asc"}]}`, uploader)
	return queryPapers(stub, query)
}

func ReviewingPapersByUploaderSortByUploadTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	uploader := args[0]
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_uploader":"%s"},{"paper_status":"reviewing"}]},"sort":[{"paper_upload_time":"desc"}]}`, uploader)
	return queryPapers(stub, query)
}

func PaperById(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	id := args[0]
	paperBytes, err := utils.GetByKeys(stub, lib.ObjectTypePaper, id)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(paperBytes)
}
