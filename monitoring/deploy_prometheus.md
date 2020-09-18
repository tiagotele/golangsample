# Installing MyApp via Helm Chart
helm3 install myapp .

# Installing Prometheus
```bash
kubectl create namespace monitoring  
helm3 install prometheus stable/prometheus --set alertmanager.service.type=NodePort --set server.service.type=NodePort --namespace monitoring
```

helm3 upgrade prometheus stable/prometheus  -f ~/git/charts/stable/prometheus/values.yaml  --set alertmanager.service.type=NodePort --set server.service.type=NodePort --namespace monitoring

# Instaling Grafana

```bash
helm3 repo add bitnami https://charts.bitnami.com/bitnami
helm3 install grafana bitnami/grafana --version 3.4.1 --set service.type="NodePort"
```

> Add Prometheus URL by hand  (IP From service)
> Import dashboards  
> Add single dashboard
