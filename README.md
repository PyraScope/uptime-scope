# Uptime Scope
A Prometheus Exporter for Uptime Monitoring

## Running agent:
```
go mod tidy
go run main.go
```

## Running with Docker Compose:
```
docker compose up --build -d
```

Stop:
```
docker compose down
```

## Metrics:
```
http://localhost:8181/metrics
```

## Prometheus.yml example:
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

## Ports:
| Default Port | Name | Listing |
|---|---|---|
| 8181 | UptimeScope | [Prometheus](https://github.com/prometheus/prometheus/wiki/Default-port-allocations) |

## Contribution

Please check our [Contributing Guide](./docs/CONTRIBUTING.md) on how you can contribute.
