//
// MinIO Inc [madmin-go]
// Copyright (c) 2014-2024 MinIO.
// All rights reserved. No warranty, explicit or implicit, provided.
//

package madmin

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

// NetperfNodeResult - stats from each server
type NetperfNodeResult struct {
	Endpoint string `json:"endpoint"`
	TX       uint64 `json:"tx"`
	RX       uint64 `json:"rx"`
	Error    string `json:"error,omitempty"`
}

// NetperfResult - aggregate results from all servers
type NetperfResult struct {
	NodeResults []NetperfNodeResult `json:"nodeResults"`
}

// Netperf - perform netperf on the MinIO servers
func (adm *AdminClient) Netperf(ctx context.Context, duration time.Duration) (result NetperfResult, err error) {
	queryVals := make(url.Values)
	queryVals.Set("duration", duration.String())

	resp, err := adm.executeMethod(ctx,
		http.MethodPost, requestData{
			relPath:     adminAPIPrefix + "/speedtest/net",
			queryValues: queryVals,
		})
	if err != nil {
		return result, err
	}
	if resp.StatusCode != http.StatusOK {
		return result, httpRespToErrorResponse(resp)
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}
