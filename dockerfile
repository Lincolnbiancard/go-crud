# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.16-alpine

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# ... the rest of the Dockerfile is ...
# ...   omitted from this example   ...

COPY *.go ./

RUN go build -o /docker-gs-ping

# Install Gorm ORM 
# go get -u gorm.io/gorm

EXPOSE 8080

CMD [ "/docker-gs-ping" ]