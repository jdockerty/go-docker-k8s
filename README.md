# Task Scheduler Dashboard

Small personal project for using Golang, Docker, and Kubernetes. Tutorial by [callicoder](https://www.callicoder.com/deploy-containerized-go-app-kubernetes/)

### Notes

### Docker

Dockerfile is commented to provide explanation of each layer.

Starting with the command `docker build -t gokubedash .`  

This builds the image using the Dockerfile, whilst tagging it as _'gokubedash'_, everything from the current directory, which is not part of the .dockerignore, is included in the image. 


The resulting image was then tagged with _jdockerty/gokubedash:1.1_ as this provides the _jdockerty/_ namespace in a DockerHub. _(This may not have been entirely necessary upon reflection, as you can add multiple tags in the previous step.)_   
`docker tag gokubedash jdockerty/gokubedash:1.1`  


The image was then run as a container to test that it was functioning properly. This maps the localhost (127.0.0.1) port 8080, to the port which is exposed from the container, which is also 8080. `docker run -p 127.0.0.1:8080:8080 gokubedash`

After this, I pushed the image to DockerHub for use with Kubernetes. `docker push jdockerty/gokubedash:1.1`

### Kubernetes
Beforehand, I had to install [Kubernetes](https://kubernetes.io/docs/tasks/tools/install-kubectl/) and [Minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/). Kubernetes is for running containers at scale and Minikube is so that a Kubernetes cluster can be run locally.

The process begins with `minikube start`

A YAML file, known as a manifest, is used to start Kubernetes deployments. This is the _k8sconfig.yml_ file.

To begin the K8s deployment, we issue the command `kubectl apply -f k8sconfig.yml`, the -f flag is used to denote the file we are using for our deployment's configuration. This applies our resources, the deployment and a service.

We can see the deployments using the command `kubectl get deployments`, of which the output shown below.
![Get deployments K8s command](https://github.com/jdockerty/taskschedulerdashboard/blob/master/k8simages/get%20deployments.png)

The pods can also be seen using the command `kubectl get pods`.
![Get pods K8s command](https://github.com/jdockerty/taskschedulerdashboard/blob/master/k8simages/get%20pods.png)

Since the pods within the cluster cannot be accessed from outside, by default, we must port-forward one of the pods to test it is working. `kubectl port-forward gokubedash-6bf64bdc89-4r89s 8080:8080`

![Port forwarding K8s](https://github.com/jdockerty/taskschedulerdashboard/blob/master/k8simages/port%20forward.png)

We can then visit localhost:8080 to test whether it is working properly.
![Testing K8s port forward](https://github.com/jdockerty/taskschedulerdashboard/blob/master/k8simages/port%20forward%20testing.png)

Kubernetes allows for very fast scaling, we increase the number of pods that are running by issuing the command 

`kubectl scale --replicas=5 deployment/gokubedash` 

We initial set the number of pods (replicas) to 3, so now we are scaling it to 5 with a single command.

![Scaling K8s](https://github.com/jdockerty/taskschedulerdashboard/blob/master/k8simages/scaling%20k8.png)

### Elastic Kubernetes Service (Amazon EKS)
_This side implementation was done out of interest from using EKS within AWS and wasn't part of the original tutorial post. This was done through AWS documentation and [this post](https://codeburst.io/getting-started-with-kubernetes-deploy-a-docker-container-with-kubernetes-in-5-minutes-eb4be0e96370)_

Firstly, the Docker image we have for our _gokubedash_ was uploaded into AWS Elastic Container Registry (ECR). We have to log into ECR before we can do this, the credentials are provided using the command, the ECR registry had been created prior to this. `aws ecr get-login --no-include-email --region eu-west-2`

The output of this command is then used to log into the ECR registry 

`docker login -u AWS -p <password output from previous command> <ECR domain>`

From here we simply tag the current image we were using and then push it into ECR.

`docker tag jdockerty/go-k8s:latest <ECR domain>/gokubedash:latest`

`docker push <ECR domain>/gokubedash:latest`

We created the EKS cluster and node group on AWS. This was done through following the [getting started guide](https://docs.aws.amazon.com/eks/latest/userguide/getting-started-console.html) and Management Console prompts.

![EKS Cluster console](https://github.com/jdockerty/taskschedulerdashboard/blob/master/k8simages/EKS%20cluster.png)

From there, we must update the EKS cluster with the AWS CLI, this was done through the command 

`aws eks --region eu-west-2 update-kubeconfig --name NewCluster` 

This updated the EKS cluster with the k8sconfig.yml we had on our local machine.

From there we simply need to run the command 

`kubectl run gokubedash --image=<ECR domain>/gokubedash:latest --port=8080`

This runs our application, under the name gokubedash, using the image provided inside of ECR, on port 8080, which is the port exposed through the Go code.

From here we can issue `kubectl get pods` to see the running pods.

![EKS Get pods](https://github.com/jdockerty/taskschedulerdashboard/blob/master/k8simages/get%20nodes%20EKS.png)

Since we know that the deployment cannot be accessed from outside yet, as the pods are given private addresses within a cluster, it must be exposed to the outside world using a load balancer. This can be achieved with the command

`kubectl expose deployment gokubedash --type=LoadBalancer --port=8080 --target-port=8080 `

This exposes our deployment with the load balancer listening on port 8080, forwarding the traffic to 8080 as well, which is what our application is listening on for connections.

The load balancer is given an external address for us to access, we can see this by using the command 

`kubectl get svc`

![EKS kubectl get svc](https://github.com/jdockerty/taskschedulerdashboard/blob/master/k8simages/EKS%20get%20svc.png)

We can then browse to the provided external domain on port 8080 and view our application from one of the pods.

![EKS app working](https://github.com/jdockerty/taskschedulerdashboard/blob/master/k8simages/EKS%20app%20working.png)
