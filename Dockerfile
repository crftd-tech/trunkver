FROM golang:1.23 AS builder
ARG VERSION

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY smoke.yaml Makefile *.go ./
COPY cmd cmd
COPY internal internal

RUN make test out/trunkver_linux_amd64 spec "VERSION=${VERSION}" 

# Can't be scratch because we need sh and tee for the Github Action
# so we can write the trunkver to GITHUB_OUTPUT
FROM busybox

COPY --from=builder /app/out/trunkver_linux_amd64 /trunkver

ENTRYPOINT ["/trunkver"]
