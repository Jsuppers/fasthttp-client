ARG PROJECT=fasthttp-client

FROM golang:1.13 as builder

ARG PROJECT
WORKDIR /go/${PROJECT}
COPY . /go/${PROJECT}
RUN make build

FROM scratch

ARG PROJECT
COPY --from=builder /go/${PROJECT}/bin/${PROJECT} service

CMD ["./service"]
