//
// MinIO Inc [madmin-go]
// Copyright (c) 2014-2024 MinIO.
// All rights reserved. No warranty, explicit or implicit, provided.
//

package madmin

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/dustin/go-humanize"
	"github.com/prometheus/common/expfmt"
	"github.com/prometheus/prom2json"
)

// MetricsRespBodyLimit sets the top level limit to the size of the
// metrics results supported by this library.
var (
	MetricsRespBodyLimit = int64(humanize.GiByte)
)

// NodeMetrics - returns Node Metrics in Prometheus format
//
//	The client needs to be configured with the endpoint of the desired node
func (client *MetricsClient) NodeMetrics(ctx context.Context) ([]*prom2json.Family, error) {
	return client.GetMetrics(ctx, "node")
}

// ClusterMetrics - returns Cluster Metrics in Prometheus format
func (client *MetricsClient) ClusterMetrics(ctx context.Context) ([]*prom2json.Family, error) {
	return client.GetMetrics(ctx, "cluster")
}

// BucketMetrics - returns Bucket Metrics in Prometheus format
func (client *MetricsClient) BucketMetrics(ctx context.Context) ([]*prom2json.Family, error) {
	return client.GetMetrics(ctx, "bucket")
}

// ResourceMetrics - returns Resource Metrics in Prometheus format
func (client *MetricsClient) ResourceMetrics(ctx context.Context) ([]*prom2json.Family, error) {
	return client.GetMetrics(ctx, "resource")
}

// GetMetrics - returns Metrics of given subsystem in Prometheus format
func (client *MetricsClient) GetMetrics(ctx context.Context, subSystem string) ([]*prom2json.Family, error) {
	reqData := metricsRequestData{
		relativePath: "/v2/metrics/" + subSystem,
	}

	// Execute GET on /minio/v2/metrics/<subSys>
	resp, err := client.executeGetRequest(ctx, reqData)
	if err != nil {
		return nil, err
	}
	defer closeResponse(resp)

	if resp.StatusCode != http.StatusOK {
		return nil, httpRespToErrorResponse(resp)
	}

	return ParsePrometheusResults(io.LimitReader(resp.Body, MetricsRespBodyLimit))
}

func ParsePrometheusResults(reader io.Reader) (results []*prom2json.Family, err error) {
	// We could do further content-type checks here, but the
	// fallback for now will anyway be the text format
	// version 0.0.4, so just go for it and see if it works.
	var parser expfmt.TextParser
	metricFamilies, err := parser.TextToMetricFamilies(reader)
	if err != nil {
		return nil, fmt.Errorf("reading text format failed: %v", err)
	}
	results = make([]*prom2json.Family, 0, len(metricFamilies))
	for _, mf := range metricFamilies {
		results = append(results, prom2json.NewFamily(mf))
	}
	return results, nil
}
