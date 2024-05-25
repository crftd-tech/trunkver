FROM golang:1.22 AS builder
ARG VERSION=0.0.0-HEAD-local

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY internal internal

RUN go test -v
RUN go build -o trunkver -ldflags "-X main.Version=${VERSION}"

FROM debian AS smoke

ADD --chmod=755 --chown=0 https://github.com/SamirTalwar/smoke/releases/download/v2.4.0/smoke-v2.4.0-Linux-x86_64 /smoke
COPY smoke.yaml /smoke.yaml
COPY --from=builder /app/trunkver /trunkver
ENV PATH="/:${PATH}"
RUN /smoke /smoke.yaml

# Can't be scratch because we need sh and tee for the Github Action
# so we can write the trunkver to GITHUB_OUTPUT
FROM busybox

COPY --from=smoke /trunkver /trunkver

ENTRYPOINT ["/trunkver"]
