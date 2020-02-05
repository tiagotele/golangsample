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
docker run -it -p 8093:8093 tiagotele/golangsample:<VERSION_TAG>
```

## Pushing to DockerHub
```
docker login 
docker push tiagotele/golangsample:<VERSION_TAG>
```

## Docker image
Docker image are available at [DockerHub](https://hub.docker.com/repository/docker/tiagotele/golangsample).