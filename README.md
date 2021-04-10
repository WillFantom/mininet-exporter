# Mininet Prometheus Exporter

[![Publish Rolling Release](https://github.com/WillFantom/mininet-exporter/actions/workflows/rolling.yml/badge.svg?branch=main)](https://github.com/WillFantom/mininet-exporter/actions/workflows/rolling.yml) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/willfantom/mininet-exporter) ![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/willfantom/mininet-exporter?sort=semver) ![GitHub](https://img.shields.io/github/license/willfantom/mininet-exporter)

Export metrics of virtual topologies to a prometheus instance. Must be using the NG-CDI Mininet fork that adds in a REST API (see [here](https://github.com/ng-cdi/mininet)).

Although this could perhaps be best done by creating some instrumentation within Mininet itself, having the API separate to the exporter tool allows for more use cases to be met with the least amount of code.

## Build

To build the docker image for this, run the following command:
```bash
docker build --rm -f Dockerfile --build-arg EXPORTER_VERSION=rolling -t mn-exporter:local .
```

## Configuration

An example config can is given as [`tests.yaml`](./tests.yaml)...

### TODO

 - IPerf Collector
 - Custom Collector
