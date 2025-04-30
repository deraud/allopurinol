package middleware

import (
	"healthcare-allopurinol/metrics"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func PrometheusMiddleware(c *gin.Context) {
	if c.Request.URL.Path == "/metrics" {
		c.Next()
		return
	}

	start := time.Now()
	path := c.FullPath()

	c.Next()

	duration := time.Since(start)
	status := strconv.Itoa(c.Writer.Status())

	metrics.HttpRequestsTotal.WithLabelValues(c.Request.Method, path, status).Inc()
	metrics.HttpRequestDuration.WithLabelValues(path).Observe(duration.Seconds())
	metrics.InFlightGauge.Inc()
}
