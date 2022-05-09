// Copyright 2022 AreSZerA. All rights reserved.
// This file defines functions the paper DAO and functions to update and query.

package models

import (
	"ClaytonUniversityLibrary/blockchain"
	"encoding/json"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"strconv"
	"time"
)

type jsonPaper struct {
	ID          string    `json:"paper_id"`
	Uploader    string    `json:"paper_uploader"`
	UploadTime  int64     `json:"paper_upload_time"`
	Title       string    `json:"paper_title"`
	Abstract    string    `json:"paper_abstract"`
	Authors     []string  `json:"paper_authors"`
	Keywords    []string  `json:"paper_keywords"`
	Reviewers   [3]string `json:"paper_reviewers"`
	Status      string    `json:"paper_status"`
	PublishTime int64     `json:"paper_publish_time"`
}

func (p jsonPaper) convert() Paper {
	authors := ""
	for i := 0; i < len(p.Authors); i++ {
		authors += p.Authors[i]
		if i < len(p.Authors)-1 {
			authors += ", "
		}
	}
	return Paper{
		ID:          p.ID,
		Uploader:    p.Uploader,
		UploadTime:  time.UnixMicro(p.UploadTime / 1000).Format("2006-01-02 15:04:05"),
		Title:       p.Title,
		Abstract:    p.Abstract,
		Authors:     authors,
		Keywords:    p.Keywords,
		Status:      p.Status,
		PublishTime: time.UnixMicro(p.PublishTime / 1000).Format("2006-01-02 15:04:05"),
		FirstChar:   string(authors[0]),
	}
}

type Paper struct {
	ID          string
	Uploader    string
	UploadTime  string
	Title       string
	Abstract    string
	Authors     string
	Keywords    []string
	Status      string
	PublishTime string
	FirstChar   string
}

func CountPublishedPapers() (int, error) {
	n, err := blockchain.Query(blockchain.FuncRetrieveAcceptedPapersSortByPublishTime)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(n.Payload))
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
	var jsonPapers []jsonPaper
	err = json.Unmarshal(resp.Payload, &jsonPapers)
	if err != nil {
		return nil, err
	}
	var papers []Paper
	for _, paper := range jsonPapers {
		papers = append(papers, paper.convert())
	}
	return papers, nil
}

func UploadPaper(email string, title string, abstract string, authors string, keywords string) (Paper, error) {
	resp, err := blockchain.Execute(blockchain.FuncCreatePaper, []byte(email), []byte(title), []byte(abstract), []byte(authors), []byte(keywords))
	if err != nil {
		return Paper{}, err
	}
	var paper jsonPaper
	err = json.Unmarshal(resp.Payload, &paper)
	if err != nil {
		return Paper{}, err
	}
	return paper.convert(), nil
}
