package prometheus

const Cluster_All_Pods_Cpu_Usage = `sum(rate(container_cpu_usage_seconds_total{container!="",container!="POD"}[3m]))`
