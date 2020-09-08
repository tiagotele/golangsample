# About Helm v3
helm3 repo list
helm3 plugin install https://github.com/helm/helm-2to3
helm3 2to3
helm3 2to3 move config -h
helm3 2to3 move config --dry-run
helm3 2to3 move config 
helm3 repo list
helm3 plugin list
helm3 search repo stable
helm3 repo update
helm3 search repo stable

# Installing Prometheus
kubectl create namespace monitoring
helm3 install prometheus stable/prometheus --namespace monitoring

