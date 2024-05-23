FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN go test -v
RUN go build -o trunkver

# Can't be scratch because we need sh and tee for the Github Action 
# so we can write the trunkver to GITHUB_OUTPUT
FROM busybox

COPY --from=builder /app/trunkver /trunkver

ENTRYPOINT ["/trunkver"]