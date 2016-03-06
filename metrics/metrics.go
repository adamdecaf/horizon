package metrics

import (
	"os"
	"sync"
	"time"
	metrics "github.com/rcrowley/go-metrics"
)

var registry metrics.Registry = metrics.NewRegistry()

var mu sync.Mutex
var meters map[string]metrics.Meter

func Meter(name string) metrics.Meter {
	exists := meters[name]

	if exists == nil {
		mu.Lock()
		meters[name] = metrics.NewMeter()
		mu.Unlock()
	}

	return meters[name]
}

func report_metrics_to_stdout() {
	out := os.Stdout
	metrics.WriteOnce(registry, out)
}

func InitializeStdoutReporter() {
	if run := os.Getenv("STDOUT_REPORTING_ENABLED"); run == "yes" {
		t := time.Tick(1 * time.Minute)
		for _ = range t {
			report_metrics_to_stdout()
		}
	}
}
