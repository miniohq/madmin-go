//
// MinIO Inc [madmin-go]
// Copyright (c) 2014-2024 MinIO.
// All rights reserved. No warranty, explicit or implicit, provided.
//

package madmin

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// ExportBucketMetadata makes an admin call to export bucket metadata of a bucket
func (adm *AdminClient) ExportBucketMetadata(ctx context.Context, bucket string) (io.ReadCloser, error) {
	path := adminAPIPrefix + "/export-bucket-metadata"
	queryValues := url.Values{}
	queryValues.Set("bucket", bucket)

	resp, err := adm.executeMethod(ctx,
		http.MethodGet, requestData{
			relPath:     path,
			queryValues: queryValues,
		},
	)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		closeResponse(resp)
		return nil, httpRespToErrorResponse(resp)
	}
	return resp.Body, nil
}

// MetaStatus status of metadata import
type MetaStatus struct {
	IsSet bool   `json:"isSet"`
	Err   string `json:"error,omitempty"`
}

// BucketStatus reflects status of bucket metadata import
type BucketStatus struct {
	ObjectLock   MetaStatus `json:"olock"`
	Versioning   MetaStatus `json:"versioning"`
	Policy       MetaStatus `json:"policy"`
	Tagging      MetaStatus `json:"tagging"`
	SSEConfig    MetaStatus `json:"sse"`
	Lifecycle    MetaStatus `json:"lifecycle"`
	Notification MetaStatus `json:"notification"`
	Quota        MetaStatus `json:"quota"`
	Cors         MetaStatus `json:"cors"`
	Err          string     `json:"error,omitempty"`
}

// BucketMetaImportErrs reports on bucket metadata import status.
type BucketMetaImportErrs struct {
	Buckets map[string]BucketStatus `json:"buckets,omitempty"`
}

// ImportBucketMetadata makes an admin call to set bucket metadata of a bucket from imported content
func (adm *AdminClient) ImportBucketMetadata(ctx context.Context, bucket string, contentReader io.ReadCloser) (r BucketMetaImportErrs, err error) {
	content, err := io.ReadAll(contentReader)
	if err != nil {
		return r, err
	}

	path := adminAPIPrefix + "/import-bucket-metadata"
	queryValues := url.Values{}
	queryValues.Set("bucket", bucket)

	resp, err := adm.executeMethod(ctx,
		http.MethodPut, requestData{
			relPath:     path,
			queryValues: queryValues,
			content:     content,
		},
	)
	defer closeResponse(resp)

	if err != nil {
		return r, err
	}

	if resp.StatusCode != http.StatusOK {
		return r, httpRespToErrorResponse(resp)
	}

	err = json.NewDecoder(resp.Body).Decode(&r)
	return r, err
}
