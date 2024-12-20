//
// MinIO Inc [madmin-go]
// Copyright (c) 2014-2024 MinIO.
// All rights reserved. No warranty, explicit or implicit, provided.
//

package madmin

import (
	"context"
	"net/http"
	"net/url"
)

// ServiceRestart - restarts the MinIO cluster
//
// Deprecated: use ServiceRestartV2 instead
func (adm *AdminClient) ServiceRestart(ctx context.Context) error {
	return adm.serviceCallAction(ctx, ServiceActionRestart)
}

// ServiceStop - stops the MinIO cluster
//
// Deprecated: use ServiceStopV2
func (adm *AdminClient) ServiceStop(ctx context.Context) error {
	return adm.serviceCallAction(ctx, ServiceActionStop)
}

// ServiceUnfreeze - un-freezes all incoming S3 API calls on MinIO cluster
//
// Deprecated: use ServiceUnfreezeV2
func (adm *AdminClient) ServiceUnfreeze(ctx context.Context) error {
	return adm.serviceCallAction(ctx, ServiceActionUnfreeze)
}

// serviceCallAction - call service restart/update/stop API.
func (adm *AdminClient) serviceCallAction(ctx context.Context, action ServiceAction) error {
	queryValues := url.Values{}
	queryValues.Set("action", string(action))

	// Request API to Restart server
	resp, err := adm.executeMethod(ctx,
		http.MethodPost, requestData{
			relPath:     adminAPIPrefix + "/service",
			queryValues: queryValues,
		},
	)
	defer closeResponse(resp)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return httpRespToErrorResponse(resp)
	}

	return nil
}
