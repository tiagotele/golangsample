# Installing Prometheus
kubectl create namespace monitoring
helm3 install prometheus stable/prometheus --namespace monitoring
