FROM golang:1.15.3

WORKDIR $GOPATH/src/gin-example
ADD ./ $GOPATH/src/gin-example

ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io,direct"

RUN go build .
ENTRYPOINT ["./k8sgo"]

EXPOSE 8083