//
// MinIO Inc [madmin-go]
// Copyright (c) 2014-2024 MinIO.
// All rights reserved. No warranty, explicit or implicit, provided.
//

//go:build (linux && arm) || (linux && ppc64) || (linux && ppc64le) || (linux && s390x) || (linux && riscv64)
// +build linux,arm linux,ppc64 linux,ppc64le linux,s390x linux,riscv64

package kernel

func utsnameStr(in []uint8) string {
	out := make([]byte, 0, len(in))
	for i := 0; i < len(in); i++ {
		if in[i] == 0x00 {
			break
		}
		out = append(out, byte(in[i]))
	}
	return string(out)
}
