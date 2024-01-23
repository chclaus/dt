FROM golang:alpine
RUN apk add --no-cache git
RUN go install github.com/chclaus/dt@latest

FROM alpine:latest  
COPY --from=0 /go/bin/dt /usr/local/bin
ENTRYPOINT ["/usr/local/bin/dt"]
