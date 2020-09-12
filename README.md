# Sample REST API with Golang

## How to Run
```
go run main.go
```

## Building docker image
```
docker build -t tiagotele/golangsample:<VERSION_TAG> . 
```

## Running docker image
```
docker run -it -p 8081:8081 tiagotele/golangsample:<VERSION_TAG>
```

## Pushing to DockerHub
```
docker login 
docker push tiagotele/golangsample:<VERSION_TAG>
```

## Docker image
Docker image are available at [DockerHub](https://hub.docker.com/r/tiagotele/golangsample).


# For testing purposes
## Running locally with Docker compose
```
docker-compose up
```

# Helm chart
A [Helm Chart](https://helm.sh/) is available for docker image of this repository is available [here](https://github.com/tiagotele/myapp-chart).
