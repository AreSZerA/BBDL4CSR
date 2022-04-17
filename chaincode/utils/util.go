package utils

import (
	"chaincode/lib"
	"encoding/json"
	"strconv"
)

func CheckArgs(args []string, length int) error {
	if len(args) < length {
		return lib.ErrArgsLength
	}
	for _, arg := range args {
		if arg == "" {
			return lib.ErrArgsEmpty
		}
	}
	return nil
}

func GetStatus(s1 string, s2 string, s3 string) string {
	switch lib.StatusCode(s1) + lib.StatusCode(s2) + lib.StatusCode(s3) {
	case 3 * lib.StatusCodeAccepted, 2*lib.StatusCodeAccepted + lib.StatusCodeRejected:
		return lib.StatusCodeAccepted.String()
	case 3 * lib.StatusCodeRejected, 2*lib.StatusCodeRejected + lib.StatusCodeAccepted:
		return lib.StatusCodeRejected.String()
	default:
		return lib.StatusCodeReviewing.String()
	}
}

func MarshalWithOffsetAndLimit(slice interface{}, offsetString string, limitString string) ([]byte, error) {
	var objectsBytes []byte
	offset, err := strconv.Atoi(offsetString)
	if err != nil {
		return nil, lib.ErrOffsetNotInteger
	}
	if offset < 0 {
		offset = 0
	}
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		return nil, lib.ErrLimitNotInteger
	}
	if limit < 0 {
		limit = 0
	}
	switch slice.(type) {
	case []lib.User:
		if offset > len(slice.([]lib.User)) {
			objectsBytes = []byte("[]")
		} else {
			if offset+limit > len(slice.([]lib.User)) {
				objectsBytes, _ = json.Marshal(slice.([]lib.User)[offset:])
			} else {
				objectsBytes, _ = json.Marshal(slice.([]lib.User)[offset : offset+limit])
			}
		}
	case []lib.Paper:
		if offset > len(slice.([]lib.Paper)) {
			objectsBytes = []byte("[]")
		} else {
			if offset+limit > len(slice.([]lib.Paper)) {
				objectsBytes, _ = json.Marshal(slice.([]lib.Paper)[offset:])
			} else {
				objectsBytes, _ = json.Marshal(slice.([]lib.Paper)[offset : offset+limit])
			}
		}
	case []lib.PeerReview:
		if offset > len(slice.([]lib.PeerReview)) {
			objectsBytes = []byte("[]")
		} else {
			if offset+limit > len(slice.([]lib.PeerReview)) {
				objectsBytes, _ = json.Marshal(slice.([]lib.PeerReview)[offset:])
			} else {
				objectsBytes, _ = json.Marshal(slice.([]lib.PeerReview)[offset : offset+limit])
			}
		}
	default:
		return nil, lib.ErrNotBlockchainObjectSlice
	}
	return objectsBytes, nil
}
