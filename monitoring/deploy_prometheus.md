# Installing Prometheus
kubectl create namespace monitoring
helm3 install prometheus stable/prometheus --set alertmanager.service.type=NodePort --set server.service.type=NodePort --namespace monitoring

# Latest running version
2020-09-10 21:32:39.686903 -0300 -03	deployed	prometheus-11.12.1	2.20.1