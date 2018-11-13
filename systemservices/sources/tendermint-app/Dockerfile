FROM golang:1.10.4
WORKDIR /go/src/github.com/mesg-foundation/core/systemservices/sources/tendermint-app
COPY . .
RUN go install ./...
CMD ["tendermint-app"]
