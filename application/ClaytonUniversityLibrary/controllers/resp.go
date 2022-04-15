package controllers

import (
	"encoding/json"
	"time"
)

type Resp struct {
	Time   int64       `json:"time,omitempty"`
	Result interface{} `json:"result,omitempty"`
	Error  string      `json:"error,omitempty"`
}

func NewResp(result interface{}, err string) string {
	resp := Resp{Time: time.Now().UnixNano(), Result: result, Error: err}
	respBytes, _ := json.Marshal(resp)
	return string(respBytes)
}
