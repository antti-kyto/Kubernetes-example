##
## -- BUILD --
##
FROM golang as builder

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

### CGO has to be disabled cross platform builds
### Build the Go app for a linux OS
ENV CGO_ENABLED=0
RUN GOOS=linux go build -o /my-api

##
## -- DEPLOY --
##
FROM alpine:3.13.1
COPY --from=builder /my-api /my-api

EXPOSE 3000

CMD ["/my-api"]