local kp =
  (import 'kube-prometheus/kube-prometheus.libsonnet') +
  (import 'kube-prometheus/kube-prometheus-all-namespaces.libsonnet') +
  {
    _config+:: {
      namespace: 'monitoring',

      prometheus+:: {
        namespaces: [],
      },
    },
    prometheus+:: {
      serviceMonitorTraefik: {
        apiVersion: 'monitoring.coreos.com/v1',
        kind: 'ServiceMonitor',
        metadata: {
          name: 'traefik-monitor-service',
          namespace: 'default',
        },
        spec: {
          jobLabel: 'traefik',
          endpoints: [
            {
              port: 'metrics',
            },
          ],
          selector: {
            matchLabels: {
              'app': 'traefik',
            },
          },
        },
      },
      serviceMonitorDefault: {
        apiVersion: 'monitoring.coreos.com/v1',
        kind: 'ServiceMonitor',
        metadata: {
          name: 'vitess-monitor-service',
          namespace: 'default',
        },
        spec: {
          jobLabel: 'vitess-cluster',
          endpoints: [
            {
              port: 'metrics',
            },
          ],
          selector: {
            matchLabels: {
              'planetscale.com/component': 'vttablet',
            },
          },
        },
      },
    },
    grafanaDashboards+:: {
      'mysql-overview.json': (import '../../grafana-dashboards/dashboards/MySQL_Overview_mod.json'),
    },
    grafana+:: {
      dashboards+:: {
        'mysql-overview.json': (import '../../grafana-dashboards/dashboards/MySQL_Overview_mod.json'),
      },
    },
  };

{ ['setup/0namespace-' + name]: kp.kubePrometheus[name] for name in std.objectFields(kp.kubePrometheus) } +
{
  ['setup/prometheus-operator-' + name]: kp.prometheusOperator[name]
  for name in std.filter((function(name) name != 'serviceMonitor'), std.objectFields(kp.prometheusOperator))
} +
{ 'prometheus-operator-serviceMonitor': kp.prometheusOperator.serviceMonitor } +
{ ['node-exporter-' + name]: kp.nodeExporter[name] for name in std.objectFields(kp.nodeExporter) } +
{ ['kube-state-metrics-' + name]: kp.kubeStateMetrics[name] for name in std.objectFields(kp.kubeStateMetrics) } +
{ ['alertmanager-' + name]: kp.alertmanager[name] for name in std.objectFields(kp.alertmanager) } +
{ ['prometheus-' + name]: kp.prometheus[name] for name in std.objectFields(kp.prometheus) } +
{ ['prometheus-adapter-' + name]: kp.prometheusAdapter[name] for name in std.objectFields(kp.prometheusAdapter) } +
{ ['grafana-' + name]: kp.grafana[name] for name in std.objectFields(kp.grafana) }
