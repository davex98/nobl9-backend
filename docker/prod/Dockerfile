FROM golang:buster as builder

WORKDIR $GOPATH/src/gitlab/davex98/nobl9-backend/random-generator
COPY random-generator .

RUN go get -d -v
RUN go mod download
RUN go mod verify
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/random-generator .

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /go/bin/random-generator .
USER nonroot:nonroot
ENV PORT=8080
ENV CONCURRENT_REQUESTS=10

ENTRYPOINT ["/random-generator"]
