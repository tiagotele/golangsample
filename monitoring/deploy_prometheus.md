# Installing Prometheus
```bash
kubectl create namespace monitoring  
helm3 install prometheus stable/prometheus --set alertmanager.service.type=NodePort --set server.service.type=NodePort --namespace monitoring
```

> Latest running version
2020-09-10 21:32:39.686903 -0300 -03	deployed	prometheus-11.12.1	2.20.1

# Instaling Grafana

```bash
helm3 repo add bitnami https://charts.bitnami.com/bitnami
helm3 install grafana bitnami/grafana --version 3.4.1 --set service.type="NodePort"
```

> Add Prometheus URL by hand  
> Add single dashboard

# Installing MyApp via Helm Chart
helm3 install myapp myapp-chart