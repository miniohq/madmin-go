//
//  MinIO Inc [madmin-go]
//  Copyright (c) 2014-2025 MinIO.
//  All rights reserved. No warranty, explicit or implicit, provided.
//

package madmin

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

//msgp:clearomitted
//go:generate msgp

// BucketScanInfo contains information of a bucket scan in a given pool/set
type BucketScanInfo struct {
	Pool        int         `msg:"pool"`
	Set         int         `msg:"set"`
	Cycle       uint64      `msg:"cycle"`
	Ongoing     bool        `msg:"ongoing"`
	LastUpdate  time.Time   `msg:"last_update"`
	LastStarted time.Time   `msg:"last_started"`
	Completed   []time.Time `msg:"completed,omitempty"`
}

// BucketScanInfo returns information of a bucket scan in all pools/sets
func (adm *AdminClient) BucketScanInfo(ctx context.Context, bucket string) ([]BucketScanInfo, error) {
	resp, err := adm.executeMethod(ctx,
		http.MethodGet,
		requestData{relPath: adminAPIPrefix + "/scanner/status/" + bucket})
	if err != nil {
		return nil, err
	}
	defer closeResponse(resp)

	if resp.StatusCode != http.StatusOK {
		return nil, httpRespToErrorResponse(resp)
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var info []BucketScanInfo
	err = json.Unmarshal(respBytes, &info)
	if err != nil {
		return nil, err
	}

	return info, nil
}
