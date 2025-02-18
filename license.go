//
//  MinIO Inc [madmin-go]
//  Copyright (c) 2014-2025 MinIO.
//  All rights reserved. No warranty, explicit or implicit, provided.
//

package madmin

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// LicenseInfo is a structure containing MinIO license information.
type LicenseInfo struct {
	ID           string    `json:"ID"`           // The license ID
	Organization string    `json:"Organization"` // Name of the organization using the license
	Plan         string    `json:"Plan"`         // License plan. E.g. "ENTERPRISE-PLUS"
	IssuedAt     time.Time `json:"IssuedAt"`     // Point in time when the license was issued
	ExpiresAt    time.Time `json:"ExpiresAt"`    // Point in time when the license expires
	Trial        bool      `json:"Trial"`        // Whether the license is on trial
	APIKey       string    `json:"APIKey"`       // Subnet account API Key
}

// GetLicenseInfo - returns the license info
func (adm *AdminClient) GetLicenseInfo(ctx context.Context) (*LicenseInfo, error) {
	// Execute GET on /minio/admin/v3/licenseinfo to get license info.
	resp, err := adm.executeMethod(ctx,
		http.MethodGet,
		requestData{
			relPath: adminAPIPrefix + "/license-info",
		})
	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, httpRespToErrorResponse(resp)
	}

	l := LicenseInfo{}
	err = json.NewDecoder(resp.Body).Decode(&l)
	if err != nil {
		return nil, err
	}
	return &l, nil
}
