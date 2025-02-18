//
//  MinIO Inc [madmin-go]
//  Copyright (c) 2014-2025 MinIO.
//  All rights reserved. No warranty, explicit or implicit, provided.
//

package madmin

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// SpeedTestStatServer - stats of a server
type SpeedTestStatServer struct {
	Endpoint         string `json:"endpoint"`
	ThroughputPerSec uint64 `json:"throughputPerSec"`
	ObjectsPerSec    uint64 `json:"objectsPerSec"`
	Err              string `json:"err"`
}

// SpeedTestStats - stats of all the servers
type SpeedTestStats struct {
	ThroughputPerSec uint64                `json:"throughputPerSec"`
	ObjectsPerSec    uint64                `json:"objectsPerSec"`
	Response         Timings               `json:"responseTime"`
	TTFB             Timings               `json:"ttfb,omitempty"`
	Servers          []SpeedTestStatServer `json:"servers"`
}

// SpeedTestResult - result of the speedtest() call
type SpeedTestResult struct {
	Version    string `json:"version"`
	Servers    int    `json:"servers"`
	Disks      int    `json:"disks"`
	Size       int    `json:"size"`
	Concurrent int    `json:"concurrent"`
	PUTStats   SpeedTestStats
	GETStats   SpeedTestStats
}

// SpeedtestOpts provide configurable options for speedtest
type SpeedtestOpts struct {
	Size         int           // Object size used in speed test
	Concurrency  int           // Concurrency used in speed test
	Duration     time.Duration // Total duration of the speed test
	Autotune     bool          // Enable autotuning
	StorageClass string        // Choose type of storage-class to be used while performing I/O
	Bucket       string        // Choose a custom bucket name while performing I/O
	NoClear      bool          // Avoid cleanup after running an object speed test
	EnableSha256 bool          // Enable calculating sha256 for uploads
}

// Speedtest - perform speedtest on the MinIO servers
func (adm *AdminClient) Speedtest(ctx context.Context, opts SpeedtestOpts) (chan SpeedTestResult, error) {
	if !opts.Autotune {
		if opts.Duration <= time.Second {
			return nil, errors.New("duration must be greater a second")
		}
		if opts.Size <= 0 {
			return nil, errors.New("size must be greater than 0 bytes")
		}
		if opts.Concurrency <= 0 {
			return nil, errors.New("concurrency must be greater than 0")
		}
	}

	queryVals := make(url.Values)
	if opts.Size > 0 {
		queryVals.Set("size", strconv.Itoa(opts.Size))
	}
	if opts.Duration > 0 {
		queryVals.Set("duration", opts.Duration.String())
	}
	if opts.Concurrency > 0 {
		queryVals.Set("concurrent", strconv.Itoa(opts.Concurrency))
	}
	if opts.Bucket != "" {
		queryVals.Set("bucket", opts.Bucket)
	}
	if opts.Autotune {
		queryVals.Set("autotune", "true")
	}
	if opts.NoClear {
		queryVals.Set("noclear", "true")
	}
	if opts.EnableSha256 {
		queryVals.Set("enableSha256", "true")
	}
	resp, err := adm.executeMethod(ctx,
		http.MethodPost, requestData{
			relPath:     adminAPIPrefix + "/speedtest",
			queryValues: queryVals,
		})
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, httpRespToErrorResponse(resp)
	}
	ch := make(chan SpeedTestResult)
	go func() {
		defer closeResponse(resp)
		defer close(ch)
		dec := json.NewDecoder(resp.Body)
		for {
			var result SpeedTestResult
			if err := dec.Decode(&result); err != nil {
				return
			}
			select {
			case ch <- result:
			case <-ctx.Done():
				return
			}
		}
	}()
	return ch, nil
}
