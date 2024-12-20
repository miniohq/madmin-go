//
// MinIO Inc [madmin-go]
// Copyright (c) 2014-2024 MinIO.
// All rights reserved. No warranty, explicit or implicit, provided.
//

//go:build !linux
// +build !linux

package kernel

// VersionFromRelease only implemented on Linux.
func VersionFromRelease(_ string) (uint32, error) {
	return 0, nil
}

// Version only implemented on Linux.
func Version(_, _, _ int) uint32 {
	return 0
}

// CurrentRelease only implemented on Linux.
func CurrentRelease() (string, error) {
	return "", nil
}

// CurrentVersion only implemented on Linux.
func CurrentVersion() (uint32, error) {
	return 0, nil
}
