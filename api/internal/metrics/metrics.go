package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	APIRequests = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "api_requests_total",
		Help: "Total API requests ingested",
	})
	APIEnqueued = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "api_requests_enqueued_total",
		Help: "Total API requests successfully enqueued",
	})
)

func init() {
	prometheus.MustRegister(APIRequests, APIEnqueued)
}

func Handler() http.Handler {
	return promhttp.Handler()
}
