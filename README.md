# Uptime Scope
Uptime Monitor

## Running agent:
```
go mod tidy
go run main.go
```

## Metrics:
```
http://localhost:9100/metrics
```

# Prometheus.yml example:
```
scrape_configs:
  - job_name: 'uptime_agents'
    scrape_interval: 30s
    static_configs:
      - targets:
          - 'agent-eu-west:9100'
          - 'agent-us-east:9100'
          - 'agent-asia:9100'
```
