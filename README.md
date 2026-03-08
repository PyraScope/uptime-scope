# Uptime Scope
<p align="center">
<img src="https://img.shields.io/github/go-mod/go-version/Pyrascope/uptime-scope" alt="Go Version">
<img src="https://img.shields.io/badge/Architecture-amd64%20%7C%20arm64-orange" alt="Supported Archicetures">
<img src="https://img.shields.io/github/actions/workflow/status/Pyrascope/uptime-scope/ci.yml" alt="CI Build">
</p>

A Prometheus Exporter for Uptime Monitoring

## Running agent:
```
go mod tidy
go run main.go
```

## Running with Docker Compose for Development:
```
docker compose -f docker-compose-dev.yml up --build
```

## Running with Docker Compose for Production:
```
docker compose up -d
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
