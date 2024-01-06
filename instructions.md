### Prerequisites

1. [Docker](https://docs.docker.com/get-docker/)

1. Install `kind` by following the installation instructions provided in the [quick start](https://kind.sigs.k8s.io/docs/user/quick-start/)

### Creating a local Cluster

1. In your terminal, execute the following command:

   ```console
   kind create cluster
   ```

1. Print the cluster info by executing the following command:

   ```console
   kubectl cluster-info --context kind-kind
   ```

## minikube

`minikube` minikube is local Kubernetes, focusing on making it easy to learn and develop for Kubernetes."

### Starting a local Cluster

1. In your terminal, execute the following command:

   ```console
   minikube start
   ```

1. Print the cluster info by executing the following command:

   ```console
   kubectl cluster-info
   ```

### STARTING MICROSERIVES

## Requirements

- `kubectl` is connected to the Kubernetes cluster created

## Time microservice

1. In Kubernetes, the equivalent to `docker container run` is `kubectl run`. Use this command to run your container:

   ```console
   kubectl run time --image=matteoschweitzer/time:0.0.1
   ```

1. To verify that the container has started and the app is running, use:

   ```console
   kubectl get pods
   ```

1. To forward your local port **9999** to the container port **7777**, use:

   ```console
   kubectl port-forward pod/time 9999:7777
   ```

   :mag: Now, access the website [http://localhost:9999](http://localhost:9999).

1. To delete the pod, use:

   ```console
   kubectl delete pod time
   ```

## Fibonacci microservice

1. In Kubernetes, the equivalent to `docker container run` is `kubectl run`. Use this command to run your container:

   ```console
   kubectl run fibonacci --image=matteoschweitzer/fibonacci:0.0.1
   ```

1. To verify that the container has started and the app is running, use:

   ```console
   kubectl get pods
   ```

1. To forward your local port **9090** to the container port **8888**, use:

   ```console
   kubectl port-forward pod/fibonacci 9090:8888
   ```

   :mag: Now, access the website [http://localhost:9090](http://localhost:9090).

1. To delete the pod, use:

   ```console
   kubectl delete pod fibonacci
   ```

## Factorial microservice

1. In Kubernetes, the equivalent to `docker container run` is `kubectl run`. Use this command to run your container:

   ```console
   kubectl run factorial --image=matteoschweitzer/factorial:0.0.1
   ```

1. To verify that the container has started and the app is running, use:

   ```console
   kubectl get pods
   ```

1. To forward your local port **9000** to the container port **5555**, use:

   ```console
   kubectl port-forward pod/factorial 9000:5555
   ```

   :mag: Now, access the website [http://localhost:9000](http://localhost:9000).

1. To delete the pod, use:

   ```console
   kubectl delete pod factorial
   ```

## prime microservice

1. In Kubernetes, the equivalent to `docker container run` is `kubectl run`. Use this command to run your container:

   ```console
   kubectl run prime --image=matteoschweitzer/prime:0.0.1
   ```

1. To verify that the container has started and the app is running, use:

   ```console
   kubectl get pods
   ```

1. To forward your local port **9990** to the container port **6666**, use:

   ```console
   kubectl port-forward pod/prime 9990:6666
   ```

   :mag: Now, access the website [http://localhost:9990](http://localhost:9990).

1. To delete the pod, use:

   ```console
   kubectl delete pod prime
   ```
