package metrics

import (
	metrics "github.com/rcrowley/go-metrics"
	"github.com/mihasya/go-metrics-librato"
	"github.com/adamdecaf/horizon/configs"
	"time"
)

func report_metrics_to_librato(registry metrics.Registry) {
	config := configs.NewConfig()

	owner := config.Get("LIBRATO_OWNER_EMAIL")
	token := config.Get("LIBRATO_API_TOKEN")
	hostname := config.Get("LIBRATO_INSTANCE_HOSTNAME")

	librato.Librato(registry,
		1,			// interval
		owner,			// account owner email address
		token,			// Librato API token
		hostname,		// source
		[]float64{0.95},	// percentiles to send
		time.Minute,		// time unit
	)
}
