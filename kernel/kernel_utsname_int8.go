//
// MinIO Inc [madmin-go]
// Copyright (c) 2014-2024 MinIO.
// All rights reserved. No warranty, explicit or implicit, provided.
//

//go:build (linux && 386) || (linux && amd64) || (linux && arm64) || (linux && loong64) || (linux && mips64) || (linux && mips64le) || (linux && mips)
// +build linux,386 linux,amd64 linux,arm64 linux,loong64 linux,mips64 linux,mips64le linux,mips

package kernel

func utsnameStr(in []int8) string {
	out := make([]byte, 0, len(in))
	for i := 0; i < len(in); i++ {
		if in[i] == 0x00 {
			break
		}
		out = append(out, byte(in[i]))
	}
	return string(out)
}
