# EECS4222: Distributed Systems: Project 2 (25%)

**Title:** Distributed Chat Application

**Start:** March 05, 2023, @ 9:00am EST  

**Due:** April 09, 2023, by 9:00pm EST

## Objective
This project aims to design and implement a distributed chat application that supports multi-client and real-time communication. The chat application will be built using a distributed architecture and deployed on Kubernetes to ensure reliable and efficient communication between clients, high scalability, and fault tolerance. The application will provide a web user interface for users to send and receive messages.

## Prerequisites
The following prerequisites must be met to successfully complete this project.

* A Kubernetes cluster (as required by Project 1)
* Docker >= 19.03.8 (as required by Project 1; Hint: see [Appendix 4](#4-docker-tutorials--links--docs) for Docker Tutorials/Docs)
* Go >= 1.19
* Python >= 3.8.0 (if applicable)
* JavaScript (if applicable)
* Access to the private Docker Registry [`harbor.pacslab.ca`](https://harbor.pacslab.ca) (obtain your credential [here](https://forms.gle/JhwQTvFPnUAcnhRZ6))

## Requirements
The overall requirements fall into three parts.

### Part I: Implement a Distributed Chat Application
For Part I, the overall requirements include the following:
1. Implement a distributed chat application that supports multi-client and real-time communication **using Go**. When a client sends a new message, other clients can see it immediately.
2. Use the WebSocket protocol to achieve real-time communications.
3. The chat application should have a web user interface where users can read and send messages.
4. Each message should at least contain the following information.
    * Name of the sender
    * Email of the sender
    * Time when the message was sent
    * Topic of the message
    * Content of the message
5. Define Kubernetes objects in `chatapp.yaml` to properly coordinate components of your application.
6. Use node port `30222` to expose the web user interface.
7. The whole application should be up and running using the following commands:
    ```
    kubectl create namespace chatapp
    kubectl apply -f chatapp.yaml -n chatapp
    ```
    The web user interface should be accessible by going to the address `http://MASTER_IP:30222`.

Hint: see [Appendix 1](#1-example-chat-application)

### Part II: Feature Implementation
Part II involves implementing an additional feature for the distributed chat application you built for Part I. You may choose **one** feature from the following list and implement for this part. You need to update `info.txt` accordingly (see [Technicalities](#technicalities)).

1. Add user registration and authentication features. Only authenticated users can send messages, while guests can only view messages. On the web user interface, display the number of **presently online** guests and registered users separately.
2. Provide the ability to allow users to create private chat rooms visible only to users who have been added to the room's participants list by the creator.
3. Provide the ability to send ephemeral messages. Such messages will **not** be stored on the server, and they will just be broadcast in real-time to whoever has the web user interface open and receiving messages right now.
4. Provide the ability to allow a user to write replies and vote messages up or down. Display the replies and the number of upvotes and downvotes of each message on the web user interface.
5. Provide the ability to allow a user to receive messages based on one or more self-defined conditions (i.e., contain a specific hashtag/keyword, filters based on the time when messages were sent, or filters based on search expressions)
6. Persist data using a relational/NoSQL database (i.e., MariaDB, PostgreSQL, MySQL, or MongoDB) **other than** Redis **and** use Redis as a cache/buffer to improve throughputs.
7. Use **Go** to build your own web server, proxy for WebSocket/HTTP, and/or load balancer, **avoiding** existing solutions such as Nginx, HAProxy, Apache, and Node.js.
8. Implement a load testing tool using Locust and perform load testing for your application. Write a locust file to simulate a user who visits the web user interface to receive and send messages. Use the node port `30223` to expose the Locust user interface. (Hint: Locust provides plugins to handle the WebSocket protocol)

### Part III: Report
You need to write a report `report.pdf` (min 3 pages and max 5 pages) that explains the system architecture and implementation of your solution. The report should at least cover:
* Description of the system architecture
* Design of each component
* How the components interact with each other
* How you implemented the additional feature
* Efforts you have made to improve the performance, if any
* Resources that you referred to

Additionally, your report should include the following information, which does not count towards the page limit, depending on the additional feature you choose to implement:
* For features one to five, provide documentation on how to test the feature along with screenshots to demonstrate the feature is working.
* For features six or seven, ensure that you give clear descriptions of the system architecture and design of each component. Also, your code should be clearly and effectively commented. The code comments should include descriptions of all key variables. Specific purpose is noted for each function, control structure, input requirements, and output.
* For feature eight, your report should mention the path of locustfile in your submission. Your locustfile should be clearly and effectively commented. The code comments should include descriptions of all key variables. Specific purpose is noted for each function, control structure, input requirements, and output. Also, you should perform load tests under different parameters and provide the Locust charts and discuss the results in your report.


### Technicalities
Besides the requirements mentioned above, you should strictly follow the following technicalities/instructions:
* Clone this repository and implement your solution in the `EECS4222_Project_2` directory.
* Make sure your repository is private.
* The three lines in the `info.txt` file include information about you (full name, 9-digit student ID, and your feature choice (see [Requirements Part II](#part-ii-feature-implementation)). Please update the `info.txt` file accordingly. For example, your name is `Foo Bar`, your student ID is `123456789`, and you choose to implement feature #3 for Requirements Part II.
The `info.txt` file should be as follows:
    ```
    Foo Bar
    123456789
    3
    ```
* Define Kubernetes objects in `chatapp.yaml` to properly coordinate components of your application.
* All the images referenced in `chatapp.yaml` should have the registry hostname as `harbor.pacslab.ca` and the repository name as `YOUR_STUDENT_ID`. Therefore, your image name should be something like `harbor.pacslab.ca/YOUR_STUDENT_ID/your-image-a`. (Hint: see [Appendix 2](#2-private-image-registry))
* Push all the images used in `chatapp.yaml` (include the images you developed and those you took from Docker Hub) to the private image registry for this course. (Hint: see [Appendix 2](#2-private-image-registry))
* Use `regcred` as the name of the secret for image pulls, which means you need to add the following field for each pod/deployment. (Hint: see [Appendix 3](#3-pull-images-from-a-private-registry-in-kubernetes-cluster))
    ```
    imagePullSecrets:
    - name: regcred
    ```
* Place the source code of each Docker image **you developed** for the project in the `images` directory. For example, the [`images`](./images/) directory shows the correct directory structure when you developed three images, `harbor.pacslab.ca/YOUR_STUDENT_ID/image-a:latest`, `harbor.pacslab.ca/YOUR_STUDENT_ID/image-b:latest`, and `harbor.pacslab.ca/YOUR_STUDENT_ID/image-c:latest`. Do **not** include those official images taken from Docker Hub (e.g., Redis and MariaDB) without any modifications in the `images` directory.
* Do **not** push the images you developed to any public image registry (e.g., Docker Hub and ghcr.io).

## Submission
You need to zip your repository and submit as one zip file with the name `project2.zip` on eClass by the due date. The directory structure in project2.zip should look like this:
```
EECS4222_Project_2/
├─ images/
│  ├─ image-a/
│  │  ├─ Dockerefile
|  |  ├─ ...
│  ├─ image-b/
│  │  ├─ Dockerefile
|  |  ├─ ...
│  ├─ ...
├─ chatapp.yaml
├─ report.pdf
├─ info.txt
├─ README.md
├─ FAQ.md
├─ .gitignore
```

Also, make sure to push all the images used in `chatapp.yaml` (include the images you developed and those you took from Docker Hub) to the private image registry for this course (see [Technicalities](#technicalities)).

**Important**: To ensure you do not lose significant marks, it is important to adhere strictly to the specified directory structure and instructions, since an automated judge will evaluate the majority part of your solution programmatically. Implementations with incorrect format/directory structure/technicalities will be marked as 0.

## Evaluation
An automated judge will evaluate the majority part of your solution programmatically. TAs will evaluate any parts that cannot be assessed through automation. The grading criteria are as follows:

* Part I - Distributed Chat Application (50%): Your solution has completed **all** the tasks specified in [Requirements Part I](#part-i-implement-a-distributed-chat-application).
* Part II - Feature Implementation (30%): Your solution has completed **all** the tasks specified in [Requirements Part II](#part-ii-feature-implementation). In order to get these marks, your chat application built for Part I must function properly.
* Part III - Report (20%): The report has completed **all** the requirements specified in [Requirements Part III](#part-iii-report). In order to get full marks for this part, your chat application built for Part I and the additional feature implemented for Part II must function properly.

**More information on evaluation**: The automated judge will first build and push all your images located in the `images` directory using the following commands:
```
docker build -t harbor.pacslab.ca/YOUR_STUDENT_ID/YOUR_IMAGE_NAME:IMAGE_TAG .
docker push harbor.pacslab.ca/YOUR_STUDENT_ID/YOUR_IMAGE_NAME:IMAGE_TAG
```
After building and pushing all the images, the automated judge will deploy your solution.
```
kubectl create namespace chatapp
kubectl apply -f chatapp.yaml -n chatapp
```
Then, the judge will try to access the web user interface by going to the address `http://MASTER_IP:30222`.


## Appendix
### 1. Example Chat Application
An example chat application is provided in [this repository](https://github.com/pacslab/chatapp). You can modify its source code and rewrite it in Go to implement the chat application for this project. Please note that you need to obtain your credential for the private registry and create a secret named `regcred` in the `chatapp` namespace (see below) before you try to deploy the example chat application.

### 2. Private Image Registry
Since public image registries (e.g., Docker Hub and GitHub Container Registry) usually have a harsh limit on the number of pulls per hour, we use a [private image registry](https://harbor.pacslab.ca/) for this project. We built the registry using [Harbor](https://goharbor.io/), which is an open-source image registry designed to be a trusted cloud native repository for Kubernetes. You can obtain your credential for the private registry by submitting this [form](https://forms.gle/JhwQTvFPnUAcnhRZ6). Your account will be activated within 24 hours after you submit the form.

Harbor provides a web-based [management portal](https://harbor.pacslab.ca/). Your username is your student ID, and your password is the concatenation of "Eecs4222" and the password you entered in the form. Your student ID will be used to create a private project in the registry with a storage quota of 10 GB. There are no limits on the number of images and no rate limits for image push/pull. All images in your private project will have a prefix of `harbor.pacslab.ca/YOUR_STUDENT_ID`. For example, if your student ID is `123456789`, and you have an image `go-backend` and a tag of `latest`, you can refer to the image as `harbor.pacslab.ca/123456789/go-backend:latest`. Please note that the default `latest` tag is often omitted.

Authentication is required to push/pull images to/from the private image registry using the `docker` command. You can use the `docker login harbor.pacslab.ca` command to log into the private image registry.
```
$ docker login harbor.pacslab.ca
Username: YOUR_STUDENT_ID
Password:
WARNING! Your password will be stored unencrypted in /home/ubuntu/.docker/config.json.
Configure a credential helper to remove this warning. See
https://docs.docker.com/engine/reference/commandline/login/#credentials-store

Login Succeeded
```

The following commands show how to push the official Redis image to the private image registry.

```
docker pull redis:7.0.8-alpine
docker tag redis:7.0.8-alpine harbor.pacslab.ca/YOUR_STUDENT_ID/redis:7.0.8-alpine
docker push harbor.pacslab.ca/YOUR_STUDENT_ID/redis:7.0.8-alpine
```
The following commands demonstrate how to build your own image and push it to the private image registry.

```
cd DIRECTORY_WITH_DOCKERFILE
docker build -t harbor.pacslab.ca/YOUR_STUDENT_ID/YOUR_IMAGE_NAME:TAG .
docker push harbor.pacslab.ca/YOUR_STUDENT_ID/YOUR_IMAGE_NAME:TAG
```
For example, your student ID is `123456789`, the image name is `my-awesome-frontend`, and the tag is `latest`.
```
cd my-awesome-frontend
docker build -t harbor.pacslab.ca/123456789/my-awesome-frontend:latest .
docker push harbor.pacslab.ca/123456789/my-awesome-frontend:latest
```

### 3. Pull Images from a Private Registry in Kubernetes Cluster
A Kubernetes cluster uses the Secret of `kubernetes.io/dockerconfigjson` type to authenticate with a container registry to pull a private image. The following command creates a secret named `regcred` in the `chatapp` namespace, which allows the Kubernetes cluster to pull the image from our private registry.
```
kubectl create secret docker-registry regcred --docker-server=harbor.pacslab.ca --docker-username=YOUR_STUDENT_ID --docker-password=YOUR_PRIVATE_REGISTRY_PASSWORD -n chatapp
```
Here is a manifest for an example Deployment that needs access to your credentials in `regcred`:
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: redis
spec:
  selector:
    matchLabels:
      app: redis
  template:
    metadata:
      labels:
        app: redis
    spec:
      terminationGracePeriodSeconds: 5
      containers:
      - name: redis
        image: harbor.pacslab.ca/123456789/redis:alpine
        ports:
        - containerPort: 6379
        resources:
          requests:
            cpu: 100m
            memory: 64Mi
          limits:
            cpu: 200m
            memory: 256Mi
      restartPolicy: Always
      imagePullSecrets:
      - name: regcred
```
Please refer to [this page](https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/) for more information on using a private image registry in Kubernetes.

### 4. Docker Tutorials / Links / Docs
* [Get started with Docker](https://docs.docker.com/get-started/)
* [Dockerfile reference](https://docs.docker.com/engine/reference/builder/)
* [A Docker Tutorial for Beginners](https://docker-curriculum.com/)
* [Containerization and Docker](https://decal.ocf.berkeley.edu/archives/2018-fall/labs/a11)

### 5. Q&As
Please refer to [QA.md](./QA.md).
