//
//  MinIO Inc [madmin-go]
//  Copyright (c) 2014-2025 MinIO.
//  All rights reserved. No warranty, explicit or implicit, provided.
//

//go:build !linux
// +build !linux

package cgroup

import "errors"

// GetMemoryLimit - Not implemented in non-linux platforms
func GetMemoryLimit(_ int) (limit uint64, err error) {
	return limit, errors.New("Not implemented for non-linux platforms")
}
