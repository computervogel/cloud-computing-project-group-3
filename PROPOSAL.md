# Horizontal Pod Autoscaler and Vertical Pod Autoscaler

## Definition and Scope
The project consists of a small self-built website with 3-4 services. These services will be autoscaled using a horizontal/vertical pod autoscaler. Testing will be done by a loadgenerator. Through logs we will be able to check the resource allocation/number of pods. So we are just checking if the autoloades is working correctly. If enough time is left we try to implement a multidimensional autoscaler. Main resource will probably be: https://github.com/kubernetes/autoscaler

## Webiste
We implemented a simple website with buttons that are connected to simple services. It can for example change the colour of the output from black to red, show the current time or print out a specitic amount of fibonacci numbers.

## Milestones
-Design the [website](#website)

-Deploy the working [website](#website)

-Research autoscaling (Tools and Concepts)

-Implement the autoscalers
