package metrics

import (
        "log"
	"os"
	"sync"
	"time"
	metrics "github.com/rcrowley/go-metrics"
)

var registry metrics.Registry = metrics.DefaultRegistry

var mu sync.Mutex
var meters map[string]metrics.Meter = make(map[string]metrics.Meter)

func Meter(name string) metrics.Meter {
	exists := meters[name]

	if exists == nil {
		mu.Lock()
		meters[name] = metrics.NewMeter()
		mu.Unlock()
	}

        m := meters[name]
        metrics.Register(name, m)

	return m
}

func report_metrics_to_stdout() {
        log.Println("reporting metrics to stdout")
	out := os.Stdout
	metrics.WriteOnce(registry, out)
}

func InitializeStdoutReporter() {
	if run := os.Getenv("STDOUT_REPORTING_ENABLED"); run == "yes" {
                log.Println("starting stdout metrics reporting")
                go metrics.Log(metrics.DefaultRegistry, 1 * time.Minute, log.New(os.Stderr, "metrics: ", log.Lmicroseconds))
	}
}
