# Migration
This is a migration doc for the Observability for Kubernetes Operator from the manual and Helm Collector and proxy installation.
If you want to test the new Operator in parallel with your existing manual or Helm installation, use the [wavefront-allow-legacy-intall.yaml](../../deploy/scenarios/wavefront-allow-legacy-install.yaml) template.

## Migrate from Helm Installation

### 1. Install the Operator

Follow [the Operator installation instructions](../../README.md#Deploy-the-Kubernetes-Metrics-Collector-and-Wavefront-Proxy-with-the-Observability-for-Kubernetes-Operator).

In your `wavefront.yaml`,
 * set `spec.allowLegacyInstall` to `true`
 * set `spec.clusterName` to `<YOUR HELM CLUSTER NAME>-operator` 

### 2. Modify Your `wavefront.yaml`

Modify your `wavefront.yaml` to match your Helm installation based on the information below.
The following table lists the mapping of configurable parameters of the Wavefront Helm chart to Observability for Kubernetes Operator Custom Resource.
See [Custom Resource Scenarios](../../deploy/scenarios) for examples or refer to [wavefront.com_wavefronts.yaml](../../deploy/crd/wavefront.com_wavefronts.yaml) for information on all Custom Resource fields.

| Helm Collector Parameter | Observability for Kubernetes Operator Custom Resource `spec`. | Description |
|---|---|---|
| `clusterName` | `clusterName` | ClusterName is a unique name for the Kubernetes cluster to be identified via a metric tag in Operations for Applications. |
| `wavefront.url` | `wavefrontUrl` | The URL of your product cluster. Ex: https://<your_cluster>.wavefront.com. |
| `wavefront.token` | `wavefrontTokenSecret` | WavefrontTokenSecret is the name of the secret that contains a Wavefront API Token. |
| `collector.enabled` | `dataCollection.metrics.enable` | Metrics holds the configuration for node and cluster collectors. |
| `collector.interval` | `dataCollection.metrics.defaultCollectionInterval` | Default metrics collection interval. Defaults to 60s. |
| `collector.useProxy` | `NA` | The earlier Collector config parameter was meant to be used to support direct ingestion, which the Operator doesn't support. |
| `collector.proxyAddress` | `dataExport.externalWavefrontProxy.Url` | Url is the proxy URL that the collector sends metrics to. |
| `collector.tags` | `dataCollection.metrics.tags` | Map of tags (key/value) to add to all metrics collected. |
| `collector.filters.metricDenyList` | `dataCollection.metrics.filters.denyList` | List of metric patterns to deny. |
| `collector.filters.metricAllowList` | `dataCollection.metrics.filters.allowList` | List of metric patterns to allow. |
| `collector.discovery.enabled` | `dataCollection.metrics.enableDiscovery` | Rules based and Prometheus endpoints auto-discovery. Defaults to true. |
| `collector.resources` | `dataCollection.metrics.nodeCollector.resources` `dataCollection.metrics.clusterCollector.resources` | Compute resources required by the node and cluster collector containers. |
| `proxy.enabled` | `dataExport.wavefrontProxy.enable` | Enable is whether to enable the Wavefront proxy. Defaults to true. Disable to use `dataExport.externalWavefrontProxy.Url`. |
| `proxy.port` | `dataExport.wavefrontProxy.metricPort` | MetricPort is the port for sending Operations for Applications data format metrics. Defaults to 2878. |
| `proxy.httpProxyHost` | `dataExport.wavefrontProxy.httpProxy.secret` | Name of the secret containing the HttpProxy configuration. |
| `proxy.httpProxyPort` | `dataExport.wavefrontProxy.httpProxy.secret` | Name of the secret containing the HttpProxy configuration. |
| `proxy.useHttpProxyCAcert` | `dataExport.wavefrontProxy.httpProxy.secret` | Name of the secret containing the HttpProxy configuration. |
| `proxy.httpProxyUser` | `dataExport.wavefrontProxy.httpProxy.secret` | Name of the secret containing the HttpProxy configuration. |
| `proxy.httpProxyPassword` | `dataExport.wavefrontProxy.httpProxy.secret` | Name of the secret containing the HttpProxy configuration. |
| `proxy.tracePort` | `dataExport.wavefrontProxy.tracing.wavefront.port` | Port for sending distributed Operations for Applications format tracing data (usually 30000). |
| `proxy.jaegerPort` | `dataExport.wavefrontProxy.tracing.jaeger.port` | Port for Jaeger format tracing data (usually 30001). |
| `proxy. traceJaegerHttpListenerPort` | `dataExport.wavefrontProxy.tracing.jaeger. httpPort` | HttpPort for Jaeger Thrift format data (usually 30080). |
| `proxy. traceJaegerGrpcListenerPort` | `dataExport.wavefrontProxy.tracing.jaeger. grpcPort` | GrpcPort for Jaeger gRPC format data (usually 14250). |
| `proxy.zipkinPort` | `dataExport.wavefrontProxy.tracing.zipkin.port` | Port for Zipkin format tracing data (usually 9411). |
| `proxy.traceSamplingRate` | `dataExport.wavefrontProxy.tracing.wavefront. samplingRate` | SamplingRate Distributed tracing data sampling rate (0 to 1). |
| `proxy.traceSamplingDuration` | `dataExport.wavefrontProxy.tracing.wavefront. samplingDuration` | SamplingDuration When set to greater than 0, spans that exceed this duration will force trace to be sampled (ms). |
| `proxy. traceJaegerApplicationName` | `dataExport.wavefrontProxy.tracing.jaeger. applicationName` | Custom application name for traces received on Jaeger's HTTP or gRPC port. |
| `proxy. traceZipkinApplicationName` | `dataExport.wavefrontProxy.tracing.zipkin. applicationName` | Custom application name for traces received on Zipkin's port. |
| `proxy.histogramPort` | `dataExport.wavefrontProxy.histogram.port` | Port for Operations for Applications histogram distributions (usually 40000) |
| `proxy.histogramMinutePort` | `dataExport.wavefrontProxy.histogram.minutePort` | Port to accumulate 1-minute based histograms in the Operations for Applications data format (usually 40001). |
| `proxy.histogramHourPort` | `dataExport.wavefrontProxy.histogram.hourPort` | Port to accumulate 1-hour based histograms in the Operations for Applications data format (usually 40002). |
| `proxy.histogramDayPort` | `dataExport.wavefrontProxy.histogram.dayPort` | Port to accumulate 1-day based histograms in the Operations for Applications data format (usually 40002). |
| `proxy.deltaCounterPort` | `dataExport.wavefrontProxy.deltaCounterPort` | Port to send delta counters in the Operations for Applications data format (usually 50000). |
| `proxy.args` | `dataExport.wavefrontProxy.args` | Additional Wavefront proxy properties can be passed as command line arguments in the `--<property_name> <value>` format. Multiple properties can be specified. |
| `proxy.preprocessor.rules.yaml` | `dataExport.wavefrontProxy.preprocessor` | Name of the configmap containing a rules.yaml key with proxy preprocessing rules. |

If you have a Collector configuration with parameters not covered above, please open an issue in this repository.

### 3. Re-apply Your `wavefront.yaml` File

```shell
kubectl apply -f <path_to_your_wavefront.yaml>
```

### 4. Verify The Operator Status

Check the status of your installation:

```shell
kubectl get wavefront -n observability-system
```

You should eventually see a table that looks something like the following:

```
NAME        STATUS    PROXY           CLUSTER-COLLECTOR   NODE-COLLECTOR   LOGGING        AGE    MESSAGE
wavefront   Healthy   Running (1/1)   Running (1/1)       Running (3/3)                   2m4s   All components are healthy
```

If `STATUS` is `Unhealthy`, check [troubleshooting](../troubleshooting.md).

### 5. Verify That You Are Receiving Metrics from the Operator

 * Go to the Kubernetes Summary Dashboard, set the cluster to `<YOUR HELM CLUSTER NAME>-operator`, and verify that you are receiving metrics. It may take some time for metrics to populate this dashboard.
 * Verify that other important metrics are present with a `cluster="<YOUR HELM CLUSTER NAME>-operator"` tag.

### 6. Uninstall Helm

```shell
helm uninstall wavefront --namespace wavefront
```

### 7. Clean Up the `wavefront.yaml` File

In your `wavefront.yaml`,
 * (Optional) Update `spec.clusterName` to `<YOUR HELM CLUSTER NAME>`
 * Update `spec.allowLegacyInstall` to `false`

```shell
kubectl apply -f <path_to_your_wavefront.yaml>
```

## Migrate from Manual Installation 

### Wavefront Proxy Configuration

#### References:
* See [Custom Resource Scenarios](../../deploy/scenarios) for proxy configuration examples.
* Copy and save your existing Collector configMaps and any other configurations.
* Uninstall your currently deployed Collector and Wavefront proxy.
* Make sure the Observability for Kubernetes Operator is already installed.
* Create a Kubernetes secret with your Wavefront API Token `kubectl create -n observability-system secret generic wavefront-secret --from-literal token=YOUR_WAVEFRONT_TOKEN`

Most of the proxy configurations could be set using environment variables for proxy container.
Here are the different proxy environment variables and how they map to the Operator config.

| Proxy Environment Variables | Observability for Kubernetes Operator Custom Resource `spec` |
|---|---|
| `WAVEFRONT_URL` | `wavefrontUrl` Ex: https://<your_cluster>.wavefront.com |
| `WAVEFRONT_TOKEN` | `WAVEFRONT_TOKEN` is now stored in a Kubernetes secret; see **References** above. |
| `WAVEFRONT_PROXY_ARGS` | `dataExport.wavefrontProxy.*` Refer to the below table for details. |

Below are the proxy arguments that are specified in `WAVEFRONT_PROXY_ARGS`, which are currently supported natively in the Custom Resource. 

| Wavefront Proxy args | Observability for Kubernetes Operator Custom Resource `spec` |
|---|---|
|`--preprocessorConfigFile` | `dataExport.wavefrontProxy.preprocessor` ConfigMap |
|`--proxyHost` | `dataExport.wavefrontProxy.httpProxy.secret` Secret |
|`--proxyPort` | `dataExport.wavefrontProxy.httpProxy.secret` Secret |
|`--proxyUser` | `dataExport.wavefrontProxy.httpProxy.secret` Secret |
|`--proxyPassword` | `dataExport.wavefrontProxy.httpProxy.secret` Secret |
|`--pushListenerPorts` | `dataExport.wavefrontProxy.metricPort` |
|`--deltaCounterPorts` | `dataExport.wavefrontProxy.deltaCounterPort` |
|`--traceListenerPorts` | `dataExport.wavefrontProxy.tracing.wavefront.port` |
|`--traceSamplingRate` | `dataExport.wavefrontProxy.tracing.wavefront.samplingRate` |
|`--traceSamplingDuration` | `dataExport.wavefrontProxy.tracing.wavefront.samplingDuration` |
|`--traceZipkinListenerPorts` | `dataExport.wavefrontProxy.tracing.zipkin.port` |
|`--traceZipkinApplicationName` | `dataExport.wavefrontProxy.tracing.zipkin.applicationName` |
|`--traceJaegerListenerPorts` | `dataExport.wavefrontProxy.tracing.jaeger.port` |
|`--traceJaegerHttpListenerPorts` | `dataExport.wavefrontProxy.tracing.jaeger.httpPort` |
|`--traceJaegerGrpcListenerPorts` | `dataExport.wavefrontProxy.tracing.jaeger.grpcPort` |
|`--traceJaegerApplicationName` | `dataExport.wavefrontProxy.tracing.jaeger.applicationName` |
|`--histogramDistListenerPorts` | `dataExport.wavefrontProxy.histogram.port` |
|`--histogramMinuteListenerPorts` | `dataExport.wavefrontProxy.histogram.minutePort` |
|`--histogramHourListenerPorts` | `dataExport.wavefrontProxy.histogram.hourPort` |
|`--histogramDayListenerPorts` | `dataExport.wavefrontProxy.histogram.dayPort` |

Other supported Custom Resource configuration:
* `dataExport.wavefrontProxy.args`: Used to set any `WAVEFRONT_PROXY_ARGS` configuration not mentioned in the above table. See [wavefront-proxy-args.yaml](../../deploy/scenarios/wavefront-proxy-args.yaml) for an example.
* `dataExport.wavefrontProxy.resources`: Used to set container resource request or limits for Wavefront proxy. See [wavefront-pod-resources.yaml](../../deploy/scenarios/wavefront-pod-resources.yaml) for an example.
* `dataExport.externalWavefrontProxy.Url`: Used to set an external Wavefront proxy. See [wavefront-collector-external-proxy.yaml](../../deploy/scenarios/wavefront-collector-external-proxy.yaml) for an example.

### Wavefront Collector to Kubernetes Metrics Collector Configuration

Wavefront Collector `ConfigMap` changes:
* Wavefront Collector ConfigMap changed from `wavefront-collector` to `wavefront` namespace.
* `sinks.proxyAddress` changed from `wavefront-proxy.default.svc.cluster.local:2878` to `wavefront-proxy:2878`.
* Change `collector.yaml` to `config.yaml`

Custom Resource `spec` changes:
* Update Custom Resource configuration`dataCollection.metrics.customConfig` with the created ConfigMap name.
See [wavefront-collector-existing-configmap.yaml](../../deploy/scenarios/wavefront-collector-existing-configmap.yaml) for an example.

Other supported Custom Resource configurations:
* `dataCollection.metrics.nodeCollector.resources`: Used to set container resource request or limits for Wavefront node collector.
* `dataCollection.metrics.clusterCollector.resources`: Used to set container resource request or limits for Wavefront cluster collector.
See [wavefront-pod-resources.yaml](../../deploy/scenarios/wavefront-pod-resources.yaml) for an example.

### Future Support

If you have feedback, or come across something that cannot be configured with the Operator, please open an issue in this repository.
