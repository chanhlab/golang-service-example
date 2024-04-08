FROM golang:1.15-alpine as builder
RUN apk add --no-cache ca-certificates git && \
  wget -qO/go/bin/dep https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 && \
  chmod +x /go/bin/dep

WORKDIR /go/src/github.com/chanhlab/golang-service-example

COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

COPY . .

# Build
RUN cd cmd/ && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -installsuffix cgo -o /server

FROM alpine as release
COPY --from=builder /server /server

CMD ["/server api"]
