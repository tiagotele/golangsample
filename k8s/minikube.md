# Create minikube cluster

minikube start --vm-driver=virtualbox

## Show nodes, plugins


kubectl create -f pod.yml
kubectl get pods
## Create pod

## Show ways to expose a pod

### Port-forward
kubectl port-forward myapp 8080:8080
Then access http://localhost:8080/

### Expose via service 

kubectl expose pod myapp --type=NodePort

minikube service hello-minikube --url

minikube service list

## Replication Controller

kubectl create -f replic-controller.yml 

### Scale via overriding yml file

kubectl scale --replicas=4 -f replic-controller.yml

### Scale via replica controller name
kubectl scale --replicas=1 rc/myapp


## Deployment
kubectl create -f deployment.yml # 2.0 version

kubectl get deployments
kubectl get rs

kubectl rollout status deployment/myapp-deployment

kubectl set image deployment/myapp-deployment myapp=tiagotele/golangsample:3.0

kubectl rollout status deployment/myapp-deployment

kubectl rollout history deployment/myapp-deployment

> To record info about deployment
> kubectl create -f deployment.yml --record

kubectl rollout undo deployment/myapp-deployment

## Services
Pods are acessed by services, not directly
Logcial bridge between "mortal" pods and other servies or end-users.

kubectl expose creates a service

Types:
- ClusterIP: default. reachable within cluster
- NodePort: 
- LoadBalancer

## Labels

You can tag resource(s)
Not unique: you use mulitple lables.

Used on Node selector.

First add label to node

```
get nodes --show-labels
```

To add label to a node
```
kubectl label nodes node1 key=value
```

Secondly add pod to node that use that label.
```
kubectl label nodes minikube turma=uni7
```

## Health checks

2 ways:
- run command on container periodically
- periodic checks URL(HTTP)

## Healthchecks and Readness
- Healthcheck see if container is running
- Readiness indicate if container is ready to serve requests.



## Pod states

### Pod status
High level

- pending
- succeeded
- failed
- unknown

kubectl describe pod to get info about pod

### 5 pod condition
Contition pod has passed.

- PodScheduled
- Ready
- Initialized
- Unschedulable
- ContainersReady

### containers state
- Running
- Terminated
- Waiting

# Secrets

Can be environment variables
Can be environment volume

# Configmap
https://github.com/fabiogomezdiaz/spring-boot-configmap-demos/blob/master/1-basic-configmap/k8s/deployment.yaml