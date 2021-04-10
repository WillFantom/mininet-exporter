# Mininet Prometheus Exporter

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
