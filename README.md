# Uptime Scope
Uptime Monitor

## Running agent:
```
go mod tidy
go run main.go
```

## Metrics:
```
http://localhost:8181/metrics
```

# Prometheus.yml example:
```
scrape_configs:
  - job_name: 'uptime_scope'
    scrape_interval: 30s
    static_configs:
      - targets:
          - 'uptimescope-eu-west:8181'
          - 'uptimescope-us-east:8181'
          - 'uptimescope-asia:8181'
```

## Contribution

Please check our [Contributing Guide](./docs/CONTRIBUTING.md) on how you can contribute.
