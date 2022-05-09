// Copyright 2022 AreSZerA. All rights reserved.
// This file defines errors for package route to use.

package lib

import "errors"

var (
	ErrArgsLength               = errors.New("argument length too shorty")
	ErrArgsEmpty                = errors.New("empty argument contained")
	ErrUserNotFound             = errors.New("user not found")
	ErrUserExists               = errors.New("user exists")
	ErrReviewerNotEnough        = errors.New("reviewer not enough")
	ErrPaperNotFound            = errors.New("paper not found")
	ErrPeerReviewNotFound       = errors.New("peer review not found")
	ErrLimitNotInteger          = errors.New("limit is not integer")
	ErrOffsetNotInteger         = errors.New("offset is not integer")
	ErrNotBlockchainObjectSlice = errors.New("not blockchain object slice")
	ErrDataHasLost              = errors.New("data has lost")
	ErrPeerReviewDone           = errors.New("peer review is done")
)
