#
FROM golang:latest AS reflex

RUN mkdir /gobin
ENV GOBIN "/gobin"
ENV PATH "${GOBIN}:${PATH}"

RUN go get github.com/cespare/reflex
RUN go install github.com/cespare/reflex

#
FROM golang:latest AS builder

RUN mkdir /gobin
RUN mkdir -p /gow/src/github.com/clementauger/practical-golang-docker/producer

WORKDIR /gow/src/github.com/clementauger/practical-golang-docker/
COPY . .

WORKDIR /gow/src/github.com/clementauger/practical-golang-docker/producer
RUN go mod download
RUN CGO_ENABLED=0 go build -o /gobin/producer .

#
FROM golang:latest AS dev

RUN mkdir /gobin
ENV GOBIN "/gobin"
ENV PATH "${GOBIN}:${PATH}"
COPY --from=reflex /gobin/reflex /gobin/reflex

WORKDIR /gow/src/github.com/clementauger/practical-golang-docker
COPY . .
WORKDIR producer/

EXPOSE ${EXPOSE}
CMD ["make", "reload"]

#
FROM scratch

COPY --from=builder /gobin/producer /bin/producer

EXPOSE ${EXPOSE}
CMD ["/bin/producer"]
