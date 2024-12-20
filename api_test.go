//
// MinIO Inc [madmin-go]
// Copyright (c) 2014-2024 MinIO.
// All rights reserved. No warranty, explicit or implicit, provided.
//

// Package madmin_test
package madmin_test

import (
	"testing"

	"github.com/minio/madmin-go/v3"
)

func TestMinioAdminClient(t *testing.T) {
	_, err := madmin.New("localhost:9000", "food", "food123", true)
	if err != nil {
		t.Fatal(err)
	}
}
