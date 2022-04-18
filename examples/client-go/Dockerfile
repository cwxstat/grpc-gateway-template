FROM golang:latest as builder

WORKDIR /project

COPY . /project
# Install Delve for debugging
RUN go install github.com/go-delve/delve/cmd/dlv@latest

ENV GO111MODULE on
ENV DEBUG true

ENTRYPOINT ["sleep", "100d"]

#  docker build --no-cache -t clientgo:v0.0.1 -f Dockerfile .
#  kind load --name=metal docker-image clientgo:v0.0.1

#  kubectl create deployment --image=clientgo:v0.0.1 clientgo

#  kubectl port-forward --address 0.0.0.0 webdev 2345:2345 
#  kubectl port-forward --address 0.0.0.0 webdev 8080:8080
# kubectl expose deployment webdev --port=8080 --target-port=8080
# kubectl port-forward webdev 2345:2345
# kubectl exec -it webdev -- /bin/bash
# dlv debug ./main.go --listen=0.0.0.0:2345 --api-version=2  --headless