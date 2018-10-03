FROM golang:1.11.0-stretch

WORKDIR /core

COPY . .

RUN go install -mod=readonly ./vendor/github.com/go-bindata/go-bindata/...

# Attach volumes
# VOLUME /core

# Execute go bin data
RUN /go/bin/go-bindata
