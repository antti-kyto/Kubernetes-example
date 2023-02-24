# go Docker & Kubernetes example 

Running minikube on WSL2

## Docker 

Build Docker image:
```console
docker build -t my-golang-api .
```

Run Docker container (use -d to Run container in background ):
```console
docker run -p 127.0.0.1:3000:3000/tcp my-golang-api
```

Login Docker hub:
```console
docker login --username={username}
```

Tag image:
```console
docker tag {imageId} {username}/{reponame}:{tag}
```

Push:
```console
docker push {username}/{reponame}:{tag}
```

## Kubernetes

Start minikube:
```console
# If minikube was started without ports, run:
minikube delete
# Then
minikube start --ports=127.0.0.1:30007:30007
```

Apply k8s files:
```console
kubectl apply -f k8s/myapp-config.yaml  # ConfigMap
kubectl apply -f k8s/myapp.yaml         # Deployment & Service
```

if ports not defined on start:
```console
minikube service my-api
```

Test:
```console
curl http://127.0.0.1:30007/
```