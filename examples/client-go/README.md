

```bash

k create ns client-go
kubens client-go
kubectl create clusterrolebinding default-view --clusterrole=view --serviceaccount=client-go:default

kubectl create clusterrolebinding cluster-admin --clusterrole=cluster-admin --serviceaccount=client-go:default
#




docker build --no-cache -t clientgo:v0.0.1 -f Dockerfile .
kind load --name=metal docker-image clientgo:v0.0.1
kubectl create deployment --image=clientgo:v0.0.1 clientgo
```