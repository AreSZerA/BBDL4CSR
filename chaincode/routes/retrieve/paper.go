// Copyright 2022 AreSZerA. All rights reserved.
// This file defines functions for retrieving papers.

package retrieve

import (
	"chaincode/lib"
	"chaincode/utils"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// queryPapers is used in this package to simplify the process of query.
func queryPapers(stub shim.ChaincodeStubInterface, query string) ([]lib.Paper, error) {
	// retrieve results by query
	results, err := utils.GetByQuery(stub, query)
	if err != nil {
		return nil, err
	}
	var papers []lib.Paper
	// traverse the results, deserialize and append to the slice
	for _, result := range results {
		var paper lib.Paper
		// deserialize the string to user object
		_ = json.Unmarshal(result, &paper)
		papers = append(papers, paper)
	}
	return papers, nil
}

// PapersByQuery in package `retrieve` retrieves papers by query.
// It requires one necessary argument: query.
func PapersByQuery(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	query := args[0]
	// query to get results
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	papersBytes, err := utils.MarshalByArgs(papers, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

// Papers in package `retrieve` retrieves all the papers.
// No necessary argument is required.
func Papers(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// get all the papers
	results, err := utils.GetAll(stub, lib.ObjectTypePaper)
	// traverse the results, deserialize and append to the slice
	if err != nil {
		return shim.Error(err.Error())
	}
	var papers []lib.Paper
	for _, result := range results {
		var paper lib.Paper
		// deserialize the string to paper object
		_ = json.Unmarshal(result, &paper)
		papers = append(papers, paper)
	}
	// check count and pagination
	papersBytes, err := utils.MarshalByArgs(papers, args)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

// PapersSortByTitle in package `retrieve` retrieves all the papers and sort by title in ascending order.
// No necessary argument is required.
func PapersSortByTitle(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// SELECT * FROM paper ORDER BY (title ASC)
	query := `{"sort":[{"paper_title":"asc"}]}`
	// query to get results
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	papersBytes, err := utils.MarshalByArgs(papers, args)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

// PapersSortByUploadTime in package `retrieve` retrieves all the papers and sort by upload time in descending order.
// No necessary argument is required.
func PapersSortByUploadTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// SELECT * FROM paper ORDER BY (upload_time DESC)
	query := `{"sort":[{"paper_upload_time":"desc"}]}`
	// query to get results
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	papersBytes, err := utils.MarshalByArgs(papers, args)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

// AcceptedPapersSortByTitle in package `retrieve` retrieves the accepted papers and sort by title in ascending order.
// No necessary argument is required.
func AcceptedPapersSortByTitle(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// SELECT * FROM paper WHERE status = accepted ORDER BY (title ASC)
	query := `{"selector":{"paper_status":"accepted"},"sort":[{"paper_title":"asc"}]}`
	// query to get results
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	papersBytes, err := utils.MarshalByArgs(papers, args)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

// AcceptedPapersSortByUploadTime in package `retrieve` retrieves the accepted papers and sort by upload time in descending order.
// No necessary argument is required.
func AcceptedPapersSortByUploadTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// SELECT * FROM paper WHERE status = accepted ORDER BY (upload_time DESC)
	query := `{"selector":{"paper_status":"accepted"},"sort":[{"paper_upload_time":"desc"}]}`
	// query to get results
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	papersBytes, err := utils.MarshalByArgs(papers, args)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

// AcceptedPapersSortByPublishTime in package `retrieve` retrieves the accepted papers and sort by publishing time in descending order.
// No necessary argument is required.
func AcceptedPapersSortByPublishTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// SELECT * FROM paper WHERE status = accepted ORDER BY (publish_time DESC)
	query := `{"selector":{"paper_status":"accepted"},"sort":[{"paper_publish_time":"desc"}]}`
	// query to get results
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	papersBytes, err := utils.MarshalByArgs(papers, args)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

// RejectedPapersSortByTitle in package `retrieve` retrieves the rejected papers and sort by title in ascending order.
// No necessary argument is required.
func RejectedPapersSortByTitle(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// SELECT * FROM paper WHERE status = rejected ORDER BY (title ASC)
	query := `{"selector":{"paper_status":"rejected"},"sort":[{"paper_title":"asc"}]}`
	// query to get results
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	papersBytes, err := utils.MarshalByArgs(papers, args)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

// RejectedPapersSortByUploadTime in package `retrieve` retrieves the rejected papers and sort by upload time in descending order.
// No necessary argument is required.
func RejectedPapersSortByUploadTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// SELECT * FROM paper WHERE status = rejected ORDER BY (upload_time DESC)
	query := `{"selector":{"paper_status":"rejected"},"sort":[{"paper_upload_time":"desc"}]}`
	// query to get results
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	papersBytes, err := utils.MarshalByArgs(papers, args)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

// ReviewingPapersSortByTitle in package `retrieve` retrieves the reviewing papers and sort by title in ascending order.
// No necessary argument is required.
func ReviewingPapersSortByTitle(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// SELECT * FROM paper WHERE status = reviewing ORDER BY (title ASC)
	query := `{"selector":{"paper_status":"reviewing"},"sort":[{"paper_title":"asc"}]}`
	// query to get results
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	papersBytes, err := utils.MarshalByArgs(papers, args)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

// ReviewingPapersSortByUploadTime in package `retrieve` retrieves the reviewing papers and sort by upload time in descending order.
// No necessary argument is required.
func ReviewingPapersSortByUploadTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// SELECT * FROM paper WHERE status = reviewing ORDER BY (upload_time DESC)
	query := `{"selector":{"paper_status":"reviewing"},"sort":[{"paper_upload_time":"desc"}]}`
	// query to get results
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	papersBytes, err := utils.MarshalByArgs(papers, args)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

// AcceptedPapersByTitleSortByTitle in package `retrieve` retrieves the accepted papers by title and sort by it in ascending order.
// It requires one necessary argument: title.
func AcceptedPapersByTitleSortByTitle(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	title := args[0]
	/// SELECT * FROM paper WHERE title LIKE %#{title}% AND status = accepted ORDER BY (title ASC)
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_title":{"$regex":".*?%s.*?"}},{"paper_status":"accepted"}]},"sort":[{"paper_title":"asc"}]}`, title)
	// query to get results
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	papersBytes, err := utils.MarshalByArgs(papers, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

// AcceptedPapersByTitleSortByPublishTime in package `retrieve` retrieves the accepted papers by title and sort by publishing time in descending order.
// It requires one necessary argument: title.
func AcceptedPapersByTitleSortByPublishTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	title := args[0]
	/// SELECT * FROM paper WHERE title LIKE %#{title}% AND status = accepted ORDER BY (publish_time DESC)
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_title":{"$regex":".*?%s.*?"}},{"paper_status":"accepted"}]},"sort":[{"paper_publish_time":"desc"}]}`, title)
	// query to get results
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	papersBytes, err := utils.MarshalByArgs(papers, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

// AcceptedPapersByAuthorSortByTitle in package `retrieve` retrieves the accepted papers by author and sort by title in ascending order.
// It requires one necessary argument: author.
func AcceptedPapersByAuthorSortByTitle(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	author := args[0]
	/// SELECT * FROM paper WHERE authors LIKE %#{author}% AND status = accepted ORDER BY (title ASC)
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_authors":{"$in":["%s"]}},{"paper_status":"accepted"}]},"sort":[{"paper_title":"asc"}]}`, author)
	// query to get results
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	papersBytes, err := utils.MarshalByArgs(papers, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

// AcceptedPapersByAuthorSortByPublishTime in package `retrieve` retrieves the accepted papers by author and sort by publishing time in descending order.
// It requires one necessary argument: author.
func AcceptedPapersByAuthorSortByPublishTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	author := args[0]
	/// SELECT * FROM paper WHERE authors LIKE %#{author}% AND status = accepted ORDER BY (publish_time DESC)
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_authors":{"$in":["%s"]}},{"paper_status":"accepted"}]},"sort":[{"paper_publish_time":"desc"}]}`, author)
	// query to get results
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	papersBytes, err := utils.MarshalByArgs(papers, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

// AcceptedPapersByKeywordSortByTitle in package `retrieve` retrieves the accepted papers by keyword and sort by title in ascending order.
// It requires one necessary argument: keyword.
func AcceptedPapersByKeywordSortByTitle(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	keyword := args[0]
	/// SELECT * FROM paper WHERE keywords LIKE %#{keyword}% AND status = accepted ORDER BY (title ASC)
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_keywords":{"$in":["%s"]}},{"paper_status":"accepted"}]},"sort":[{"paper_title":"asc"}]}`, keyword)
	// query to get results
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	papersBytes, err := utils.MarshalByArgs(papers, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

// AcceptedPapersByKeywordSortByPublishTime in package `retrieve` retrieves the accepted papers by keyword and sort by publishing time in descending order.
// It requires one necessary argument: keyword.
func AcceptedPapersByKeywordSortByPublishTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	keyword := args[0]
	/// SELECT * FROM paper WHERE keywords LIKE %#{keyword}% AND status = accepted ORDER BY (publish_time DESC)
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_keywords":{"$in":["%s"]}},{"paper_status":"accepted"}]},"sort":[{"paper_publish_time":"desc"}]}`, keyword)
	// query to get results
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	papersBytes, err := utils.MarshalByArgs(papers, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

// AcceptedPapersByUploaderSortByTitle in package `retrieve` retrieves the accepted papers by uploader email and sort by title in ascending order.
// It requires one necessary argument: uploader email.
func AcceptedPapersByUploaderSortByTitle(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	uploader := args[0]
	/// SELECT * FROM paper WHERE uploader = #{uploader} AND status = accepted ORDER BY (publish_time DESC)
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_uploader":"%s"},{"paper_status":"accepted"}]},"sort":[{"paper_title":"asc"}]}`, uploader)
	// query to get results
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	papersBytes, err := utils.MarshalByArgs(papers, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

// AcceptedPapersByUploaderSortByUploadTime in package `retrieve` retrieves the accepted papers by uploader email and sort by upload time in descending order.
// It requires one necessary argument: uploader email.
func AcceptedPapersByUploaderSortByUploadTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	uploader := args[0]
	/// SELECT * FROM paper WHERE uploader = #{uploader} AND status = accepted ORDER BY (upload_time DESC)
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_uploader":"%s"},{"paper_status":"accepted"}]},"sort":[{"paper_upload_time":"desc"}]}`, uploader)
	// query to get results
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	papersBytes, err := utils.MarshalByArgs(papers, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

// AcceptedPapersByUploaderSortByPublishTime in package `retrieve` retrieves the accepted papers by uploader email and sort by publishing time in descending order.
// It requires one necessary argument: uploader email.
func AcceptedPapersByUploaderSortByPublishTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	uploader := args[0]
	/// SELECT * FROM paper WHERE uploader = #{uploader} AND status = accepted ORDER BY (publish_time DESC)
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_uploader":"%s"},{"paper_status":"accepted"}]},"sort":[{"paper_publish_time":"desc"}]}`, uploader)
	// query to get results
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	papersBytes, err := utils.MarshalByArgs(papers, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

// RejectedPapersByUploaderSortByTitle in package `retrieve` retrieves the rejected papers by uploader email and sort by title in ascending order.
// It requires one necessary argument: uploader email.
func RejectedPapersByUploaderSortByTitle(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	uploader := args[0]
	/// SELECT * FROM paper WHERE uploader = #{uploader} AND status = rejected ORDER BY (title ASC)
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_uploader":"%s"},{"paper_status":"rejected"}]},"sort":[{"paper_title":"asc"}]}`, uploader)
	// query to get results
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	papersBytes, err := utils.MarshalByArgs(papers, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

// RejectedPapersByUploaderSortByUploadTime in package `retrieve` retrieves the rejected papers by uploader email and sort by upload time in descending order.
// It requires one necessary argument: uploader email.
func RejectedPapersByUploaderSortByUploadTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	uploader := args[0]
	/// SELECT * FROM paper WHERE uploader = #{uploader} AND status = rejected ORDER BY (upload_time DESC)
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_uploader":"%s"},{"paper_status":"rejected"}]},"sort":[{"paper_upload_time":"desc"}]}`, uploader)
	// query to get results
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	papersBytes, err := utils.MarshalByArgs(papers, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

// RejectedPapersByUploaderSortByPublishTime in package `retrieve` retrieves the rejected papers by uploader email and sort by publishing time in descending order.
// It requires one necessary argument: uploader email.
func RejectedPapersByUploaderSortByPublishTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	uploader := args[0]
	/// SELECT * FROM paper WHERE uploader = #{uploader} AND status = rejected ORDER BY (publish_time DESC)
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_uploader":"%s"},{"paper_status":"rejected"}]},"sort":[{"paper_publish_time":"desc"}]}`, uploader)
	// query to get results
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	papersBytes, err := utils.MarshalByArgs(papers, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

// ReviewingPapersByUploaderSortByTitle in package `retrieve` retrieves the reviewing papers by uploader email and sort by title in ascending order.
// It requires one necessary argument: uploader email.
func ReviewingPapersByUploaderSortByTitle(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	uploader := args[0]
	/// SELECT * FROM paper WHERE uploader = #{uploader} AND status = reviewing ORDER BY (title ASC)
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_uploader":"%s"},{"paper_status":"reviewing"}]},"sort":[{"paper_title":"asc"}]}`, uploader)
	// query to get results
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	papersBytes, err := utils.MarshalByArgs(papers, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

// ReviewingPapersByUploaderSortByUploadTime in package `retrieve` retrieves the reviewing papers by uploader email and sort by upload time in descending order.
// It requires one necessary argument: uploader email.
func ReviewingPapersByUploaderSortByUploadTime(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	uploader := args[0]
	/// SELECT * FROM paper WHERE uploader = #{uploader} AND status = reviewing ORDER BY (upload_time DESC)
	query := fmt.Sprintf(`{"selector":{"$and":[{"paper_uploader":"%s"},{"paper_status":"reviewing"}]},"sort":[{"paper_upload_time":"desc"}]}`, uploader)
	// query to get results
	papers, err := queryPapers(stub, query)
	if err != nil {
		return shim.Error(err.Error())
	}
	// check count and pagination
	papersBytes, err := utils.MarshalByArgs(papers, args[1:])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(papersBytes)
}

// PaperById in package `retrieve` retrieves the exact one paper by paper ID.
// It requires one necessary argument: paper ID.
func PaperById(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// check arguments
	err := utils.CheckArgs(args, 1)
	if err != nil {
		return shim.Error(err.Error())
	}
	id := args[0]
	// retrieve paper by composite key
	paperBytes, err := utils.GetByKeys(stub, lib.ObjectTypePaper, id)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(paperBytes)
}
