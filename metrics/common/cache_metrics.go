package common

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

const (
	DefaultCacheRequestMetricName = "apm_cache_request_duration_seconds"
)

var (
	DefaultCacheRequestMetric = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:                            DefaultCacheRequestMetricName,
		Buckets:                         []float64{0.00025, 0.0005, 0.001, 0.002, 0.005, 0.01, 0.25, 0.5, 1, 2.5, 5, 10},
		NativeHistogramBucketFactor:     1.4,
		NativeHistogramZeroThreshold:    0.0001,
		NativeHistogramMinResetDuration: 5 * time.Minute,
		NativeHistogramMaxZeroThreshold: 0.01,
		NativeHistogramMaxBucketNumber:  20,
	}, []string{"sdk", "cache_type", "cache_host", "command", "response_status"})
)