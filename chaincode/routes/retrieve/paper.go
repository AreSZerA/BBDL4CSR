package retrieve

import (
	"chaincode/lib"
	"chaincode/utils"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func queryPapers(stub shim.ChaincodeStubInterface, query string) ([]lib.Paper, error) {
	results, err := utils.GetByQuery(stub, query)
	if err != nil {
		return nil, err
	}
	var papers []lib.Paper
	for _, result := range results {
		var paper lib.Paper
		_ = json.Unmarshal(result, &paper)
		papers = append(papers, paper)
	}
	return papers, nil
}

func PapersByQuery(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	query := args[0]
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	papersBytes, err := utils.MarshalByArgs(papers, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

func Papers(stub shim.ChaincodeStubInterface, args []string) peer.Response {
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
	papersBytes, err := utils.MarshalByArgs(papers, args)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

func PapersSortByTitle(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	query := `{"sort":[{"paper_title":"asc"}]}`
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	papersBytes, err := utils.MarshalByArgs(papers, args)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

func PapersSortByUploadTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	query := `{"sort":[{"paper_upload_time":"desc"}]}`
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	papersBytes, err := utils.MarshalByArgs(papers, args)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

func AcceptedPapersSortByTitle(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	query := `{"selector":{"paper_status":"accepted"},"sort":[{"paper_title":"asc"}]}`
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	papersBytes, err := utils.MarshalByArgs(papers, args)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

func AcceptedPapersSortByUploadTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	query := `{"selector":{"paper_status":"accepted"},"sort":[{"paper_upload_time":"asc"}]}`
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	papersBytes, err := utils.MarshalByArgs(papers, args)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

func AcceptedPapersSortByPublishTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	query := `{"selector":{"paper_status":"accepted"},"sort":[{"paper_publish_time":"asc"}]}`
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	papersBytes, err := utils.MarshalByArgs(papers, args)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

func RejectedPapersSortByTitle(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	query := `{"selector":{"paper_status":"rejected"},"sort":[{"paper_title":"asc"}]}`
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	papersBytes, err := utils.MarshalByArgs(papers, args)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

func RejectedPapersSortByUploadTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	query := `{"selector":{"paper_status":"rejected"},"sort":[{"paper_upload_time":"desc"}]}`
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	papersBytes, err := utils.MarshalByArgs(papers, args)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

func ReviewingPapersSortByTitle(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	query := `{"selector":{"paper_status":"reviewing"},"sort":[{"paper_title":"asc"}]}`
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	papersBytes, err := utils.MarshalByArgs(papers, args)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

func ReviewingPapersSortByUploadTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	query := `{"selector":{"paper_status":"reviewing"},"sort":[{"paper_upload_time":"desc"}]}`
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	papersBytes, err := utils.MarshalByArgs(papers, args)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

func AcceptedPapersByTitleSortByTitle(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	title := args[0]
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_title":{"$regex":".*?%s.*?"}},{"paper_status":"accepted"}]},"sort":[{"paper_title":"asc"}]}`, title)
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	papersBytes, err := utils.MarshalByArgs(papers, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

func AcceptedPapersByTitleSortByPublishTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	title := args[0]
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_title":{"$regex":".*?%s.*?"}},{"paper_status":"accepted"}]},"sort":[{"paper_publish_time":"desc"}]}`, title)
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	papersBytes, err := utils.MarshalByArgs(papers, args)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

func AcceptedPapersByAuthorSortByTitle(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	author := args[0]
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_authors":{"$in":["%s"]}},{"paper_status":"accepted"}]},"sort":[{"paper_title":"asc"}]}`, author)
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	papersBytes, err := utils.MarshalByArgs(papers, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

func AcceptedPapersByAuthorSortByPublishTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	author := args[0]
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_authors":{"$in":["%s"]}},{"paper_status":"accepted"}]},"sort":[{"paper_publish_time":"desc"}]}`, author)
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	papersBytes, err := utils.MarshalByArgs(papers, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

func AcceptedPapersByKeywordSortByTitle(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	keyword := args[0]
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_keywords":{"$in":["%s"]}},{"paper_status":"accepted"}]},"sort":[{"paper_title":"asc"}]}`, keyword)
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	papersBytes, err := utils.MarshalByArgs(papers, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

func AcceptedPapersByKeywordSortByPublishTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	keyword := args[0]
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_keywords":{"$in":["%s"]}},{"paper_status":"accepted"}]},"sort":[{"paper_publish_time":"desc"}]}`, keyword)
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	papersBytes, err := utils.MarshalByArgs(papers, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

func AcceptedPapersByUploaderSortByTitle(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	uploader := args[0]
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_uploader":"%s"},{"paper_status":"accepted"}]},"sort":[{"paper_title":"asc"}]}`, uploader)
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	papersBytes, err := utils.MarshalByArgs(papers, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

func AcceptedPapersByUploaderSortByUploadTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	uploader := args[0]
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_uploader":"%s"},{"paper_status":"accepted"}]},"sort":[{"paper_upload_time":"desc"}]}`, uploader)
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	papersBytes, err := utils.MarshalByArgs(papers, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

func AcceptedPapersByUploaderSortByPublishTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	uploader := args[0]
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_uploader":"%s"},{"paper_status":"accepted"}]},"sort":[{"paper_publish_time":"desc"}]}`, uploader)
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	papersBytes, err := utils.MarshalByArgs(papers, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

func RejectedPapersByUploaderSortByTitle(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	uploader := args[0]
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_uploader":"%s"},{"paper_status":"rejected"}]},"sort":[{"paper_title":"asc"}]}`, uploader)
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	papersBytes, err := utils.MarshalByArgs(papers, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

func RejectedPapersByUploaderSortByUploadTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	uploader := args[0]
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_uploader":"%s"},{"paper_status":"rejected"}]},"sort":[{"paper_upload_time":"desc"}]}`, uploader)
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	papersBytes, err := utils.MarshalByArgs(papers, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

func RejectedPapersByUploaderSortByPublishTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	uploader := args[0]
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_uploader":"%s"},{"paper_status":"rejected"}]},"sort":[{"paper_publish_time":"desc"}]}`, uploader)
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	papersBytes, err := utils.MarshalByArgs(papers, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

func ReviewingPapersByUploaderSortByTitle(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	uploader := args[0]
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_uploader":"%s"},{"paper_status":"reviewing"}]},"sort":[{"paper_title":"asc"}]}`, uploader)
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	papersBytes, err := utils.MarshalByArgs(papers, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

func ReviewingPapersByUploaderSortByUploadTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	uploader := args[0]
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_uploader":"%s"},{"paper_status":"reviewing"}]},"sort":[{"paper_upload_time":"desc"}]}`, uploader)
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	papersBytes, err := utils.MarshalByArgs(papers, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
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
