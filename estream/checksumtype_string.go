//
//  MinIO Inc [madmin-go]
//  Copyright (c) 2014-2025 MinIO.
//  All rights reserved. No warranty, explicit or implicit, provided.
//

// Code generated by "stringer -type=checksumType -trimprefix=checksumType"; DO NOT EDIT.

package estream

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[checksumTypeNone-0]
	_ = x[checksumTypeXxhash-1]
	_ = x[checksumTypeUnknown-2]
}

const _checksumType_name = "NoneXxhashUnknown"

var _checksumType_index = [...]uint8{0, 4, 10, 17}

func (i checksumType) String() string {
	if i >= checksumType(len(_checksumType_index)-1) {
		return "checksumType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _checksumType_name[_checksumType_index[i]:_checksumType_index[i+1]]
}
