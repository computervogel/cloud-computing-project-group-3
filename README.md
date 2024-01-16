# Cloud Computing Project Group 3

## Team members:
- Maximilian Fink
- Matteo Schweitzer
- Daniel Lovrinovic

## Project proposal
[Proposal](./PROPOSAL.md)

## Tutorial
### Prerequisites

1. [Docker](https://docs.docker.com/get-docker/)

1. Install `kind` by following the installation instructions provided in the [quick start](https://kind.sigs.k8s.io/docs/user/quick-start/)

1. Install the k6 test loader [Installation](https://grafana.com/docs/k6/latest/get-started/installation/)

1. Install the minikube [Installation](https://kubernetes.io/de/docs/tasks/tools/install-minikube/)

#### Creating a local Cluster

1. In your terminal, execute the following command:

   ```console
   kind create cluster
   ```

### Minikube

`minikube` minikube is local Kubernetes, focusing on making it easy to learn and develop for Kubernetes.

#### Starting a local Cluster

1. In your terminal, execute the following command:

   ```console
   minikube start
   ```

#### STARTING MICROSERIVES

### Requirements

- `kubectl` is connected to the Kubernetes cluster created

- Also run this command in order to track the behaviour of the autoscalers:

```console
   minikube addons enable metrics-server
```


### Start microservices

If you want to build the microservices by yourself and run them locally, then go to the [Microservice Instructions](microservices/INSTRUCTIONS_MICROSERVICES.md).

Open terminal and run:

```console
   kubectl apply -f <path_to_microservices.yaml>
```

### Time microservice

1. To forward your local port **9999** to the container port **7777**, use:

   ```console
   kubectl port-forward deployment/time 9999:7777
   ```

   :mag: Now, access the website [http://localhost:9999](http://localhost:9999).

### Fibonacci microservice

1. To forward your local port **9090** to the container port **8888**, use:

   ```console
   kubectl port-forward deployment/fibonacci 9090:8888
   ```

   :mag: Now, access the website [http://localhost:9090](http://localhost:9090).

### Factorial microservice

1. To forward your local port **9000** to the container port **5555**, use:

   ```console
   kubectl port-forward deployment/factorial 9000:5555
   ```

   :mag: Now, access the website [http://localhost:9000](http://localhost:9000).

### Prime microservice

1. To forward your local port **9990** to the container port **6666**, use:

   ```console
   kubectl port-forward deployment/prime 9990:6666
   ```

   :mag: Now, access the website [http://localhost:9990](http://localhost:9990).


#### STARTING VERTICAL AUTOSCALER

1. Open a terminal and run:

   ```console
   git clone https://github.com/kubernetes/autoscaler.git
   ```

   ```console
   cd autoscaler/vertical-pod-autoscaler/
   ```

   ```console
   ./hack/vpa-up.sh
   ```

   ```console
   
   kubectl apply -f https://raw.githubusercontent.com/kubernetes/autoscaler/vpa-release-1.0/vertical-pod-autoscaler/deploy/vpa-rbac.yaml

   kubectl apply -f https://raw.githubusercontent.com/kubernetes/autoscaler/vpa-release-1.0/vertical-pod-autoscaler/deploy/vpa-v1-crd-gen.yaml

   kubectl apply -f <path_to_vpa.yaml>
   ```

1. to check if the autoscaler is running, run:

   ```console
   kubectl get pods -n kube-system
   ```
   
1. to see the autoscaler working, run:

   ```console
   kubectl describe vpa
   ```
   
1. to stop the autoscaler, run:

   ```console
   ./hack/vpa-down.sh
   ```

#### Output of the vertical pod autoscaler:
This screenshot shows how the cpu load changes by the time the vpa has to handle a big amount of load. If you look at the "Recommendations" and then at the "Target" you will first see 25m, but as the load increases it goes up to 50m, which is the given maximum by our vertical-autoscaler.yaml-file.

![vpa](vpa.png)


#### STARTING HORIZONTAL AUTOSCALER

1. Open a terminal and run:

   ```console
   kubectl autoscale deployment <name> --cpu-percent=50 --min=1 --max=10
   ```
   or
   ```console
   kubectl apply -f <path_to_hpa.yaml>
   ```

1. to check if the autoscaler is running, run:

   ```console
   kubectl get hpa
   ```
1. to see the load, run:

   ```console
   kubectl get hpa <name> --watch
   ```

1. to stop the autoscaler, run:

   ```console
   kubectl delete hpa <name>
   ```

#### Output of setting up the connection:

![Connection](Connection.png)
   

#### Output of the pod load:
Here you can see pretty good how the horizontal pod autoscaler handles too much load and generates automatically more pods to reduce the load.

![Pod load](PodLoad.png)


#### Output of the pod scaling:

![Scaling the pods](Scale.png)


#### LOAD-TESTER

1. K6 load-tester is used for load-generating:

```console
k6 run <path_to_respective_test.js>
```

The load tester automatically runs the respective tests with the given parameters inside the test files (e.g. testFibonacci). In our case the test repeatedly goes to the dedicated localhost-URL to generate load onto the deployment, which is tested. "vus" stands for virtual users and regulates how many "Users" use the services simultanouesly, defining the load. "duration" is simply the time the test runs for.

SIDENOTE: Both autoscalers are currently configured to test the fibonacci-microservice. If you want to test other services, you simply have to change the name-parameters inside the autoscalers to the desired deployment name. After changing the names, you have to run the following command again:

for horizontal autoscaler:
   ```console
   kubectl apply -f <path_to_hpa.yaml>
   ```
for vertical autoscaler:
   ```console
   kubectl apply -f <path_to_vpa.yaml>
   ```

## Lessons learned
-Kubernetes allows great management of microservices, especially when using autoscalers.

-Microservices are very useful regarding the robustness of for example a website. If one microservice gets heavy traffic, the other services are effected minimally.

-Autoscalers, no matter which type, can help in extreme real-life situations, where a service gets much more traffic than usual, but it is important that it still runs smoothly throughout the heavy-    load time

-Even though microservices have many benefits, it requires more setup time, than simply coding all the services in one file, without the use of kubernetes. The developers have to decide if the additional steps, needed to setup a kubernetes cluster, are worth it in the long run.
