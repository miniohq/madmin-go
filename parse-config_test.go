//
//  MinIO Inc [madmin-go]
//  Copyright (c) 2014-2025 MinIO.
//  All rights reserved. No warranty, explicit or implicit, provided.
//

package madmin

import (
	"reflect"
	"testing"
)

func TestParseServerConfigOutput(t *testing.T) {
	tests := []struct {
		Name        string
		Config      string
		Expected    []SubsysConfig
		ExpectedErr error
	}{
		{
			Name:   "single target config data only",
			Config: "subnet license= api_key= proxy=",
			Expected: []SubsysConfig{
				{
					SubSystem: SubnetSubSys,
					Target:    "",
					KV: []ConfigKV{
						{
							Key:         "license",
							Value:       "",
							EnvOverride: nil,
						},
						{
							Key:         "api_key",
							Value:       "",
							EnvOverride: nil,
						},
						{
							Key:         "proxy",
							Value:       "",
							EnvOverride: nil,
						},
					},
					kvIndexMap: map[string]int{
						"license": 0,
						"api_key": 1,
						"proxy":   2,
					},
				},
			},
		},
		{
			Name: "single target config + env",
			Config: `# MINIO_SUBNET_API_KEY=xxx
# MINIO_SUBNET_LICENSE=2
subnet license=1 api_key= proxy=`,
			Expected: []SubsysConfig{
				{
					SubSystem: SubnetSubSys,
					Target:    "",
					KV: []ConfigKV{
						{
							Key:   "api_key",
							Value: "",
							EnvOverride: &EnvOverride{
								Name:  "MINIO_SUBNET_API_KEY",
								Value: "xxx",
							},
						},
						{
							Key:   "license",
							Value: "1",
							EnvOverride: &EnvOverride{
								Name:  "MINIO_SUBNET_LICENSE",
								Value: "2",
							},
						},
						{
							Key:         "proxy",
							Value:       "",
							EnvOverride: nil,
						},
					},
					kvIndexMap: map[string]int{
						"license": 1,
						"api_key": 0,
						"proxy":   2,
					},
				},
			},
		},
		{
			Name: "multiple targets no env",
			Config: `logger_webhook enable=off endpoint= auth_token= client_cert= client_key= queue_size=100000
logger_webhook:1 endpoint=http://localhost:8080/ auth_token= client_cert= client_key= queue_size=100000
`,
			Expected: []SubsysConfig{
				{
					SubSystem: LoggerWebhookSubSys,
					Target:    "",
					KV: []ConfigKV{
						{
							Key:   "enable",
							Value: "off",
						},
						{
							Key:   "endpoint",
							Value: "",
						},
						{
							Key:   "auth_token",
							Value: "",
						},
						{
							Key:   "client_cert",
							Value: "",
						},
						{
							Key:   "client_key",
							Value: "",
						},
						{
							Key:   "queue_size",
							Value: "100000",
						},
					},
					kvIndexMap: map[string]int{
						"enable":      0,
						"endpoint":    1,
						"auth_token":  2,
						"client_cert": 3,
						"client_key":  4,
						"queue_size":  5,
					},
				},
				{
					SubSystem: LoggerWebhookSubSys,
					Target:    "1",
					KV: []ConfigKV{
						{
							Key:   "endpoint",
							Value: "http://localhost:8080/",
						},
						{
							Key:   "auth_token",
							Value: "",
						},
						{
							Key:   "client_cert",
							Value: "",
						},
						{
							Key:   "client_key",
							Value: "",
						},
						{
							Key:   "queue_size",
							Value: "100000",
						},
					},
					kvIndexMap: map[string]int{
						"endpoint":    0,
						"auth_token":  1,
						"client_cert": 2,
						"client_key":  3,
						"queue_size":  4,
					},
				},
			},
		},
	}

	for i, test := range tests {
		r, err := ParseServerConfigOutput(test.Config)
		if err != nil {
			if err.Error() != test.ExpectedErr.Error() {
				t.Errorf("Test %d (%s) got unexpected error: %v", i, test.Name, err)
			}
			// got an expected error.
			continue
		}
		if !reflect.DeepEqual(test.Expected, r) {
			t.Errorf("Test %d (%s) expected:\n%#v\nbut got:\n%#v\n", i, test.Name, test.Expected, r)
		}
	}
}
