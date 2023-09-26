# Use official golang image from docker hub with golang version 1.20.5 as dev environment
FROM golang:1.20.5 as dev

# Set environment variables for goproxy and gopath
ENV GOPROXY=https://goproxy.io,direct
ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

RUN go mod download


COPY . .
RUN go install github.com/cespare/reflex@latest



# Makefile as entrypoint
ENTRYPOINT ["make"]
CMD ["dev-docker"]

# Use official golang image from docker hub with golang version 1.20.5 as production environment
FROM golang:1.20.5 as prod

# Set environment variables for goproxy and gopath
ENV GOPROXY=https://goproxy.io,direct
ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

RUN go mod download


COPY . .
