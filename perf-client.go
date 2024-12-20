//
// MinIO Inc [madmin-go]
// Copyright (c) 2014-2024 MinIO.
// All rights reserved. No warranty, explicit or implicit, provided.
//

package madmin

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/dustin/go-humanize"
)

// ClientPerfExtraTime - time for get lock or other
type ClientPerfExtraTime struct {
	TimeSpent int64 `json:"dur,omitempty"`
}

// ClientPerfResult - stats from  client to server
type ClientPerfResult struct {
	Endpoint  string `json:"endpoint,omitempty"`
	Error     string `json:"error,omitempty"`
	BytesSend uint64
	TimeSpent int64
}

// clientPerfReader - wrap the reader
type clientPerfReader struct {
	count     uint64
	startTime time.Time
	endTime   time.Time
	buf       []byte
}

// Start - reader start
func (c *clientPerfReader) Start() {
	buf := make([]byte, 128*humanize.KiByte)
	rand.Read(buf)
	c.buf = buf
	c.startTime = time.Now()
}

// End - reader end
func (c *clientPerfReader) End() {
	c.endTime = time.Now()
}

// Read - reader send data
func (c *clientPerfReader) Read(p []byte) (n int, err error) {
	n = copy(p, c.buf)
	c.count += uint64(n)
	return n, nil
}

var _ io.Reader = &clientPerfReader{}

const (
	// MaxClientPerfTimeout for max time out for client perf
	MaxClientPerfTimeout = time.Second * 30
	// MinClientPerfTimeout for min time out for client perf
	MinClientPerfTimeout = time.Second * 5
)

// ClientPerf - perform net from client to MinIO servers
func (adm *AdminClient) ClientPerf(ctx context.Context, dur time.Duration) (result ClientPerfResult, err error) {
	if dur > MaxClientPerfTimeout {
		dur = MaxClientPerfTimeout
	}
	if dur < MinClientPerfTimeout {
		dur = MinClientPerfTimeout
	}
	ctx, cancel := context.WithTimeout(ctx, dur)
	defer cancel()
	queryVals := make(url.Values)
	reader := &clientPerfReader{}
	reader.Start()
	_, err = adm.executeMethod(ctx, http.MethodPost, requestData{
		queryValues:   queryVals,
		relPath:       adminAPIPrefix + "/speedtest/client/devnull",
		contentReader: reader,
	})
	reader.End()
	if errors.Is(err, context.DeadlineExceeded) && ctx.Err() != nil {
		err = nil
	}

	resp, err := adm.executeMethod(context.Background(), http.MethodPost, requestData{
		queryValues: queryVals,
		relPath:     adminAPIPrefix + "/speedtest/client/devnull/extratime",
	})
	if err != nil {
		return ClientPerfResult{}, err
	}
	var extraTime ClientPerfExtraTime
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&extraTime)
	if err != nil {
		return ClientPerfResult{}, err
	}
	durSpend := reader.endTime.Sub(reader.startTime).Nanoseconds()
	if extraTime.TimeSpent > 0 {
		durSpend = durSpend - extraTime.TimeSpent
	}
	if durSpend <= 0 {
		return ClientPerfResult{}, fmt.Errorf("unexpected spent time duration, mostly NTP errors on the server")
	}
	return ClientPerfResult{
		BytesSend: reader.count,
		TimeSpent: durSpend,
		Error:     "",
		Endpoint:  adm.endpointURL.String(),
	}, err
}
