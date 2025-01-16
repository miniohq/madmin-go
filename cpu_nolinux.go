//
//  MinIO Inc [madmin-go]
//  Copyright (c) 2014-2025 MinIO.
//  All rights reserved. No warranty, explicit or implicit, provided.
//

//go:build !linux
// +build !linux

package madmin

import (
	"errors"
)

func getCPUFreqStats() ([]CPUFreqStats, error) {
	return nil, errors.New("Not implemented for non-linux platforms")
}
