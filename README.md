# grpc-gateway-template


# Namespace demo

Steps to run the Namespace demo

```bash
# check out namespace_basic_no_connect branch
git co namespace_basic_no_connect

# Next. This will generate protobuf files
make

# Run code
go run main.go
```

After running the above, you should see the following:

```bash
INFO: 2022/04/20 18:18:47 Serving gRPC on https://0.0.0.0:10000
INFO: 2022/04/20 18:18:47 [core] parsed scheme: "dns"
INFO: 2022/04/20 18:18:47 [core] ccResolverWrapper: sending update to cc: {[{0.0.0.0:10000  <nil> 0 <nil>}] <nil> <nil>}
INFO: 2022/04/20 18:18:47 [core] ClientConn switching balancer to "pick_first"
INFO: 2022/04/20 18:18:47 [core] Channel switches to new LB policy "pick_first"
INFO: 2022/04/20 18:18:47 [core] Subchannel Connectivity change to CONNECTING
INFO: 2022/04/20 18:18:47 [core] Subchannel picks a new address "0.0.0.0:10000" to connect
INFO: 2022/04/20 18:18:47 [core] Channel Connectivity change to CONNECTING
INFO: 2022/04/20 18:18:47 [core] Subchannel Connectivity change to READY
INFO: 2022/04/20 18:18:47 [core] Channel Connectivity change to READY
INFO: 2022/04/20 18:18:47 Serving gRPC-Gateway and OpenAPI Documentation on https://0.0.0.0:11000

```



You can attach to the http interface `https://0.0.0.0:11000`

<img width="885" alt="image" src="https://user-images.githubusercontent.com/755710/164297030-c4757a15-474e-4875-9c4f-fd0c94901a7d.png">



### gRPC run on 10000

```bash
# Need to be in client directory to pick up certs
cd client
go run main.go 
2022/04/20 18:24:21 Create time: seconds:1650479061 nanos:70050683
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
