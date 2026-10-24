# Can't be scratch because we need sh and tee for the Github Action
# so we can write the trunkver to GITHUB_OUTPUT
FROM busybox:1.38.0-glibc@sha256:99813cb178ace2cd9c6c53ee3bddef5d2f5708aa438335111d7a9e3c5e1aaebe

ARG TARGETOS
ARG TARGETARCH
COPY dist/trunkver_${TARGETOS}_${TARGETARCH} /usr/local/bin/trunkver
COPY dist/trunkver_${TARGETOS}_${TARGETARCH}.cosign.bundle /usr/local/share/trunkver.cosign.bundle

ENTRYPOINT ["/usr/local/bin/trunkver"]
