//
//  MinIO Inc [madmin-go]
//  Copyright (c) 2014-2025 MinIO.
//  All rights reserved. No warranty, explicit or implicit, provided.
//

//go:build !linux
// +build !linux

package madmin

// GetNetInfo returns information of the given network interface
// Not implemented for non-linux platforms
func GetNetInfo(addr string, iface string) NetInfo {
	return NetInfo{
		NodeCommon: NodeCommon{
			Addr:  addr,
			Error: "Not implemented for non-linux platforms",
		},
		Interface: iface,
	}
}
