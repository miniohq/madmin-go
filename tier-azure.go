//
//  MinIO Inc [madmin-go]
//  Copyright (c) 2014-2025 MinIO.
//  All rights reserved. No warranty, explicit or implicit, provided.
//

package madmin

import "errors"

//go:generate msgp -file $GOFILE

// ServicePrincipalAuth holds fields for a successful SP authentication with Azure
type ServicePrincipalAuth struct {
	TenantID     string `json:",omitempty"`
	ClientID     string `json:",omitempty"`
	ClientSecret string `json:",omitempty"`
}

// TierAzure represents the remote tier configuration for Azure Blob Storage.
type TierAzure struct {
	Endpoint     string `json:",omitempty"`
	AccountName  string `json:",omitempty"`
	AccountKey   string `json:",omitempty"`
	Bucket       string `json:",omitempty"`
	Prefix       string `json:",omitempty"`
	Region       string `json:",omitempty"`
	StorageClass string `json:",omitempty"`

	SPAuth ServicePrincipalAuth `json:",omitempty"`
}

// IsSPEnabled returns true if all SP related fields are provided
func (ti TierAzure) IsSPEnabled() bool {
	return ti.SPAuth.TenantID != "" && ti.SPAuth.ClientID != "" && ti.SPAuth.ClientSecret != ""
}

// AzureOptions supports NewTierAzure to take variadic options
type AzureOptions func(*TierAzure) error

// AzureServicePrincipal helper to supply optional service principal credentials
func AzureServicePrincipal(tenantID, clientID, clientSecret string) func(az *TierAzure) error {
	return func(az *TierAzure) error {
		if tenantID == "" {
			return errors.New("empty tenant ID unsupported")
		}
		if clientID == "" {
			return errors.New("empty client ID unsupported")
		}
		if clientSecret == "" {
			return errors.New("empty client secret unsupported")
		}
		az.SPAuth.TenantID = tenantID
		az.SPAuth.ClientID = clientID
		az.SPAuth.ClientSecret = clientSecret
		return nil
	}
}

// AzurePrefix helper to supply optional object prefix to NewTierAzure
func AzurePrefix(prefix string) func(az *TierAzure) error {
	return func(az *TierAzure) error {
		az.Prefix = prefix
		return nil
	}
}

// AzureEndpoint helper to supply optional endpoint to NewTierAzure
func AzureEndpoint(endpoint string) func(az *TierAzure) error {
	return func(az *TierAzure) error {
		az.Endpoint = endpoint
		return nil
	}
}

// AzureRegion helper to supply optional region to NewTierAzure
func AzureRegion(region string) func(az *TierAzure) error {
	return func(az *TierAzure) error {
		az.Region = region
		return nil
	}
}

// AzureStorageClass helper to supply optional storage class to NewTierAzure
func AzureStorageClass(sc string) func(az *TierAzure) error {
	return func(az *TierAzure) error {
		az.StorageClass = sc
		return nil
	}
}

// NewTierAzure returns a TierConfig of Azure type. Returns error if the given
// parameters are invalid like name is empty etc.
func NewTierAzure(name, accountName, accountKey, bucket string, options ...AzureOptions) (*TierConfig, error) {
	if name == "" {
		return nil, ErrTierNameEmpty
	}

	az := &TierAzure{
		AccountName: accountName,
		AccountKey:  accountKey,
		Bucket:      bucket,
		// Defaults
		Endpoint:     "",
		Prefix:       "",
		Region:       "",
		StorageClass: "",
	}

	for _, option := range options {
		err := option(az)
		if err != nil {
			return nil, err
		}
	}

	return &TierConfig{
		Version: TierConfigVer,
		Type:    Azure,
		Name:    name,
		Azure:   az,
	}, nil
}
