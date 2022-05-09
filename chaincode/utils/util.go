// Copyright 2022 AreSZerA. All rights reserved.
// This file defines some utility functions.

package utils

import (
	"chaincode/lib"
	"encoding/json"
	"strconv"
)

// CheckArgs accepts an argument slice and an integer as the acceptable length,
// returns lib.ErrArgsLength when the number of arguments is not enough,
// lib.ErrArgsEmpty when exists empty augments.
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

// GetStatus returns lib.StatusReviewing or lib.StatusAccepted or lib.StatusRejected according to the passed strings.
func GetStatus(s1 string, s2 string, s3 string) string {
	// convert string statuses to numeric statuses, return the final status according to their sum
	// Accepted: 4 + 4 + 4 = 12, or 4 + 4 + 2 = 10
	// Rejected: 2 + 2 + 2 = 6, or 2 + 2 + 4 = 8
	// any other combinations will not obtain the above results
	switch lib.StatusCode(s1) + lib.StatusCode(s2) + lib.StatusCode(s3) {
	case 3 * lib.StatusCodeAccepted, 2*lib.StatusCodeAccepted + lib.StatusCodeRejected:
		return lib.StatusCodeAccepted.String()
	case 3 * lib.StatusCodeRejected, 2*lib.StatusCodeRejected + lib.StatusCodeAccepted:
		return lib.StatusCodeRejected.String()
	default:
		return lib.StatusCodeReviewing.String()
	}
}

// MarshalByArgs serialises or count the lib.BlockchainObjectInterface slice in accordance with arguments.
func MarshalByArgs(slice interface{}, args []string) ([]byte, error) {
	var result []byte
	// do different operations for different types
	// actually, for []lib.User, []lib.Papers, and []lib.PeerReview, the ideas and procedures are the same
	//
	// Q: Why not use []interface{} or []lib.BlockchainObjectInterface as the first argument type?
	// A: []lib.User, []lib.Paper, and []lib.PeerReview are not []lib.BlockchainObjectInterface.
	//    To convert them to []lib.BlockchainObjectInterface, it may cost a lot to traverse the slice.
	//
	// Q: Why not use []lib.BlockchainObjectInterface to replace other cases in switch?
	// A: The reason is the same. If used so, this function always returns lib.ErrNotBlockchainObjectSlice.
	switch slice.(type) {
	case []lib.User:
		// when the length of the arguments is 0, that is, no extra arguments in API
		if len(args) == 0 {
			// serialise the original slice to return
			result, _ = json.Marshal(slice.([]lib.User))
		} else if len(args) == 1 {
			// when passed one argument and the argument is "count"
			if args[0] == "count" {
				// the result is a numeric value
				result = []byte(strconv.Itoa(len(slice.([]lib.User))))
			} else {
				// otherwise, ignore the argument
				result, _ = json.Marshal(slice.([]lib.User))
			}
		} else {
			// when passed more argument and the argument is "count"
			if args[0] == "count" {
				result = []byte(strconv.Itoa(len(slice.([]lib.User))))
			} else {
				// otherwise, enter pagination mode
				// use the first argument as offset and the second as limit, return error if any one of them is NaN
				// if the offset or limit is negative, set it as 0
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
				// when the offset is larger than the number of objects, the result is an empty JSON array
				if offset > len(slice.([]lib.User)) {
					result = []byte("[]")
				} else {
					// otherwise, check if offset+limit is greater than the slice length
					// if so, cut the slice from offset to the end
					if offset+limit > len(slice.([]lib.User)) {
						result, _ = json.Marshal(slice.([]lib.User)[offset:])
					} else {
						// otherwise, cut the slice from offset to offset+limit
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
		// return error when the type is not acceptable
		return nil, lib.ErrNotBlockchainObjectSlice
	}
	return result, nil
}
