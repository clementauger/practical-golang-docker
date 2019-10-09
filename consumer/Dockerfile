#
FROM golang:latest AS grelot

RUN mkdir /gobin
ENV GOBIN "/gobin"
ENV PATH "${GOBIN}:${PATH}"

RUN go get github.com/clementauger/grelot
RUN go install github.com/clementauger/grelot

#
FROM golang:latest AS builder

RUN mkdir /gobin
RUN mkdir -p /go/src/github.com/clementauger/practical-golang-docker/consumer

WORKDIR /go/src/github.com/clementauger/practical-golang-docker/
COPY . .

WORKDIR consumer
RUN go mod download
RUN CGO_ENABLED=0 go build -o /gobin/consumer .

#
FROM golang:latest AS dev

RUN mkdir /gobin
ENV GOBIN "/gobin"
ENV PATH "${GOBIN}:${PATH}"
COPY --from=grelot /gobin/grelot /gobin/grelot

WORKDIR /go/src/github.com/clementauger/practical-golang-docker
COPY . .
WORKDIR consumer

EXPOSE ${EXPOSE}
CMD ["make", "reload"]

#
FROM scratch

COPY --from=builder /gobin/consumer /bin/consumer

EXPOSE ${EXPOSE}
CMD ["/bin/consumer"]
