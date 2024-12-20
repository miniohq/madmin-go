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
	"strconv"
)

// ServerPeerUpdateStatus server update peer binary update result
type ServerPeerUpdateStatus struct {
	Host           string                 `json:"host"`
	Err            string                 `json:"err,omitempty"`
	CurrentVersion string                 `json:"currentVersion"`
	UpdatedVersion string                 `json:"updatedVersion"`
	WaitingDrives  map[string]DiskMetrics `json:"waitingDrives,omitempty"`
}

// ServerUpdateStatusV2 server update status
type ServerUpdateStatusV2 struct {
	DryRun  bool                     `json:"dryRun"`
	Results []ServerPeerUpdateStatus `json:"results,omitempty"`
}

// ServerUpdateOpts specifies the URL (optionally to download the binary from)
// also allows a dry-run, the new API is idempotent which means you can
// run it as many times as you want and any server that is not upgraded
// automatically does get upgraded eventually to the relevant version.
type ServerUpdateOpts struct {
	UpdateURL string
	DryRun    bool
}

// ServerUpdateV2 - updates and restarts the MinIO cluster to latest version.
// optionally takes an input URL to specify a custom update binary link
func (adm *AdminClient) ServerUpdateV2(ctx context.Context, opts ServerUpdateOpts) (us ServerUpdateStatusV2, err error) {
	queryValues := url.Values{}
	queryValues.Set("type", "2")
	queryValues.Set("updateURL", opts.UpdateURL)
	queryValues.Set("dry-run", strconv.FormatBool(opts.DryRun))

	// Request API to Restart server
	resp, err := adm.executeMethod(ctx,
		http.MethodPost, requestData{
			relPath:     adminAPIPrefix + "/update",
			queryValues: queryValues,
		},
	)
	defer closeResponse(resp)
	if err != nil {
		return us, err
	}

	if resp.StatusCode != http.StatusOK {
		return us, httpRespToErrorResponse(resp)
	}

	if err = json.NewDecoder(resp.Body).Decode(&us); err != nil {
		return us, err
	}

	return us, nil
}

// ServerUpdateStatus - contains the response of service update API
type ServerUpdateStatus struct {
	// Deprecated: this struct is fully deprecated since Jan 2024.
	CurrentVersion string `json:"currentVersion"`
	UpdatedVersion string `json:"updatedVersion"`
}

// ServerUpdate - updates and restarts the MinIO cluster to latest version.
// optionally takes an input URL to specify a custom update binary link
func (adm *AdminClient) ServerUpdate(ctx context.Context, updateURL string) (us ServerUpdateStatus, err error) {
	queryValues := url.Values{}
	queryValues.Set("updateURL", updateURL)

	// Request API to Restart server
	resp, err := adm.executeMethod(ctx,
		http.MethodPost, requestData{
			relPath:     adminAPIPrefix + "/update",
			queryValues: queryValues,
		},
	)
	defer closeResponse(resp)
	if err != nil {
		return us, err
	}

	if resp.StatusCode != http.StatusOK {
		return us, httpRespToErrorResponse(resp)
	}

	if err = json.NewDecoder(resp.Body).Decode(&us); err != nil {
		return us, err
	}

	return us, nil
}
