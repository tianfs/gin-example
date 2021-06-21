FROM golang:1.15.3-alpine

WORKDIR $GOPATH/src/gin-example
ADD ./ $GOPATH/src/gin-example

ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io,direct"
ENV CONFIG_ENV=dev

#RUN go mod download

RUN go build .

ENTRYPOINT ["./gin-example"]

EXPOSE 8084