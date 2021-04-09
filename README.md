# Mininet Prometheus Exporter

Export metrics of virtual topologies to a prometheus instance. Must be using the NG-CDI Mininet fork that adds in a REST API (see [here](https://github.com/ng-cdi/mininet)).

Although this could perhaps be best done by creating some instrumentation within Mininet itself, having the API separate to the exporter tool allows for more use cases to be met with the least amount of code.

### TODO

 - Versioning System
 - GitHub Actions CI
 - YAML config for test sets
 - IPerf Collector
 - Custom Collector
