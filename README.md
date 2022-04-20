# grpc-gateway-template


# Namespace demo

Steps to run the Namespace demo

```code
# check out namespace_basic_no_connect branch
git co namespace_basic_no_connect

# Next. This will generate protobuf files
make

# Run code
go run main.go
```










# Loading Cluster

## Step 1:

```bash
make complete-kind

```

## Step 2:

Test with Nginx

```bash
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  labels:
    run: nginx
  name: nginx
spec:
  containers:
  - name: nginx 
    image: nginx
    ports:
    - containerPort: 80
EOF


```

## Step 3:

LoadBalancer

```bash
kubectl expose pod nginx --type=LoadBalancer

```
