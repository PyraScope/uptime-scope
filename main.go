package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// --- Configuration ---
var checks []string
var region string

type Config struct {
	Region string   `yaml:"region"`
	Checks []string `yaml:"checks"`
}

// --- Metrics ---
var (
	successGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "uptime_check_success",
			Help: "Whether the last check succeeded (1) or failed (0)",
		},
		[]string{"target", "region"},
	)

	durationGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "uptime_check_duration_seconds",
			Help: "Duration of the last uptime check in seconds",
		},
		[]string{"target", "region"},
	)

	statusCodeGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "uptime_check_status_code",
			Help: "HTTP status code from the last check",
		},
		[]string{"target", "region"},
	)
)

func init() {
	prometheus.MustRegister(successGauge)
	prometheus.MustRegister(durationGauge)
	prometheus.MustRegister(statusCodeGauge)
}

func loadConfig(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return fmt.Errorf("failed to parse YAML: %w", err)
	}

	checks = cfg.Checks
	region = cfg.Region
	return nil
}

func runChecks() {
	client := &http.Client{Timeout: 10 * time.Second}

	for _, url := range checks {
		start := time.Now()
		resp, err := client.Get(url)
		duration := time.Since(start).Seconds()

		durationGauge.WithLabelValues(url, region).Set(duration)

		if err != nil {
			successGauge.WithLabelValues(url, region).Set(0)
			statusCodeGauge.WithLabelValues(url, region).Set(0)
			log.Printf("[FAIL] %s: %v", url, err)
			continue
		}

		statusCodeGauge.WithLabelValues(url, region).Set(float64(resp.StatusCode))
		resp.Body.Close()

		if resp.StatusCode >= 200 && resp.StatusCode < 400 {
			successGauge.WithLabelValues(url, region).Set(1)
			log.Printf("[OK] %s (%d) in %.2fs", url, resp.StatusCode, duration)
		} else {
			successGauge.WithLabelValues(url, region).Set(0)
			log.Printf("[FAIL] %s (%d) in %.2fs", url, resp.StatusCode, duration)
		}
	}
}

func main() {
	if err := loadConfig("checks.yaml"); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Periodic checker
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		for range ticker.C {
			runChecks()
		}
	}()

	// Metrics endpoint
	http.Handle("/metrics", promhttp.Handler())
	log.Println("Agent running on :8181")
	log.Fatal(http.ListenAndServe(":8181", nil))
}
