package metrics

import (
        "log"
	"os"
	"sync"
	"time"
	metrics "github.com/rcrowley/go-metrics"
	"github.com/adamdecaf/horizon/utils"
)

var registry metrics.Registry = metrics.DefaultRegistry

var mu sync.Mutex
var meters map[string]metrics.Meter = make(map[string]metrics.Meter)
var timers map[string]metrics.Timer = make(map[string]metrics.Timer)

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

func Timer(name string) metrics.Timer {
	exists := timers[name]

	if exists == nil {
		mu.Lock()
		timers[name] = metrics.NewTimer()
		mu.Unlock()
	}

	t := timers[name]
	metrics.Register(name, t)

	return t
}

func InitializeStdoutReporter() {
	config := utils.NewConfig()

	if run := config.Get("STDOUT_REPORTING_ENABLED"); run == "yes" {
                log.Println("starting stdout metrics reporting")
                go metrics.Log(metrics.DefaultRegistry, 1 * time.Minute, log.New(os.Stderr, "metrics: ", log.Lmicroseconds))
	}

	if run := config.Get("LIBRATO_REPORTING_ENABLED"); run == "yes" {
		log.Println("starting librato metrics reporting")
		go report_metrics_to_librato(metrics.DefaultRegistry)
	}
}
