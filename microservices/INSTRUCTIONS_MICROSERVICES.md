### THIS TUTORIAL IS ONLY NEEDED IF YOU WANT TO RUN THE MICROSERVICES WITHOUT MINIKUBE

## Fibonacci Sequence

1. Open Integrated Terminal for Microservice:

    click on the fibonacci folder choose "Open in Integrated Terminal"

1. Create image from the provided Dockerfile:

    ```console
    docker image build -t [YOUR-DOCKERHUB-ACCOUNT]/fibonacci:0.0.1 ./
    ```

1. Run the a container from the image and expose the container port: **8888** to the host port: **9090**.

    ```console
    docker container run -p 9090:8888 [YOUR-DOCKERHUB-ACCOUNT]/fibonacci:0.0.1
    ```

    * *Optional*: Run the container in *detached mode* shown by the option `--detach` or `-d`, meaning that a Docker container runs in the background of your terminal. It does not receive input or display output.

    ```console
    docker container run -d -p 9090:8888 [YOUR-DOCKERHUB-ACCOUNT]/fibonacci:0.0.1
    ```

1. Open a browser and go to: http://localhost:9090

1. See your container running on your local Docker daemon:

    ```console
    docker ps
    ```
    
    ```
    CONTAINER ID        IMAGE                    COMMAND             CREATED             STATUS              PORTS                    NAMES
    789d08da1704        xyz/fibonacci:0.0.1        "/usr/myapp"        21 seconds ago      Up 19 seconds       0.0.0.0:9999->8888/tcp   foxy_joliot
    ```

1. Stop your container:

    ```console
    docker stop 789d08da1704
    ```


## Show Time

1. Open Integrated Terminal for Microservice:

    click on the time folder choose "Open in Integrated Terminal"

1. Create image from the provided Dockerfile:

    ```console
    docker image build -t [YOUR-DOCKERHUB-ACCOUNT]/time:0.0.1 ./
    ```

1. Run the a container from the image and expose the container port: **7777** to the host port: **9999**.

    ```console
    docker container run -p 9999:7777 [YOUR-DOCKERHUB-ACCOUNT]/time:0.0.1
    ```

    * *Optional*: Run the container in *detached mode* shown by the option `--detach` or `-d`, meaning that a Docker container runs in the background of your terminal. It does not receive input or display output.

    ```console
    docker container run -d -p 9999:7777 [YOUR-DOCKERHUB-ACCOUNT]/time:0.0.1
    ```

1. Open a browser and go to: http://localhost:9999

1. See your container running on your local Docker daemon:

    ```console
    docker ps
    ```
    
    ```
    CONTAINER ID        IMAGE                    COMMAND             CREATED             STATUS              PORTS                    NAMES
    789d08da1704        xyz/time:0.0.1        "/usr/myapp"        21 seconds ago      Up 19 seconds       0.0.0.0:9999->8888/tcp   foxy_joliot
    ```

1. Stop your container:

    ```console
    docker stop 789d08da1704
    ```

## Prime Numbers

1. Open Integrated Terminal for Microservice:

    click on the prime folder choose "Open in Integrated Terminal"

1. Create image from the provided Dockerfile:

    ```console
    docker image build -t [YOUR-DOCKERHUB-ACCOUNT]/prime:0.0.1 ./
    ```

1. Run the a container from the image and expose the container port: **6666** to the host port: **9990**.

    ```console
    docker container run -p 9990:6666 [YOUR-DOCKERHUB-ACCOUNT]/prime:0.0.1
    ```

    * *Optional*: Run the container in *detached mode* shown by the option `--detach` or `-d`, meaning that a Docker container runs in the background of your terminal. It does not receive input or display output.

    ```console
    docker container run -d -p 9990:6666 [YOUR-DOCKERHUB-ACCOUNT]/prime:0.0.1
    ```

1. Open a browser and go to: http://localhost:9990

1. See your container running on your local Docker daemon:

    ```console
    docker ps
    ```
    
    ```
    CONTAINER ID        IMAGE                    COMMAND             CREATED             STATUS              PORTS                    NAMES
    789d08da1704        xyz/prime:0.0.1        "/usr/myapp"        21 seconds ago      Up 19 seconds       0.0.0.0:9999->8888/tcp   foxy_joliot
    ```

1. Stop your container:

    ```console
    docker stop 789d08da1704
    ```

## Factorial Numbers

1. Open Integrated Terminal for Microservice:

    click on the factorial folder choose "Open in Integrated Terminal"

1. Create image from the provided Dockerfile:

    ```console
    docker image build -t [YOUR-DOCKERHUB-ACCOUNT]/factorial:0.0.1 ./
    ```

1. Run the a container from the image and expose the container port: **5555** to the host port: **9000**.

    ```console
    docker container run -p 9000:5555 [YOUR-DOCKERHUB-ACCOUNT]/factorial:0.0.1
    ```

    * *Optional*: Run the container in *detached mode* shown by the option `--detach` or `-d`, meaning that a Docker container runs in the background of your terminal. It does not receive input or display output.

    ```console
    docker container run -d -p 9000:5555 [YOUR-DOCKERHUB-ACCOUNT]/factorial:0.0.1
    ```

1. Open a browser and go to: http://localhost:9000

1. See your container running on your local Docker daemon:

    ```console
    docker ps
    ```
    
    ```
    CONTAINER ID        IMAGE                    COMMAND             CREATED             STATUS              PORTS                    NAMES
    789d08da1704        xyz/factorial:0.0.1        "/usr/myapp"        21 seconds ago      Up 19 seconds       0.0.0.0:9999->8888/tcp   foxy_joliot
    ```

1. Stop your container:

    ```console
    docker stop 789d08da1704
    ```
