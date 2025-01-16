//
//  MinIO Inc [madmin-go]
//  Copyright (c) 2014-2025 MinIO.
//  All rights reserved. No warranty, explicit or implicit, provided.
//

//go:build linux
// +build linux

package madmin

import (
	"fmt"

	"github.com/safchain/ethtool"
)

// GetNetInfo returns information of the given network interface
func GetNetInfo(addr string, iface string) (ni NetInfo) {
	ni.Addr = addr
	ni.Interface = iface

	ethHandle, err := ethtool.NewEthtool()
	if err != nil {
		ni.Error = err.Error()
		return
	}
	defer ethHandle.Close()

	di, err := ethHandle.DriverInfo(ni.Interface)
	if err != nil {
		ni.Error = fmt.Sprintf("Error getting driver info for %s: %s", ni.Interface, err.Error())
		return
	}

	ni.Driver = di.Driver
	ni.FirmwareVersion = di.FwVersion

	return
}
