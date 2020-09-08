helm install --name elasticsearch --version 7.9.1 elastic/elasticsearch --set replicas=1

helm install --name kibana --version 7.9.1 elastic/kibana --set resources.requests.cpu="200m" --set resources.limits.cpu="800m" --set resources.requests.memory="1Gi" 


helm install --name fluentd stable/fluentd-elasticsearch --set  elasticsearch.host="elasticsearch-master"