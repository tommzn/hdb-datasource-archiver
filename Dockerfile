FROM --platform=linux/${TARGETARCH:-amd64} golang:1.19-alpine as builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

ENV CGO_ENABLED=1
ENV GO111MODULE=on
ENV GOOS=${TARGETOS}
ENV GOARCH=${TARGETARCH}

WORKDIR /go/build

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

RUN apk -U add ca-certificates
RUN apk update && apk upgrade && apk add pkgconf git bash build-base sudo gcc musl-dev

COPY .  .

RUN go build -tags musl -ldflags '-extldflags "-static"' -o build_artifact_bin

FROM --platform=linux/${TARGETARCH:-amd64} gcr.io/distroless/static:nonroot


LABEL org.opencontainers.image.source=https://github.com/tommzn/hdb-datasource-archiver

WORKDIR /go

COPY --from=builder /go/build/build_artifact_bin hdb-bin
USER nonroot:nonroot

ENTRYPOINT ["/go/hdb-bin"]
