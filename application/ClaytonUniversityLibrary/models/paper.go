package models

import (
	"ClaytonUniversityLibrary/blockchain"
	"encoding/json"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

type Paper struct {
	ID         string    `json:"paper_id"`
	Uploader   string    `json:"paper_uploader"`
	UploadTime int64     `json:"paper_upload_time"`
	Title      string    `json:"paper_title"`
	Abstract   string    `json:"paper_abstract"`
	Authors    []string  `json:"paper_authors"`
	Keywords   []string  `json:"paper_keywords,omitempty"`
	Reviewers  [3]string `json:"paper_reviewers"`
	Status     string    `json:"paper_status"`
	ReviewTime int64     `json:"paper_review_time"`
}

func FindPublishedPapersByTitle(title string) ([]Paper, error) {
	var resp channel.Response
	var err error
	if title == "" {
		resp, err = blockchain.Query(blockchain.FuncRetrieveAcceptedPapersSortByPublishTime)
	} else {
		resp, err = blockchain.Query(blockchain.FuncRetrieveAcceptedPapersByTitleSortByPublishTime, []byte(title))
	}
	if err != nil {
		return nil, err
	}
	var papers []Paper
	err = json.Unmarshal(resp.Payload, &papers)
	if err != nil {
		return nil, err
	}
	return papers, nil
}

func UploadPaper(email string, title string, abstract string, authors string, keywords string) (Paper, error) {
	resp, err := blockchain.Execute(blockchain.FuncCreatePaper, []byte(email), []byte(title), []byte(abstract), []byte(authors), []byte(keywords))
	if err != nil {
		return Paper{}, err
	}
	var paper Paper
	err = json.Unmarshal(resp.Payload, &paper)
	if err != nil {
		return Paper{}, err
	}
	return paper, nil
}
