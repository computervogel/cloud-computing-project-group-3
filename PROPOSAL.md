# Horizontal Pod Autoscaler and Vertical Pod Autoscaler

## Definition and Scope
The project consists of a small self-built website with 3-4 services. These services will be autoscaled using a horizontal/vertical pod autoscaler. Testing will be done by a loadgenerator. Through logs we will be able to check the resource allocation/number of pods. So we are just checking if the autoloades is working correctly. If enough time is left we try to implement a multidimensional autoscaler. Main resource will probably be: https://github.com/kubernetes/autoscaler

## Website
We implemented a simple website with buttons that are connected to simple services. It can for example change print prime numbers, show the current time or print out the first 40 fibonacci numbers.

## Autoscaler
For our project we use the K6 autoscaler.

## Milestones
-Design the [website](#website)

-Deploy the working [website](#website)

-Research autoscaling (Tools and Concepts)

-Implement the autoscalers

## Responsibilities
-Horizontal Pod Autoscaler - Maximilian Fink
-Vertical Pod Autoscaler - Daniel Lovrinovic
-Website, Tests and Microservices - Matteo Schweitzer

But in the end, we all worked together on all of these topics and implemented and tested them.
