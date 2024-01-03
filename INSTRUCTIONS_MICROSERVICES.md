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

* `kubectl` is connected to the Kubernetes cluster created in the previous [exercise 3.1](../exercise%203.1)

## Time microservice

1. In Kubernetes, the equivalent to `docker container run` is `kubectl run`. Use this command to run your container:

    ```console
    kubectl run time --image=agrimmer/demo:time
    ```

1. To verify that the container has started and the app is running, use:
    
    ```console
    kubectl get pods
    ```

1. To forward your local port **9999** to the container port **8888**, use:
    
    ```console
    kubectl port-forward pod/time 9999:8888
    ```

    :mag: Now, access the website [http://localhost:9999](http://localhost:9999). What is shown there? 

1. To delete the pod, use:

    ```console
    kubectl delete pod demo
    ```
