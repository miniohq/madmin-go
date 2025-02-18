//
//  MinIO Inc [madmin-go]
//  Copyright (c) 2014-2025 MinIO.
//  All rights reserved. No warranty, explicit or implicit, provided.
//

package madmin

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// QuotaType represents bucket quota type
type QuotaType string

const (
	// HardQuota specifies a hard quota of usage for bucket
	HardQuota QuotaType = "hard"
)

// IsValid returns true if quota type is one of Hard
func (t QuotaType) IsValid() bool {
	return t == HardQuota
}

// BucketQuota holds bucket quota restrictions
type BucketQuota struct {
	Quota    uint64    `json:"quota"`    // Deprecated Aug 2023
	Size     uint64    `json:"size"`     // Indicates maximum size allowed per bucket
	Rate     uint64    `json:"rate"`     // Indicates bandwidth rate allocated per bucket
	Requests uint64    `json:"requests"` // Indicates number of requests allocated per bucket
	Type     QuotaType `json:"quotatype,omitempty"`
}

// IsValid returns false if quota is invalid
// empty quota when Quota == 0 is always true.
func (q BucketQuota) IsValid() bool {
	if q.Quota > 0 {
		return q.Type.IsValid()
	}
	// Empty configs are valid.
	return true
}

// GetBucketQuota - get info on a user
func (adm *AdminClient) GetBucketQuota(ctx context.Context, bucket string) (q BucketQuota, err error) {
	queryValues := url.Values{}
	queryValues.Set("bucket", bucket)

	reqData := requestData{
		relPath:     adminAPIPrefix + "/get-bucket-quota",
		queryValues: queryValues,
	}

	// Execute GET on /minio/admin/v3/get-quota
	resp, err := adm.executeMethod(ctx, http.MethodGet, reqData)

	defer closeResponse(resp)
	if err != nil {
		return q, err
	}

	if resp.StatusCode != http.StatusOK {
		return q, httpRespToErrorResponse(resp)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return q, err
	}
	if err = json.Unmarshal(b, &q); err != nil {
		return q, err
	}

	return q, nil
}

// SetBucketQuota - sets a bucket's quota, if quota is set to '0'
// quota is disabled.
func (adm *AdminClient) SetBucketQuota(ctx context.Context, bucket string, quota *BucketQuota) error {
	data, err := json.Marshal(quota)
	if err != nil {
		return err
	}

	queryValues := url.Values{}
	queryValues.Set("bucket", bucket)

	reqData := requestData{
		relPath:     adminAPIPrefix + "/set-bucket-quota",
		queryValues: queryValues,
		content:     data,
	}

	// Execute PUT on /minio/admin/v3/set-bucket-quota to set quota for a bucket.
	resp, err := adm.executeMethod(ctx, http.MethodPut, reqData)

	defer closeResponse(resp)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return httpRespToErrorResponse(resp)
	}

	return nil
}
