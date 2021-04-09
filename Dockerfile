ARG GO_VERSION=1.16
FROM golang:${GO_VERSION}-alpine as builder

WORKDIR /src
ARG EXPORTER_VERSION=
RUN test -n "$EXPORTER_VERSION"
COPY ./go.mod ./go.mod
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-X main.Version=${EXPORTER_VERSION}" -o mn-exporter .

FROM scratch

LABEL maintainer="Will Fantom <willf@ntom.dev>"

WORKDIR /exporter
COPY --from=builder /src/mn-exporter .
ENTRYPOINT [ "./mn-exporter" ]
