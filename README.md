# EventBus

## Prerequisites
- Install go (preferably version go1.18.1 or greater) and add to path
- Install swag for generating documentation and add in path (in the Swagger - Prerequisites)
- docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.13-management

## Build package
- Build the package by specifying the entry point file
```bash
go build cmd\kafka\main.go
```

## Run package
- Run the package by specifying the entry point file
```bash
go run cmd\kafka\main.go
```

## Swagger - Prerequisites
- Install swag
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```
- Update PATH to include $GOPATH/bin where swag should be installed

## Swagger - Generation command 
```bash
swag init --parseDependency --parseDepth 3 --dir ./cmd/kafka/,./pkg/api/v1/
```

