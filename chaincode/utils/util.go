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

func MarshalByArgs(slice interface{}, args []string) ([]byte, error) {
	var result []byte
	switch slice.(type) {
	case []lib.User:
		if len(args) == 0 {
			result, _ = json.Marshal(slice.([]lib.User))
		} else if len(args) == 1 {
			if args[0] == "count" {
				result = []byte(strconv.Itoa(len(slice.([]lib.User))))
			} else {
				result, _ = json.Marshal(slice.([]lib.User))
			}
		} else {
			if args[0] == "count" {
				result = []byte(strconv.Itoa(len(slice.([]lib.User))))
			} else {
				offset, err := strconv.Atoi(args[0])
				if err != nil {
					return nil, lib.ErrOffsetNotInteger
				}
				if offset < 0 {
					offset = 0
				}
				limit, err := strconv.Atoi(args[1])
				if err != nil {
					return nil, lib.ErrLimitNotInteger
				}
				if limit < 0 {
					limit = 0
				}
				if offset > len(slice.([]lib.User)) {
					result = []byte("[]")
				} else {
					if offset+limit > len(slice.([]lib.User)) {
						result, _ = json.Marshal(slice.([]lib.User)[offset:])
					} else {
						result, _ = json.Marshal(slice.([]lib.User)[offset : offset+limit])
					}
				}
			}
		}
	case []lib.Paper:
		if len(args) == 0 {
			result, _ = json.Marshal(slice.([]lib.Paper))
		} else if len(args) == 1 {
			if args[0] == "count" {
				result = []byte(strconv.Itoa(len(slice.([]lib.Paper))))
			} else {
				result, _ = json.Marshal(slice.([]lib.Paper))
			}
		} else {
			if args[0] == "count" {
				result = []byte(strconv.Itoa(len(slice.([]lib.Paper))))
			} else {
				offset, err := strconv.Atoi(args[0])
				if err != nil {
					return nil, lib.ErrOffsetNotInteger
				}
				if offset < 0 {
					offset = 0
				}
				limit, err := strconv.Atoi(args[1])
				if err != nil {
					return nil, lib.ErrLimitNotInteger
				}
				if limit < 0 {
					limit = 0
				}
				if offset > len(slice.([]lib.Paper)) {
					result = []byte("[]")
				} else {
					if offset+limit > len(slice.([]lib.Paper)) {
						result, _ = json.Marshal(slice.([]lib.Paper)[offset:])
					} else {
						result, _ = json.Marshal(slice.([]lib.Paper)[offset : offset+limit])
					}
				}
			}
		}
	case []lib.PeerReview:
		if len(args) == 0 {
			result, _ = json.Marshal(slice.([]lib.PeerReview))
		} else if len(args) == 1 {
			if args[0] == "count" {
				result = []byte(strconv.Itoa(len(slice.([]lib.PeerReview))))
			} else {
				result, _ = json.Marshal(slice.([]lib.PeerReview))
			}
		} else {
			if args[0] == "count" {
				result = []byte(strconv.Itoa(len(slice.([]lib.PeerReview))))
			} else {
				offset, err := strconv.Atoi(args[0])
				if err != nil {
					return nil, lib.ErrOffsetNotInteger
				}
				if offset < 0 {
					offset = 0
				}
				limit, err := strconv.Atoi(args[1])
				if err != nil {
					return nil, lib.ErrLimitNotInteger
				}
				if limit < 0 {
					limit = 0
				}
				if offset > len(slice.([]lib.PeerReview)) {
					result = []byte("[]")
				} else {
					if offset+limit > len(slice.([]lib.PeerReview)) {
						result, _ = json.Marshal(slice.([]lib.PeerReview)[offset:])
					} else {
						result, _ = json.Marshal(slice.([]lib.PeerReview)[offset : offset+limit])
					}
				}
			}
		}
	default:
		return nil, lib.ErrNotBlockchainObjectSlice
	}
	return result, nil
}
