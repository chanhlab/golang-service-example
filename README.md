[![Build Status](https://travis-ci.org/chanhteam/golang-service-example.svg?branch=master)](https://travis-ci.org/chanhteam/golang-service-example)
![Tests](https://github.com/chanhteam/golang-service-example/workflows/Tests/badge.svg)
[![codecov](https://codecov.io/gh/chanhteam/golang-service-example/branch/master/graph/badge.svg?token=S02QXOCQJ9)](https://codecov.io/gh/chanhteam/golang-service-example)

## Get Started

### Migrate Database

```
./scripts/run-migration.sh
```

### Start Grpc and Rest API

```
./scripts/start-api.sh
```

### Configurations

```
// .env

# API PORTS
GRPC_PORT=5000
HTTP_PORT=5001

# MYSQL CONFIGURATION
MYSQL_HOST=127.0.0.1
MYSQL_DATABASE=golang_sample
MYSQL_USERNAME=admin
MYSQL_PASSWORD=Hie8oox9ahhohsh

# LOG CONFIGURATION
LOG_LEVEL=-1
LOG_TIME_FORMAT=2006-01-02T15:04:05Z07:00
```

### Build Binary

```
./scripts/build.sh
```
```
cd build/
```

### Update Protocol Buffers

```
git submodule update --remote
```

## Structure

```
├── Dockerfile
├── Dockerfile.test
├── Makefile
├── README.md
├── build
│   ├── api
│   ├── migration
│   └── worker
├── cSpell.json
├── cmd
│   ├── api
│   │   └── main.go
│   ├── migration
│   │   └── main.go
│   └── worker
│       └── main.go
├── config
│   └── config.go
├── coverage.html
├── coverage.out
├── docker-compose.yml
├── docker-env
│   └── development.env
├── go.mod
├── go.sum
├── internal
│   ├── grpc
│   │   └── grpc.go
│   ├── models
│   │   ├── credential.go
│   │   ├── credential_test.go
│   │   └── mocks
│   │       └── credential_mock.go
│   ├── rest
│   │   └── rest.go
│   ├── services
│   │   ├── credential.go
│   │   └── credential_test.go
│   └── workers
│       ├── credential_worker.go
│       └── middleware.go
├── pkg
│   ├── db
│   │   └── mysql
│   │       ├── mysql.go
│   │       └── mysql_test.go
│   ├── env.go
│   ├── env_test.go
│   ├── grpc
│   │   └── middleware
│   │       ├── context.go
│   │       └── logger.go
│   ├── logger
│   │   └── logger.go
│   ├── rest
│   │   └── middleware
│   │       ├── logger.go
│   │       ├── request_id.go
│   │       └── tracing.go
│   ├── timestamp.go
│   └── timestamp_test.go
├── pre-commit
├── protobuf
│   └── v1
│       └── credential
│           ├── credential.pb.go
│           └── credential.pb.gw.go
└── protos
    ├── README.md
    ├── cSpell.json
    ├── proto-generate.sh
    └── v1
        └── credential.proto
```


## Services
### Credential

## Unit Testing

```
go test -covermode=count -coverprofile=coverage.out fmt  ./... -v cover
go test -covermode=count -coverprofile=coverage.out ./... -v cover
go tool cover -html=coverage.out -o coverage.html
```
