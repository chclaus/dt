FROM golang:alpine
RUN apk add --no-cache git
RUN go get -u github.com/chclaus/dt

FROM alpine:latest  
COPY --from=0 /go/bin/dt /usr/local/bin
ENTRYPOINT ["/usr/local/bin/dt"]
