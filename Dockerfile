FROM golang:1.15-alpine as builder
RUN apk add --no-cache ca-certificates git && \
  wget -qO/go/bin/dep https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 && \
  chmod +x /go/bin/dep

WORKDIR /go/src/github.com/chanhteam/golang-service-example

COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

COPY . .

# Build API
RUN cd cmd/api && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -installsuffix cgo -o /api

# Build Migration
RUN cd cmd/migration && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -installsuffix cgo -o /migration

# Build Worker
RUN cd cmd/worker && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -installsuffix cgo -o /worker

FROM alpine as release
COPY --from=builder /api /api
COPY --from=builder /migration /migration
COPY --from=builder /worker /worker

CMD ["/api","/migration","/worker"]
