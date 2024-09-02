# all-http-200-under-port-9999

This repository contains a Go application that starts HTTP servers on all ports from 2 to 9999. Each server responds with `200 OK` to every HTTP request.  
The application is packaged in a Docker container for easy deployment.

## Features

- Launches HTTP servers on all ports from 2 to 9999.
- Responds with `200 OK` to every HTTP request.
- Lightweight and easy to deploy as a Docker container.

## DockerHub
The Docker image for this application is available on DockerHub at:
[DockerHub Repository](https://hub.docker.com/r/souta/all-http-200-under-port-9999)

### Use the Docker Image from DockerHub

```
$ docker pull souta/all-http-200-under-port-9999
$ docker run -d -p 2-9999:2-9999 souta/all-http-200-under-port-9999
```

### Push the Docker Image to DockerHub

If you want to push the Docker image to DockerHub, make sure you are logged in to Docker and then push the image:

```
$ docker login
$ docker push your-dockerhub-username/all-http-200-under-port-9999
```

## Use Case: Passing ECS Health Checks
This Docker container is ideal for passing health checks on AWS ECS.  
Since the container listens on all ports from 2 to 9999 and returns 200 OK for any HTTP request, you can configure your ECS service's health check to target any of these ports. Here's how you might configure the health check:

- Health Check Path: / (or any valid path)
- Port: Choose any port between 2 and 9999 
- Protocol: HTTP 
- Response Code: 200

By configuring your ECS service with these settings, the container will always return a 200 OK response, ensuring that the health check passes successfully.

