//
//  MinIO Inc [madmin-go]
//  Copyright (c) 2014-2025 MinIO.
//  All rights reserved. No warranty, explicit or implicit, provided.
//

package madmin

//go:generate msgp -file $GOFILE

// TierMinIO represents the remote tier configuration for MinIO object storage backend.
type TierMinIO struct {
	Endpoint  string `json:",omitempty"`
	AccessKey string `json:",omitempty"`
	SecretKey string `json:",omitempty"`
	Bucket    string `json:",omitempty"`
	Prefix    string `json:",omitempty"`
	Region    string `json:",omitempty"`
}

// MinIOOptions supports NewTierMinIO to take variadic options
type MinIOOptions func(*TierMinIO) error

// MinIORegion helper to supply optional region to NewTierMinIO
func MinIORegion(region string) func(m *TierMinIO) error {
	return func(m *TierMinIO) error {
		m.Region = region
		return nil
	}
}

// MinIOPrefix helper to supply optional object prefix to NewTierMinIO
func MinIOPrefix(prefix string) func(m *TierMinIO) error {
	return func(m *TierMinIO) error {
		m.Prefix = prefix
		return nil
	}
}

func NewTierMinIO(name, endpoint, accessKey, secretKey, bucket string, options ...MinIOOptions) (*TierConfig, error) {
	if name == "" {
		return nil, ErrTierNameEmpty
	}
	m := &TierMinIO{
		AccessKey: accessKey,
		SecretKey: secretKey,
		Bucket:    bucket,
		Endpoint:  endpoint,
	}

	for _, option := range options {
		err := option(m)
		if err != nil {
			return nil, err
		}
	}

	return &TierConfig{
		Version: TierConfigVer,
		Type:    MinIO,
		Name:    name,
		MinIO:   m,
	}, nil
}
