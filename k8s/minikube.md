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