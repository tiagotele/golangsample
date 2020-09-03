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

# To record info about deployment
# kubectl create -f deployment.yml --record

kubectl rollout undo deployment/myapp-deployment
