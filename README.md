## Word Count Application

A basic HTTP Golang application that downloads a book and counts the words.

To build:

```
go build *.go
```

To run:

```
go run *.go
```

### Run from Docker:

To build:

```
docker build . -t text-analytics
```

To run on http://localhost:8080:

```
docker run -p 8080:7979 -it text-analytics
```

### Run from Minikube:

Start Minikube:

```
minikube start
```

To re-use the Docker daemon inside the Minikube instance (otherwise Minikube is not able to find the image):

```
eval $(minikube docker-env)
```

To build the Docker image with a tag:

```
docker build . -t text-analytics
```

To enable the Ingress extension in Minikube:

```
minikube addons enable ingress
```

To create mandatory resources for Ngnix Ingress in the Minikube cluster:

```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/nginx-0.28.0/deploy/static/mandatory.yaml
```

To check the running pods in the system for Ngnix Ingress controller:

```
kubectl get pods -n kube-system
```

To set a host name locally for the Minikube IP:

```
echo "$(minikube ip) hello.world" | sudo tee -a /etc/hosts
```

To deploy the Text Analytics (deployment and service) together with Load Balancer:

```
kubectl apply -f ./kubernetes/
```

After this step curl on the initialized path must be working :)

```
curl http://hello.world/text-collect
```
