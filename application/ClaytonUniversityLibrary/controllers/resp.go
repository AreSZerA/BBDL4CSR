// Copyright 2022 AreSZerA. All rights reserved.
// This file defines the Resp struct to simplify the message constructing procedure.

package controllers

import (
	"encoding/json"
	"time"
)

type Resp struct {
	// Current timestamp.
	Time int64 `json:"time,omitempty"`
	// Main content of the response message
	Result interface{} `json:"result,omitempty"`
	// Error message.
	Error string `json:"error,omitempty"`
}

// NewResp constructs Resp struct then serializes it to string.
func NewResp(result interface{}, err string) string {
	resp := Resp{Time: time.Now().UnixNano(), Result: result, Error: err}
	respBytes, _ := json.Marshal(resp)
	return string(respBytes)
}
