ARG GOLANG_VERSION=1.17
FROM golang:${GOLANG_VERSION}-buster as builder
ARG GOPROXY=https://goproxy.cn
WORKDIR ${GOPATH}/src/github.com/projectxpolaris/webcontainer

COPY go.mod .
RUN go mod download
COPY . .
RUN go mod tidy
RUN go build -o ${GOPATH}/bin/webcontainer ./main.go

FROM debian:buster-slim
COPY --from=builder /usr/local/lib /usr/local/lib
COPY --from=builder /etc/ssl/certs /etc/ssl/certs


COPY --from=builder /go/bin/webcontainer /usr/local/bin/webcontainer
ADD ./static /static
ENTRYPOINT ["/usr/local/bin/webcontainer","run"]